package repository

import (
	"errors"

	model "contact-go/model"
)

type repository struct {}

func NewContactRepository () *repository {
	return &repository{}
}

func (repo *repository) getLastID() (lastID int, err error) {
	list := repo.List()

	if len(list) == 0 {
		lastID = 0
	} else {
		for _, v := range list {
			if lastID < int(v.Id) {
				lastID = int(v.Id)
			}
		}
	}

	return
}

func (repo *repository) GetIndexById(id int) (index int, value model.Contact, err error) {
	for i, v := range model.ContactSlice {
		if v.Id == id {
			index = int(i)
			value = v
			return index, value, nil
		}
	}
	return -1, model.Contact{}, errors.New("id not found")
}

func (repo *repository) List() (result []model.Contact) {
	return model.ContactSlice
}

func (repo *repository) Add(req model.ContactRequest) (err error) {
	lastID, err := repo.getLastID()
	if err != nil {
		return
	}

	contact := model.Contact{
		Id:     lastID + 1,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}
	model.ContactSlice = append(model.ContactSlice, contact)
	return
}

func (repo *repository) Update(id int, req model.ContactRequest) (err error) {
	index, value, err := repo.GetIndexById(id)
	if err != nil {
		return
	}

	if req.Name == "" {
		req.Name = value.Name
	}

	if req.NoTelp == "" {
		req.NoTelp = value.NoTelp
	}

	model.ContactSlice[index] = model.Contact{
		Id:     value.Id,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}

	return
}

func (repo *repository) Delete(id int) (err error) {
	index, _, err := repo.GetIndexById(id)
	if err != nil {
		return
	}

	deletedItemIndex := index
	model.ContactSlice = append(model.ContactSlice[:deletedItemIndex], model.ContactSlice[deletedItemIndex+1:]...)
	return
}