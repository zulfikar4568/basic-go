package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

type People struct {
	id   string
	name string
}

type Customer struct {
	id                     string
	name                   string
	email                  sql.NullString
	balance                int64
	rating                 float64
	created_at, birth_date sql.NullTime
	married                bool
}

func TestDatabaseTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// do Txn
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	for i := 0; i < 10; i++ {
		email := "Isnaen " + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("ID :", id)
	}

	err = tx.Commit() // Save to DB
	// err = tx.Rollback() // Rollback
	if err != nil {
		panic(err)
	}
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "Zulfikar " + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println(id)
	}

}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, "isnaen@gmail.com", "Ini adalah Comment")

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(insertId)
}

func TestSQLInjectionSolving(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	var username string = "admin'; #" // BAHAYA, ini bisa di SQL Injection
	var password string = "salah"

	queryUser := "SELECT username, password FROM user WHERE username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, queryUser, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)

		if err != nil {
			panic(err)
		}

		fmt.Println("Success Login with user:", username)
	} else {
		fmt.Println("Login Failed")
	}
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// var username string = "admin; DROP TABLE user;#"
	var username string = "admin'; #" // BAHAYA, ini bisa di SQL Injection
	var password string = "salah"

	queryUser := "SELECT username, password FROM user WHERE username='" + username + "' AND password='" + password + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, queryUser)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)

		if err != nil {
			panic(err)
		}

		fmt.Println("Success Login with user:", username)
	} else {
		fmt.Println("Login Failed")
	}
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	selectData := "SELECT id, name, email, balance, rating, created_at, married, birth_date FROM customer"
	rows, err := db.QueryContext(ctx, selectData)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var result []Customer

	for rows.Next() {
		var each = Customer{}
		var err = rows.Scan(&each.id, &each.name, &each.email, &each.balance, &each.rating, &each.created_at, &each.married, &each.birth_date)

		if err != nil {
			panic(err)
		}

		if each.email.Valid {
			fmt.Println(each.email.String)
		}

		result = append(result, each)
	}

	for _, each := range result {
		fmt.Println(each)
	}
}

func TestExecQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	selectData := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, selectData)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var result []People

	for rows.Next() {
		var each = People{}
		var err = rows.Scan(&each.id, &each.name)

		if err != nil {
			panic(err)
		}

		result = append(result, each)
	}

	for _, each := range result {
		fmt.Println(each)
	}
}

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('4', 'Doni')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data into sql")
}
