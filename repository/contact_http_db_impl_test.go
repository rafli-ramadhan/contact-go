package repository

import (
	"contact-go/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"testing"

	// "contact-go/helper/get-env"
	// "github.com/joho/godotenv"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Mysqlclient struct {
    suite.Suite
	db	 *sql.DB
	mock sqlmock.Sqlmock
	repo ContactRepositorier
}

// set up env
func (client *Mysqlclient) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err))
	}

	client.db = db
	client.mock = mock
	client.repo = NewContactHTTPRepository(db)
}

func (client *Mysqlclient) TearDownTest() {
	log.Println("After Each Test Executed")
}

func (client *Mysqlclient) SetupSuite() {
	log.Println("Setup Before All Test Executed")
}

func (client *Mysqlclient) TearDownSuite() {
	log.Println("After All Test Executed")
}

func (client *Mysqlclient) AfterTest() {
	log.Println("After Test Executed")
}

func (client *Mysqlclient) TestGetListContactSuccess() {
	// data dog
	row := sqlmock.NewRows([]string{"id","name","no_telp"}).AddRow(1, "Andi", "0834234235244").AddRow(2, "Umar", "0894339843943")
	client.mock.ExpectPrepare("SELECT id, name, no_telp FROM contact").WillBeClosed().ExpectQuery().WillReturnRows(row)

	list_contact, err := client.repo.List()
	if err != nil {
		client.T().Errorf("error get list contact: %s", err)
	}

	require.NoError(client.T(), err)
	require.NotEmpty(client.T(), list_contact)
}

func (client *Mysqlclient) TestGetListContactFailed1() {
	// data dog
	client.mock.ExpectPrepare("SELECT id, name, no_telp FROM contact").WillReturnError(fmt.Errorf("some error"))

	list_contact, err := client.repo.List()	
	if err != nil {
		log.Printf("error get list : %v", err)
	}

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestGetListContactFailed2() {
	// data dog
	client.mock.ExpectPrepare("SELECT id, name, no_telp FROM contact").WillBeClosed().ExpectQuery().WillReturnError(fmt.Errorf("some error"))

	list_contact, err := client.repo.List()
	if err != nil {
		log.Printf("error get list : %v", err)
	}

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContactSuccess() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO contact (name, no_telp) value (?,?)")).
		ExpectExec().
		WithArgs("Andi", "0884275327327").
		WillReturnResult(sqlmock.NewResult(1, 1))
	client.mock.ExpectCommit()

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}
	list_contact, err := client.repo.Add(req)
	if err != nil {
		client.T().Errorf("error get list contact: %s", err)
	}

	require.NoError(client.T(), err)
	require.NotEmpty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContactFailed1() {
	// data dog
	client.mock.ExpectBegin().WillReturnError(errors.New("some error"))

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}
	list_contact, err := client.repo.Add(req)
	if err != nil {
		log.Printf("error add : %v", err)
	}

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContactFailed2() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO contact (name, no_telp) value (?,?)")).
		ExpectExec().
		WithArgs("Andi", "0884275327327").
		WillReturnError(errors.New("some error"))
	client.mock.ExpectRollback()

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}
	list_contact, err := client.repo.Add(req)
	if err != nil {
		log.Printf("error add : %v", err)
	}

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContactFailed3() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO contact (name, no_telp) value (?,?)")).
		WillReturnError(errors.New("some error"))
	client.mock.ExpectRollback()

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}
	list_contact, err := client.repo.Add(req)
	if err != nil {
		log.Printf("error add : %v", err)
	}

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestAddContactFailed4() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO contact (name, no_telp) value (?,?)")).
		ExpectExec().
		WithArgs("Andi", "0884275327327").
		WillReturnResult(sqlmock.NewErrorResult(errors.New("last id error")))
	client.mock.ExpectRollback()

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}
	list_contact, err := client.repo.Add(req)
	if err != nil {
		log.Printf("error add : %v", err)
	}

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *Mysqlclient) TestUpdateContactSuccess() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("UPDATE contact SET name = ?, no_telp = ? WHERE id = ?")).
		ExpectExec().
		WithArgs("Andi", "0884275327327", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	client.mock.ExpectCommit()

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)
	if err != nil {
		client.T().Errorf("error update contact: %s", err)
	}

	require.NoError(client.T(), err)
}

func (client *Mysqlclient) TestUpdateContactFailed1() {
	// data dog
	client.mock.ExpectBegin().WillReturnError(errors.New("some error"))

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)
	if err != nil {
		log.Printf("error update : %v", err)
	}

	require.Error(client.T(), err)
}

func (client *Mysqlclient) TestUpdateContactFailed2() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("UPDATE contact SET name = ?, no_telp = ? WHERE id = ?")).
		ExpectExec().
		WithArgs("Andi", "0884275327327", 1).
		WillReturnError(errors.New("some error"))
	client.mock.ExpectRollback()

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)
	if err != nil {
		log.Printf("error update : %v", err)
	}

	require.Error(client.T(), err)
}

func (client *Mysqlclient) TestUpdateContactFailed3() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("UPDATE contact SET name = ?, no_telp = ? WHERE id = ?")).
		WillReturnError(errors.New("some error"))
	client.mock.ExpectRollback()

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)
	if err != nil {
		log.Printf("error update : %v", err)
	}

	require.Error(client.T(), err)
}

func (client *Mysqlclient) TestDeleteContactSuccess() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM contact WHERE id = ?")).
		ExpectExec().
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	client.mock.ExpectCommit()

	id := 1
	err := client.repo.Delete(id)
	if err != nil {
		client.T().Errorf("error delete contact: %s", err)
	}

	require.NoError(client.T(), err)
}

func (client *Mysqlclient) TestDeleteContactFailed1() {
	// data dog
	client.mock.ExpectBegin().WillReturnError(errors.New("some error"))

	id := 1
	err := client.repo.Delete(id)
	if err != nil {
		log.Printf("error delete : %v", err)
	}

	require.Error(client.T(), err)
}

func (client *Mysqlclient) TestDeleteContactFailed2() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM contact WHERE id = ?")).
		ExpectExec().
		WithArgs(1).
		WillReturnError(errors.New("some error"))
	client.mock.ExpectRollback()

	id := 1
	err := client.repo.Delete(id)
	if err != nil {
		log.Printf("error delete : %v", err)
	}

	require.Error(client.T(), err)
}

func (client *Mysqlclient) TestDeleteContactFailed3() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM contact WHERE id = ?")).
		WillReturnError(errors.New("some error"))
	client.mock.ExpectRollback()

	id := 1
	err := client.repo.Delete(id)
	if err != nil {
		log.Printf("error delete : %v", err)
	}

	require.Error(client.T(), err)
}

func TestRepoHTTP(t *testing.T) {
	suite.Run(t, new(Mysqlclient))
}