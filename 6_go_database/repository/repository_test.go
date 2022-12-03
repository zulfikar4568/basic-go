package repository

import (
	"context"
	"fmt"
	godatabase "go-database"
	"go-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(godatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "isnaen@gmail.com",
		Comment: "Ini adalah Komen",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(godatabase.GetConnection())

	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(godatabase.GetConnection())

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(comments)
}
