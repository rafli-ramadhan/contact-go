package handler

import (
	"bytes"
	"contact-go/mocks"
	"contact-go/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContactHTTPHandler(t *testing.T) {
	t.Run("test get list contact success", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		ucMock.On("List").Return(model.ContactResponse{
			Status:  http.StatusOK,
			Message: "Ok",
			Data:    []model.Contact{
				{
					Id: 1,
					Name: "Andi",
					NoTelp: "08987534895",
				},
				{
					Id: 2,
					Name: "Umar",
					NoTelp: "08987534895",
				},
			},
		}, nil)

		request := httptest.NewRequest("GET", "http://localhost:5000/contact", nil)
		recorder := httptest.NewRecorder()
		handler.List(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("test get list contact failed", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		ucMock.On("List").Return(model.ContactResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		}, fmt.Errorf("Get List Error"))

		request := httptest.NewRequest("GET", "http://localhost:5000/contact", nil)
		recorder := httptest.NewRecorder()
		handler.List(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})

	t.Run("test create new contact success", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		req := []model.ContactRequest{
			{
				Name: "Andi",
				NoTelp: "08987534895",
			},
			{
				Name: "Umar",
				NoTelp: "08987534895",
			},
		}
		ucMock.On("Add", req).Return(model.ContactResponse{
			Status:  http.StatusCreated,
			Message: "Created",
			Data:    []model.Contact{
				{
					Id: 1,
					Name: "Andi",
					NoTelp: "08987534895",
				},
				{
					Id: 2,
					Name: "Umar",
					NoTelp: "08987534895",
				},
			},
		}, nil)

		jsonData, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonData)

		request := httptest.NewRequest("POST", "http://localhost:5000/contact", body)
		recorder := httptest.NewRecorder()
		handler.Add(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusCreated, response.StatusCode)
	})

	t.Run("test create new contact failed 1", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		req := []model.ContactRequest{
			{
				Name: "",
				NoTelp: "08987534895",
			},
			{
				Name: "",
				NoTelp: "08987534895",
			},
		}
		ucMock.On("Add", req).Return(model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "Name should not be empty",
			Data:    nil,
		}, errors.New("name should not be empty"))

		jsonData, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonData)

		request := httptest.NewRequest("POST", "http://localhost:5000/contact", body)
		recorder := httptest.NewRecorder()
		handler.Add(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("test create new contact failed 2", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		req := []model.ContactRequest{
			{
				Name: "Andi",
				NoTelp: "",
			},
			{
				Name: "Umar",
				NoTelp: "",
			},
		}
		ucMock.On("Add", req).Return(model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "No telp should not be empty",
			Data:    nil,
		}, errors.New("no telp should not be empty"))

		jsonData, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonData)

		request := httptest.NewRequest("POST", "http://localhost:5000/contact", body)
		recorder := httptest.NewRecorder()
		handler.Add(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	
	t.Run("test create new contact failed 3", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		req := []model.ContactRequest{
			{
				Name: "Andi",
				NoTelp: "",
			},
			{
				Name: "Umar",
				NoTelp: "",
			},
		}
		ucMock.On("Add", req).Return(model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "No telp should not be empty",
			Data:    nil,
		}, errors.New("no telp should not be empty"))

		reqBody := `{
			{
				Name: "Andi,
				NoTelp: "",
			},
			{
				Name: "Umar,
				NoTelp: "",
			},
		}`
		body := bytes.NewReader([]byte(reqBody))

		request := httptest.NewRequest("POST", "http://localhost:5000/contact", body)
		recorder := httptest.NewRecorder()
		handler.Add(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("test update contact success", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		req := model.ContactRequest{
			Name: "Andi",
			NoTelp: "08987534895",
		}

		ucMock.On("Update", "1", req).Return(model.ContactResponse{
			Status:  http.StatusOK,
			Message: "Updated",
			Data:    nil,
		}, nil)

		//id := "1"
		jsonData, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonData)

		//URL := fmt.Sprintf("http://localhost:5000/contact/%s", id)
		request := httptest.NewRequest("PATCH", "http://localhost:5000/contact?id=1", body)
		recorder := httptest.NewRecorder()
		handler.Update(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("test update contact failed 1", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		req := model.ContactRequest{
			Name: "",
			NoTelp: "08987534895",
		}

		ucMock.On("Update", "1", req).Return(model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "Id should not be more than 0",
			Data:    nil,
		}, errors.New("id should not be more than 0"))

		//id := "1"
		jsonData, err := json.Marshal(req)
		if err != nil {
			t.Error(err)
		}
		body := bytes.NewReader(jsonData)

		//URL := fmt.Sprintf("http://localhost:5000/contact/%s", id)
		request := httptest.NewRequest("PATCH", "http://localhost:5000/contact?id=1", body)
		recorder := httptest.NewRecorder()
		handler.Update(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("test update contact failed 2", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		req := model.ContactRequest{
			Name: "",
			NoTelp: "08987534895",
		}

		ucMock.On("Update", "1", req).Return(model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "Id should not be more than 0",
			Data:    nil,
		}, errors.New("id should not be more than 0"))

		//id := "1"
		reqBody := `{
			{
				Name: "Andi,
				NoTelp: "",
			},
			{
				Name: "Umar,
				NoTelp: "",
			},
		}`
		body := bytes.NewReader([]byte(reqBody))

		//URL := fmt.Sprintf("http://localhost:5000/contact/%s", id)
		request := httptest.NewRequest("PATCH", "http://localhost:5000/contact?id=1", body)
		recorder := httptest.NewRecorder()
		handler.Update(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("test delete contact success", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		ucMock.On("Delete", "1").Return(model.ContactResponse{
			Status:  http.StatusOK,
			Message: "Deleted",
			Data:    nil,
		}, nil)

		request := httptest.NewRequest("DELETE", "http://localhost:5000/contact?id=1", nil)
		recorder := httptest.NewRecorder()
		handler.Delete(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("test delete contact failed", func(t *testing.T) {
		ucMock := mocks.NewUseCaseMock()
		handler := NewContactHttpDbHandler(ucMock)

		ucMock.On("Delete", "-1").Return(model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "Id should not be more than 0",
			Data:    nil,
		}, errors.New("id should not be more than 0"))

		request := httptest.NewRequest("DELETE", "http://localhost:5000/contact?id=-1", nil)
		recorder := httptest.NewRecorder()
		handler.Delete(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
}
