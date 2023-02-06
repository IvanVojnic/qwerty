package service

import (
	"context"
	"testing"

	"EFpractic2/models"
	"EFpractic2/pkg/mocks/repomocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// bookID is a test book id
const bookID = 1

var testBook1 = models.Book{BookName: "name1", BookYear: 2002, BookNew: false}
var testBook2 = models.Book{BookName: "name2", BookYear: 2023, BookNew: true}

func TestBookActSrv_CreateBook(t *testing.T) {
	repo := repomocks.NewBookAct(t)
	ctx := context.Background()
	repo.On("CreateBook", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*models.Book")).Return(nil).Once()
	service := NewBookActSrv(repo)
	err := service.CreateBook(ctx, testBook1)
	print("hello")
	require.NoError(t, err)
}

func TestBookActSrv_UpdateBook(t *testing.T) {
	repo := repomocks.NewBookAct(t)
	ctx := context.Background()
	service := NewBookActSrv(repo)
	repo.On("CreateBook", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*models.Book")).Return(nil).Once()
	errCreate := service.CreateBook(ctx, testBook1)
	require.NoError(t, errCreate)
	repo.On("UpdateBook", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("models.Book")).Return(nil).Once()
	errUpdate := service.UpdateBook(ctx, testBook2)
	require.NoError(t, errUpdate)
}

func TestBookActSrv_DeleteBook(t *testing.T) {
	repo := repomocks.NewBookAct(t)
	var bookID int
	ctx := context.Background()
	repo.On("DeleteBook", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("int")).Return(nil).Once()
	service := NewBookActSrv(repo)
	err := service.DeleteBook(ctx, bookID)
	require.NoError(t, err)
}

func TestBookActSrv_GetBook(t *testing.T) {
	repo := repomocks.NewBookAct(t)
	ctx := context.Background()
	repo.On("GetBook", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("int")).Return(models.Book{}, nil).Once()
	service := NewBookActSrv(repo)
	_, err := service.GetBook(ctx, bookID)
	require.NoError(t, err)
}
