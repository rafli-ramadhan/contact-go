package repository

import (
	client "contact-go/db"
	"contact-go/model"
	"context"
	"fmt"
	"time"
)

type contacthttp struct{}

func NewContactHTTPRepository() *contacthttp {
	return &contacthttp{}
}

func (repo *contacthttp) List() (result []model.Contact, err error) {
	db := client.GetDB("mysql").GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// prepare statement
	query := `SELECT id, name, no_telp FROM contact`
	stmt, err := trx.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	// defer stmt.Close()

	res, err := stmt.QueryContext(ctx)
	if err != nil {
		trx.Rollback()
		return
	}

	for res.Next() {
		var temp model.Contact
		res.Scan(&temp.Id, &temp.Name, &temp.NoTelp)

		result = append(result, temp)
	}

	fmt.Println(result)
	trx.Commit()
	return
}

func (repo *contacthttp) Add(req []model.ContactRequest) (contact []model.Contact, err error) {
	db := client.GetDB("mysql").GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	query := `INSERT INTO contact (name, no_telp) value (?,?)`
	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	for _, v := range req {
		res, err := stmt.ExecContext(ctx, v.Name, v.NoTelp)
		if err != nil {
			trx.Rollback()
			return []model.Contact{}, err
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			return []model.Contact{}, err
		}

		contact = append(contact, model.Contact{
			Id:   	int(lastID),
			Name: 	v.Name,
			NoTelp: v.NoTelp,
		})
	}

	trx.Commit()

	return
}

func (repo *contacthttp) Update(id int, req model.ContactRequest) (err error) {
	db := client.GetDB("mysql").GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `UPDATE contact SET name = ? and no_telp = ? WHERE id = ?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, req.Name, req.NoTelp, id)
	if err != nil {
		trx.Rollback()
		return
	}

	fmt.Println(res.RowsAffected())

	trx.Commit()
	
	return
}

func (repo *contacthttp) Delete(id int) (err error) {
	db := client.GetDB("mysql").GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `DELETE FROM contact WHERE id = ?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		trx.Rollback()
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println(rowAffected)

	trx.Commit()

	return
}
