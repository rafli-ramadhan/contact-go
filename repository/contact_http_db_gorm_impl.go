package repository

import (
	"contact-go/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type contactgormhttp struct {
	db *gorm.DB
}

func NewContactGormHTTPRepository(client *gorm.DB) *contactgormhttp {
	return &contactgormhttp{
		db: client,
	}
}

type contact struct {
	Name   string `gorm:"column:name" json:"name"`
	NoTelp string `gorm:"column:no_telp" json:"no_telp"`
}

func (repo *contactgormhttp) List() (result []model.Contact, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := repo.db.Model(&contact{}).WithContext(ctx).Find(&result)
	err = query.Error
	return
}

func (repo *contactgormhttp) Add(req []model.ContactRequest) (result []model.Contact, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newContacts []model.Contact

	for _, contact := range req {
		newContacts = append(newContacts, model.Contact{
			Name:   contact.Name,
			NoTelp: contact.NoTelp,
		})
	}
	query := repo.db.Model(&model.Contact{}).WithContext(ctx).
		Create(&newContacts)

	if err = query.Error; err != nil {
		return result, err
	}

	return newContacts, err
}

func (repo *contactgormhttp) Update(id int, req model.ContactRequest) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := repo.db.Model(&contact{}).WithContext(ctx).
		Where("id", id).
		Updates(req)

	if err = query.Error; err != nil {
		return
	}

	return
}

func (repo *contactgormhttp) Delete(id int) (err error) {
	query := repo.db.Model(&contact{}).
		Where("id", id).
		Delete(contact{})

	if err = query.Error; err != nil {
		return
	}

	return
}
