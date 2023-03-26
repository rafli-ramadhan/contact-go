package repository

import (
	"encoding/json"
	"errors"
	"os"

	model "contact-go/model"
)

type repository struct {}

func NewContactRepository () *repository {
	return &repository{}
}

func (repo *repository) getLastID() (lastID int, err error) {
	list, err := repo.List()
	if err != nil {
		return
	}

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
	list, err := repo.List()
	if err != nil {
		return
	}

	for i, v := range list {
		if v.Id == id {
			index = int(i)
			value = v
			return index, value, nil
		}
	}
	return -1, model.Contact{}, errors.New("id not found")
}

func (repo *repository) List() (result []model.Contact, err error) {
	var list []model.Contact
	// JSON -> struct
	reader, err := os.Open("contact.json")
	if err != nil {
		return []model.Contact{}, err
	}
	decoder := json.NewDecoder(reader)
	decoder.Decode(&list)

	return list, nil
}

func (repo *repository) UpdateJSON(list []model.Contact) (err error) {
	// struct -> JSON
	write, err := os.Create("contact.json")
	if err != nil {
		return
	}
	encoder := json.NewEncoder(write)
	encoder.Encode(list)
	return
}

func (repo *repository) Add(req model.ContactRequest) (err error) {
	// JSON to struct
	list, err := repo.List()
	if err != nil {
		return
	}

	lastID, err := repo.getLastID()
	if err != nil {
		return
	}

	contact := model.Contact{
		Id:     lastID + 1,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}
	list = append(list, contact)

	err = repo.UpdateJSON(list)
	if err != nil {
		return
	}
	return
}

func (repo *repository) Update(id int, req model.ContactRequest) (err error) {
	index, value, err := repo.GetIndexById(id)
	if err != nil {
		return
	}

	list, err := repo.List()
	if err != nil {
		return
	}

	if req.Name == "" {
		req.Name = value.Name
	}

	if req.NoTelp == "" {
		req.NoTelp = value.NoTelp
	}

	list[index] = model.Contact{
		Id:     value.Id,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}

	err = repo.UpdateJSON(list)
	if err != nil {
		return
	}
	return
}

func (repo *repository) Delete(id int) (err error) {
	list, err := repo.List()
	if err != nil {
		return
	}

	index, _, err := repo.GetIndexById(id)
	if err != nil {
		return
	}

	deletedItemIndex := index
	list = append(list[:deletedItemIndex], list[deletedItemIndex+1:]...)

	err = repo.UpdateJSON(list)
	if err != nil {
		return
	}
	return
}