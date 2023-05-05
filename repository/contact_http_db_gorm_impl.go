package repository

import (
	"contact-go/model"
	//"context"
	"database/sql"
	//"log"
	//"time"

	// "gorm.io/gorm"
)

type contactgormhttp struct{
	db *sql.DB
}

func NewContactGormHTTPRepository(client *sql.DB) *contactgormhttp {
	return &contactgormhttp{
		db: client,
	}
}

func (repo *contactgormhttp) List() (result []model.Contact, err error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// // prepare statement
	// query := `SELECT id, name, no_telp FROM contact`
	// stmt, err := repo.db.PrepareContext(ctx, query)
	// if err != nil {
	// 	log.Print("prepare list")
	// 	return
	// }

	// res, err := stmt.QueryContext(ctx)
	// if err != nil {
	// 	return
	// }

	// for res.Next() {
	// 	var temp model.Contact
	// 	res.Scan(&temp.Id, &temp.Name, &temp.NoTelp)
	// 	result = append(result, temp)
	// }

	return
}

func (repo *contactgormhttp) Add(req []model.ContactRequest) (result []model.Contact, err error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	// defer cancel()

	// query := `INSERT INTO contact (name, no_telp) value (?,?)`
	// trx, err := repo.db.BeginTx(ctx, nil)
	// if err != nil {
	// 	return
	// }

	// stmt, err := repo.db.PrepareContext(ctx, query)
	// if err != nil {
	// 	log.Print("prepare add")
	// 	return
	// }

	// for _, v := range req {
	// 	res, err := stmt.ExecContext(ctx, v.Name, v.NoTelp)
	// 	if err != nil {
	// 		trx.Rollback()
	// 		return []model.Contact{}, err
	// 	}

	// 	lastID, err := res.LastInsertId()
	// 	if err != nil {
	// 		return []model.Contact{}, err
	// 	}

	// 	result = append(result, model.Contact{
	// 		Id:   	int(lastID),
	// 		Name: 	v.Name,
	// 		NoTelp: v.NoTelp,
	// 	})
	// }

	// trx.Commit()

	return
}

func (repo *contactgormhttp) Update(id int, req model.ContactRequest) (err error) {
	// defer repo.db.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// trx, err := repo.db.BeginTx(ctx, nil)
	// if err != nil {
	// 	return
	// }

	// query := `UPDATE contact SET name = ?, no_telp = ? WHERE id = ?`
	// stmt, err := repo.db.PrepareContext(ctx, query)
	// if err != nil {
	// 	log.Print("prepare update")
	// 	return
	// }

	// _, err = stmt.ExecContext(ctx, req.Name, req.NoTelp, id)
	// if err != nil {
	// 	trx.Rollback()
	// 	return
	// }

	// trx.Commit()

	return
}

func (repo *contactgormhttp) Delete(id int) (err error) {
	// defer repo.db.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// trx, err := repo.db.BeginTx(ctx, nil)
	// if err != nil {
	// 	return
	// }

	// query := `DELETE FROM contact WHERE id = ?`
	// stmt, err := repo.db.PrepareContext(ctx, query)
	// if err != nil {
	// 	log.Print("prepare delete")
	// 	return
	// }

	// _, err = stmt.ExecContext(ctx, id)
	// if err != nil {
	// 	trx.Rollback()
	// 	return
	// }

	// trx.Commit()

	return
}
