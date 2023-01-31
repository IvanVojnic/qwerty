package repository

import (
	"EFpractic2/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"math/rand"
	"strings"
)

type UserAuthPostgres struct {
	db *pgxpool.Pool
}

func NewUserAuthPostgres(db *pgxpool.Pool) *UserAuthPostgres {
	return &UserAuthPostgres{db: db}
}

func (r *UserAuthPostgres) CreateAuthUser(ctx context.Context, user *models.UserAuth) (string, int, error) {
	rt := generateRT()
	var id int
	sqlRow := `insert into userauth (name, age, regular, password, refreshtoken) values($1, $2, $3, $4, $5) returning id`
	err := r.db.QueryRow(ctx, sqlRow, user.UserName, user.UserAge, user.UserIsRegular, user.Password, rt).Scan(&id)
	if err != nil {
		return " ", 0, fmt.Errorf("error while user creating: %v", err)
	}
	return rt, id, err
}

func (r *UserAuthPostgres) GetAuthUser(ctx context.Context, id int) (models.UserAuth, error) {
	var user = models.UserAuth{}
	return user, nil
}

func generateRT() string {
	sec1 := rand.New(rand.NewSource(60))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[sec1.Intn(len(chars))])
	}
	str := b.String() // Например "ExcbsVQs"
	return str
}

func (r *UserAuthPostgres) GetUserId(ctx context.Context, token string) (int, error) {
	user := models.UserAuth{}
	err := r.db.QueryRow(ctx, "select id from users where refreshToken=$1", token).Scan(
		&user.UserId, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password)
	if err != nil {
		return user.UserId, fmt.Errorf("get user error %w", err)
	}
	return user.UserId, nil
}
