package mocks

import (
	"contact-go/model"
	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	mock.Mock
}

func NewUseCaseMock() *UsecaseMock {
	return &UsecaseMock{}
}

func (m *UsecaseMock) List() (model.ContactResponse, error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called()
	result := ret.Get(0).(model.ContactResponse)
	var err error
	if ret.Get(1) != nil {
		// type assertion -> mengubah interface kosong menjadi suatu tipe data yang diperlukan
		err = ret.Get(1).(error)
	}
	return result, err
}

func (m *UsecaseMock) Add(req []model.ContactRequest) (model.ContactResponse, error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(req)
	result := ret.Get(0).(model.ContactResponse)
	var err error
	if ret.Get(1) != nil {
		// type assertion -> mengubah interface kosong menjadi suatu tipe data yang diperlukan
		err = ret.Get(1).(error)
	}
	return result, err
}

func (m *UsecaseMock) Update(idStr string, req model.ContactRequest) (model.ContactResponse, error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(idStr, req)
	result := ret.Get(0).(model.ContactResponse)
	var err error
	if ret.Get(1) != nil {
		// type assertion -> mengubah interface kosong menjadi suatu tipe data yang diperlukan
		err = ret.Get(1).(error)
	}
	return result, err
}

func (m *UsecaseMock) Delete(idStr string) (model.ContactResponse, error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(idStr)
	result := ret.Get(0).(model.ContactResponse)
	var err error
	if ret.Get(1) != nil {
		// type assertion -> mengubah interface kosong menjadi suatu tipe data yang diperlukan
		err = ret.Get(1).(error)
	}
	return result, err
}