package repository

import (
	"contact-go/model"
	"regexp"

	"errors"
	"fmt"
	"log"
	"testing"

	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormMysqlclient struct {
	suite.Suite
	mock   sqlmock.Sqlmock
	repo   ContactRepositorier
	sqlDB  *sql.DB
	gormDB *gorm.DB
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
				Conn:                      dbMock,
				SkipInitializeWithVersion: true,
			},
		), &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Info),
			// SkipDefaultTransaction: true,
			// PrepareStmt: true,
		},
	)
	if err != nil {
		panic(fmt.Sprintf("Error database gorm connection %s", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		client.Require().NoError(err)
	}

	client.mock = mock
	client.repo = NewContactGormHTTPRepository(db)
	client.sqlDB = sqlDB
	client.gormDB = db
}

func (client *GormMysqlclient) BeforeTest() {
	log.Println("Before Test")
}

func (client *GormMysqlclient) TearDownTest() {
	stmtManager, ok := client.gormDB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		for _, stmt := range stmtManager.Stmts {
			stmt.Close()
		}
	}
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
	row := sqlmock.NewRows([]string{"id", "name", "no_telp"}).AddRow(1, "Andi", "0834234235244").AddRow(2, "Umar", "0894339843943")
	client.mock.ExpectQuery("SELECT `contacts`.`id`,`contacts`.`name`,`contacts`.`no_telp` FROM `contacts`").
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
	client.mock.ExpectQuery("SELECT `contacts`.`id`,`contacts`.`name`,`contacts`.`no_telp` FROM `contacts`").
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
	sql := regexp.QuoteMeta("INSERT INTO `contacts` (`name`,`no_telp`) VALUES (?,?)")
	client.mock.ExpectBegin()
	client.mock.ExpectExec(sql).
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
	fmt.Println(list_contact)
	if err != nil {
		client.T().Errorf("error get list contact: %s", err)
	}

	require.NoError(client.T(), err)
	require.NotEmpty(client.T(), list_contact)

	if err := client.mock.ExpectationsWereMet(); err != nil {
		client.T().Errorf("error expectation: %s", err)
	}
}

func (client *GormMysqlclient) TestAddContactFailed1() {
	// data dog
	sql := regexp.QuoteMeta("INSERT INTO `contacts` (`name`,`no_telp`) VALUES (?,?)")
	client.mock.ExpectExec(sql).
		WithArgs("Andi", "0884275327327").
		WillReturnError(errors.New("some error"))

	req := []model.ContactRequest{
		{
			Name:   "Andi",
			NoTelp: "0884275327327",
		},
	}
	list_contact, err := client.repo.Add(req)
	
	require.Error(client.T(), err)
	require.Empty(client.T(), list_contact)
}

func (client *GormMysqlclient) TestUpdateContactSuccess() {
	// data dog
	sql := regexp.QuoteMeta("UPDATE `contacts` SET `name`=?,`no_telp`=? WHERE `id` = ?")
	client.mock.ExpectBegin()
	client.mock.ExpectExec(sql).
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

func (client *GormMysqlclient) TestUpdateContactFailed1() {
	sql := regexp.QuoteMeta("UPDATE `contacts` SET `name`=?,`no_telp`=? WHERE `id` = ?")
	client.mock.ExpectExec(sql).
		WithArgs("Andi", "0884275327327", 1).
		WillReturnError(errors.New("some error"))

	id := 1
	req := model.ContactRequest{
		Name:   "Andi",
		NoTelp: "0884275327327",
	}
	err := client.repo.Update(id, req)

	require.Error(client.T(), err)
}

func (client *GormMysqlclient) TestDeleteContactSuccess() {
	// data dog
	sql := regexp.QuoteMeta("DELETE FROM `contacts` WHERE `id` = ?")
	client.mock.ExpectBegin()
	client.mock.ExpectExec(sql).
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

func TestRepoGormHTTP(t *testing.T) {
	suite.Run(t, new(GormMysqlclient))
}
