package repository

import (
	"EFpractic2/models"
	"EFpractic2/pkg/utils"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

type UserAuthPostgres struct {
	db *pgxpool.Pool
}

func NewUserAuthPostgres(db *pgxpool.Pool) *UserAuthPostgres {
	return &UserAuthPostgres{db: db}
}

func (r *UserAuthPostgres) CreateAuthUser(ctx context.Context, user *models.UserAuth) (int, error) {
	var id int
	sqlRow := `insert into usersauth (name, age, regular, password, refreshtoken) values($1, $2, $3, $4, $5) returning id`
	err := r.db.QueryRow(ctx, sqlRow, user.UserName, user.UserAge, user.UserIsRegular, user.Password, "").Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error while user creating: %v", err)
	}
	return id, err
}

func (r *UserAuthPostgres) GetAuthUser(ctx context.Context, id int) (models.UserAuth, error) {
	var user = models.UserAuth{}
	return user, nil
}

func (r *UserAuthPostgres) UpdateRefreshToken(ctx context.Context, rt string, id int) error {
	_, errInsert := r.db.Exec(ctx, "UPDATE usersauth SET refreshtoken = $1 WHERE id = $2", rt, id)
	if errInsert != nil {
		return fmt.Errorf("update user error %w", errInsert)
	}
	return nil
}

func (r *UserAuthPostgres) GetUserById(ctx context.Context, userId int) (models.UserAuth, error) {
	user := models.UserAuth{}
	err := r.db.QueryRow(ctx, "select usersauth.id, usersauth.name, usersauth.age, usersauth.regular, usersauth.password, usersauth.refreshtoken from usersauth where id=$1", userId).Scan(
		&user.UserId, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password, &user.RefreshToken)
	if err != nil {
		return user, fmt.Errorf("get user error %w", err)
	}

	return user, nil
}

func (r *UserAuthPostgres) GetUserWithRefreshToken(ctx context.Context, rt string) (models.UserAuth, error) {
	user := models.UserAuth{}
	userId, errRT := utils.ParseToken(rt)
	if errRT != nil {
		log.WithFields(log.Fields{
			"ERROR": errRT,
		}).Info("refresh token expired")
		return models.UserAuth{}, errRT
	}
	err := r.db.QueryRow(ctx, "select usersauth.id, usersauth.name, usersauth.age, usersauth.regular, usersauth.password, usersauth.refreshtoken from usersauth where id=$1", userId).Scan(
		&user.UserId, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password, &user.RefreshToken)
	if err != nil {
		log.WithFields(log.Fields{
			"ERROR": errRT,
		}).Info("Error while getting user")
		return models.UserAuth{}, errRT
	}
	return user, nil
}

/*
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
}*/
