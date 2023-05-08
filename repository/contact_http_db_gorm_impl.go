package repository

import (
	"contact-go/model"
	"gorm.io/gorm"
)

type contactgormhttp struct{
	db *gorm.DB
}

func NewContactGormHTTPRepository(client *gorm.DB) *contactgormhttp {
	return &contactgormhttp{
		db: client,
	}
}

type contact struct {
	Name   string	`gorm:"column:name" json:"name"`
	NoTelp string	`gorm:"column:no_telp" json:"no_telp"`
}

func (repo *contactgormhttp) List() (result []model.Contact, err error) {
	query := repo.db.Model(&contact{}).Find(&result)
	err = query.Error
	return
}

func (repo *contactgormhttp) Add(req []model.ContactRequest) (result []model.Contact, err error) {
	query := repo.db.Model(&contact{}).
		Begin().
		Create(&req)

	err = query.Error
	if err != nil {
		query.Rollback()
		return
	}

	err = query.Commit().Error

	return
}

func (repo *contactgormhttp) Update(id int, req model.ContactRequest) (err error) {
	query := repo.db.Model(&contact{}).
		Begin().
		Where("id", id).
		Updates(req)

	err = query.Error
	if err != nil {
		query.Rollback()
		return
	}

	err = query.Commit().Error

	return
}

func (repo *contactgormhttp) Delete(id int) (err error) {
	query := repo.db.Model(&contact{}).
		Begin().
		Where("id", id).
		Delete(contact{})

	err = query.Error
	if err != nil {
		query.Rollback()
		return
	}

	err = query.Commit().Error
	query.Commit()

	return
}