package handler

import (
	"EFpractic2/models"
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var testBook1 = models.Book{BookName: "name1", BookYear: 2002, BookNew: false}
var testBook2 = models.Book{BookName: "name2", BookYear: 2023, BookNew: true}

var bookID = 1

func TestHandler_CreateBook(t *testing.T) {
	testInit()
	ctx := context.Background()
	srv := NewMockBookAct(t)
	srv.On("CreateBook", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("models.Book")).Return(nil).Once()
	err := srv.CreateBook(ctx, testBook1)
	require.NoError(t, err)
}

func TestHandler_DeleteBook(t *testing.T) {
	testInit()
	ctx := context.Background()
	srv := NewMockBookAct(t)
	srv.On("DeleteBook", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("int")).Return(nil).Once()
	err := srv.DeleteBook(ctx, bookID)
	require.NoError(t, err)
}

func TestHandler_UpdateBook(t *testing.T) {
	testInit()
	ctx := context.Background()
	srv := NewMockBookAct(t)
	srv.On("UpdateBook", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("models.Book")).Return(nil).Once()
	err := srv.UpdateBook(ctx, testBook2)
	require.NoError(t, err)
}

func TestHandler_GetBook(t *testing.T) {
	testInit()
	ctx := context.Background()
	srv := NewMockBookAct(t)
	srv.On("GetBook", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("int")).Return(models.Book{}, nil).Once()
	_, err := srv.GetBook(ctx, bookID)
	require.NoError(t, err)
}
