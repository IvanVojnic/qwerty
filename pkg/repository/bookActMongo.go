package repository

import (
	"EFpractic2/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

// BookActMongo is a wrapper to db object
type BookActMongo struct {
	db   *mongo.Database
	coll *mongo.Collection
}

// NewBookActMongo used to init BookAP
func NewBookActMongo(db *mongo.Database) *BookActMongo {
	return &BookActMongo{
		db:   db,
		coll: db.Collection("books"),
	}
}

// CreateBook used to create book
func (r *BookActMongo) CreateBook(ctx context.Context, book *models.Book) error {
	ID := uuid.New()
	book.BookID = ID
	_, err := r.coll.InsertOne(ctx, book)
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	return nil
}

// GetBookId used to get book with id
func (r *BookActMongo) GetBookId(ctx context.Context, bookName string) (uuid.UUID, error) {
	filter := bson.D{{"name", bookName}}
	var bookGet models.Book
	err := r.coll.FindOne(ctx, filter).Decode(&bookGet)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error while getting book ID %s", err)
	}
	return bookGet.BookID, nil
}

// UpdateBook used to update book
func (r *BookActMongo) UpdateBook(ctx context.Context, book models.Book) error {
	filter := bson.D{
		{"_id", book.BookID},
	}
	update := bson.D{
		{"$set", bson.D{
			{"bookname", book.BookName},
			{"bookyear", book.BookYear},
			{"booknew", book.BookNew},
		}},
	}
	_, err := r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error while updating %s", err)
	}
	return nil
}

// GetBook used to get book
func (r *BookActMongo) GetBookByName(ctx context.Context, bookName string) (models.Book, error) {
	book := models.Book{}
	filter := bson.D{{"bookname", bookName}} //nolint:govet
	err := r.coll.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return models.Book{}, fmt.Errorf("error while getting book %s", err)
	}
	return book, nil
}

// DeleteBook used to delete book
func (r *BookActMongo) DeleteBook(ctx context.Context, bookName string) error {
	filter := bson.D{{"bookname", bookName}} //
	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error while deleting %s", err)
	}
	return nil
}

// GetAllBooks used to get all books
func (r *BookActMongo) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	books := make([]models.Book, 0)
	query := bson.M{}
	cur, errGetAll := r.coll.Find(ctx, query)
	if errGetAll != nil {
		return books, fmt.Errorf("error get all, %s", errGetAll)
	}
	for cur.Next(ctx) {
		var book models.Book
		errCur := cur.Decode(&book)
		if errCur != nil {
			return books, fmt.Errorf("error while getting all books from cursor %s", errCur)
		}
		books = append(books, book)
	}
	if err := cur.Err(); err != nil {
		return books, fmt.Errorf("error %s", err)
	}
	cur.Close(ctx)
	return books, nil
}
