package mocks

import (
	"contact-go/model"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func NewRepoMock() *RepoMock {
	return &RepoMock{}
}

func (m *RepoMock) List() (result []model.Contact, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called()
	// mengembalikan parameter output berdasarkan index
	result = ret.Get(0).([]model.Contact)
	if ret.Get(1) != nil {
		// type assertion -> mengubah interface kosong menjadi suatu tipe data yang diperlukan
		err = ret.Get(1).(error)
	}
	return result, err
}

func (m *RepoMock) Add(req []model.ContactRequest) (result []model.Contact, err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(req)
	// mengembalikan parameter output berdasarkan index
	result = ret.Get(0).([]model.Contact)
	if ret.Get(1) != nil {
		// type assertion -> mengubah interface kosong menjadi suatu tipe data yang diperlukan
		err = ret.Get(1).(error)
	}
	return result, err
}

func (m *RepoMock) Update(id int, req model.ContactRequest) (err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(id, req)
	// mengembalikan parameter output berdasarkan index
	if ret.Get(0) != nil {
		// type assertion -> mengubah interface kosong menjadi suatu tipe data yang diperlukan
		err = ret.Get(0).(error)
	}
	return err
}

func (m *RepoMock) Delete(id int) (err error) {
	// sebagai indikator parameter diperoleh
	ret := m.Called(id)
	// mengembalikan parameter output berdasarkan index
	if ret.Get(0) != nil {
		// type assertion -> mengubah interface kosong menjadi suatu tipe data yang diperlukan
		err = ret.Get(0).(error)
	}
	return err
}