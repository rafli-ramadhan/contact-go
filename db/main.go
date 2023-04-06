package client

/*
import (
	"bootcamp/internal/db"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/goombaio/namegenerator"
)

var listContent map[string]string = map[string]string{
	"anjani": "0882278432523",
	"albert": "0832424539342",
	"tifani": "0832424539323",
}

var listName = []string{"Alli", "Anto", "Andi", "Andrew", "Anthoni", "Aug", "Bella", "Cika", "Dian", "Daniel", "El", "Gerald"}

func main() {
	mysql()
}

func mysql() {
	db := client.GetDB("mysql").GetMysqlConnection()

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	/*
	queryInsert := `insert into pelanggan (name, no_telp) values ("adji", "088227867533")`

	res, err := db.ExecContext(ctx, queryInsert)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}	
	fmt.Println(lastInsertID)

	row, err := res.RowsAffected()
	if err != nil{
		panic(err)
	}
	fmt.Println(row)
	*/

	/*// Query parameter -> to prevent SQL Injection
	queryInsert := `insert into pelanggan (name, no_telp) values (?, ?)`

	res, err := db.ExecContext(ctx, queryInsert, "andi", "088227867533")
	if err != nil {
		panic(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	
	fmt.Println(lastInsertID)
	*/

	/* // prepare statement 
	queryInsert := `insert into pelanggan (name, no_telp) values (?, ?)`
	prepareMysql, err := db.PrepareContext(ctx, queryInsert)
	for i, v := range listContent {
		_, err := prepareMysql.ExecContext(ctx, i, v)
		if err != nil {
			panic(err)
		}
	}

	for i, v := range listContent {
		_, err := db.ExecContext(ctx, queryInsert, i, v)
		if err != nil {
			panic(err)
		}
	}*/

	/*
	querySelect := `select id, name, no_telp from pelanggan`
	res, err := db.QueryContext(ctx, querySelect)
	if err != nil {
		panic(err)
	}

	var id int
	var name string
	var noTelp string

	for res.Next() {
		if err := res.Scan(&id, &name, &noTelp); err != nil {
			panic(err)
		}
		fmt.Printf("%d %s %s\n", id, name, noTelp)
	}*/

	/*
	querySelect := `select id, name, no_telp from pelanggan`
	res, err := db.QueryContext(ctx, querySelect)
	if err != nil {
		panic(err)
	}

	var id int
	var name string
	var noTelp sql.NullString

	for res.Next() {
		if err := res.Scan(&id, &name, &noTelp); err != nil {
			panic(err)
		}
		fmt.Printf("%d %s %v\n", id, name, noTelp)
	}*/

	/*// Database Transaction -> to prevent auto-commit when SQL execute
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into comments (email, comment) value (?,?)`
	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	stmnt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}

	seed := time.Now().UTC().UnixNano()
    nameGenerator := namegenerator.NewNameGenerator(seed)

	for i:=0 ;i<70 ;i++ {
		name := strings.Trim(nameGenerator.Generate(), " ")
		res, err := stmnt.ExecContext(ctx, fmt.Sprintf("%v@gmail.com", name), "test")
		if err != nil {
			panic(err)
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println(lastInsertID)
	}
	
	if err != nil {
		trx.Rollback()
	} else {
		trx.Commit()
	}

	defer db.Close()
}*/