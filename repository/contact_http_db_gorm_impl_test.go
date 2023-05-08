package repository

import (
	"contact-go/model"
	"errors"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

type GormMysqlclient struct {
    suite.Suite
	mock sqlmock.Sqlmock
	repo ContactRepositorier
	sqlDB *sql.DB
}

// set up sql mock
func (client *GormMysqlclient) SetupTest() {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error database connection %s", err))
	}

	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DriverName:         	   "mysql",
				Conn:               	   dbMock,
				SkipInitializeWithVersion: true,
			},
		), &gorm.Config{
			Logger: 				logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,
			PrepareStmt:			true,
		},
	)
    if err != nil {
		panic(fmt.Sprintf("Error database gorm connection %s", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		client.Require().NoError(err)
	}

	client.mock  = mock
	client.repo  = NewContactGormHTTPRepository(db)
	client.sqlDB = sqlDB
}

func (client *GormMysqlclient) TearDownTest() {
	defer client.sqlDB.Close()
	log.Println("After Each Test Executed")
}

func (client *GormMysqlclient) SetupSuite() {
	log.Println("Setup Before All Test Executed")
}

func (client *GormMysqlclient) TearDownSuite() {
	log.Println("After All Test Executed")
}

func (client *GormMysqlclient) AfterTest() {
	log.Println("After Test Executed")
}

func (client *GormMysqlclient) TestGetListContactSuccess() {
	// data dog
	row := sqlmock.NewRows([]string{"id","name","no_telp"}).AddRow(1, "Andi", "0834234235244").AddRow(2, "Umar", "0894339843943")
	client.mock.ExpectPrepare("SELECT `contacts`.`id`,`contacts`.`name`,`contacts`.`no_telp` FROM `contacts`").
		WillBeClosed().
		ExpectQuery().
		WillReturnRows(row)

	list_contact, err := client.repo.List()
	if err != nil {
		client.T().Errorf("error get list contact: %s", err)
	}

	require.NoError(client.T(), err)
	require.NotEmpty(client.T(), list_contact)
}

func (client *GormMysqlclient) TestGetListContactFailed1() {
	// data dog
	client.mock.ExpectPrepare("SELECT `contacts`.`id`,`contacts`.`name`,`contacts`.`no_telp` FROM `contacts`").
		WillBeClosed().
		WillReturnError(fmt.Errorf("some error"))

	list_contact, err := client.repo.List()	
	if err != nil {
		log.Printf("error get list : %v", err)
	}

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *GormMysqlclient) TestGetListContactFailed2() {
	// data dog
	client.mock.ExpectPrepare("SELECT `contacts`.`id`,`contacts`.`name`,`contacts`.`no_telp` FROM `contacts`").
		WillBeClosed().
		ExpectQuery().
		WillReturnError(fmt.Errorf("some error"))

	list_contact, err := client.repo.List()
	if err != nil {
		log.Printf("error get list : %v", err)
	}

	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *GormMysqlclient) TestAddContactSuccess() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO `contacts` (`name`,`no_telp`) VALUES (?,?)")).WillBeClosed().
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

func (client *GormMysqlclient) TestAddContactFailed1() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO `contacts` (`name`,`no_telp`) VALUES (?,?)")).
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

func (client *GormMysqlclient) TestAddContactFailed2() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO `contacts` (`name`,`no_telp`) VALUES (?,?)")).
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

// func (client *GormMysqlclient) TestUpdateContactSuccess() {
// 	// data dog
// 	client.mock.ExpectBegin()
// 	client.mock.ExpectExec(regexp.QuoteMeta("UPDATE `contacts` SET `name`=?,`no_telp`=? WHERE `id` = ?")).
// 		WithArgs("Andi", "0884275327327", 1).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	client.mock.ExpectCommit()

// 	id := 1
// 	req := model.ContactRequest{
// 		Name:   "Andi",
// 		NoTelp: "0884275327327",
// 	}
// 	err := client.repo.Update(id, req)
// 	if err != nil {
// 		client.T().Errorf("error update contact: %s", err)
// 	}

// 	require.NoError(client.T(), err)
// }

func (client *GormMysqlclient) TestUpdateContactFailed1() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("UPDATE `contacts` SET `name`=?,`no_telp`=? WHERE `id` = ?")).
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

func (client *GormMysqlclient) TestUpdateContactFailed2() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("UPDATE `contacts` SET `name`=?,`no_telp`=? WHERE `id` = ?")).
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

// func (client *GormMysqlclient) TestDeleteContactSuccess() {
// 	// data dog
// 	client.mock.ExpectBegin()
// 	client.mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM `contacts` WHERE `id` = ?")).
// 		ExpectExec().
// 		WithArgs(1).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	client.mock.ExpectCommit()

// 	id := 1
// 	err := client.repo.Delete(id)
// 	if err != nil {
// 		client.T().Errorf("error delete contact: %s", err)
// 	}

// 	require.NoError(client.T(), err)
// }

func (client *GormMysqlclient) TestDeleteContactFailed1() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM `contacts` WHERE `id` = ?")).
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

func (client *GormMysqlclient) TestDeleteContactFailed2() {
	// data dog
	client.mock.ExpectBegin()
	client.mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM `contacts` WHERE `id` = ?")).
		WillReturnError(errors.New("some error"))
	client.mock.ExpectRollback()

	id := 1
	err := client.repo.Delete(id)
	if err != nil {
		log.Printf("error delete : %v", err)
	}

	require.Error(client.T(), err)
}

func TestRepoGormHTTP(t *testing.T) {
	suite.Run(t, new(GormMysqlclient))
}