package usecase

import (
	"fmt"
	"net/http"
	"testing"

	"contact-go/mocks"
	"contact-go/model"

	"github.com/stretchr/testify/require"
)

func TestUseCaseHTTP(t *testing.T) {
	t.Run("is valid id success", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		idStr := "2"
		id, res, err := uc.IsValidID(idStr)
		require.NoError(t, err)
		require.NotEmpty(t, id)
		require.NotEqual(t, http.StatusBadRequest, res)
	})

	t.Run("is valid id failed", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		idStr := "a"
		_, res, err := uc.IsValidID(idStr)
		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
	})

	t.Run("is valid name and no telp success", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		name := "Andi"
		noTelp := "0867576436254"
		res, err := uc.IsValidNameAndNoTelp(name, noTelp)
		require.NoError(t, err)
		require.NotEqual(t, http.StatusBadRequest, res)
	})

	t.Run("is valid name and no telp failed", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		name := "Andi"
		noTelp := ""
		res, err := uc.IsValidNameAndNoTelp(name, noTelp)
		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
	})

	t.Run("get-list-success", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		repoMock.On("List").Return([]model.Contact{
			{
				Id: 1,
				Name: "Ardi",
				NoTelp: "082828329292",
			},
			{
				Id: 2,
				Name: "Amar",
				NoTelp: "082828329292",
			},
		}, nil)
		
		res, err := uc.List()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "Ok", res.Message)
	})

	t.Run("get-list-failed", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		repoMock.On("List").Return([]model.Contact{}, fmt.Errorf("some error"))
		
		res, err := uc.List()
		require.Error(t, err)
		require.Equal(t, http.StatusInternalServerError, res.Status)
	})

	t.Run("add-contact-success", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		req := []model.ContactRequest{
			{
				Name: "Ardi",
				NoTelp: "082828329292",
			},
			{
				Name: "Amar",
				NoTelp: "082828329292",
			},
		}
	
		repoMock.On("Add", req).Return([]model.Contact{
			{
				Id: 1,
				Name: "Ardi",
				NoTelp: "082828329292",
			},
			{
				Id: 2,
				Name: "Amar",
				NoTelp: "082828329292",
			},
		}, nil)

		res, err := uc.Add(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.Status)
		require.Equal(t, "Created", res.Message)
	})

	t.Run("add-contact-failed-1", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		req := []model.ContactRequest{
			{
				Name: "Ardi",
				NoTelp: "082828329292",
			},
			{
				Name: "Amar",
				NoTelp: "082828329292",
			},
		}
	
		repoMock.On("Add", req).Return([]model.Contact{
			{
				Id: 1,
				Name: "Ardi",
				NoTelp: "082828329292",
			},
			{
				Id: 2,
				Name: "Amar",
				NoTelp: "082828329292",
			},
		}, fmt.Errorf("some error"))

		res, err := uc.Add([]model.ContactRequest{
			{
				Name: "",
				NoTelp: "082828329292",
			},
			{
				Name: "Amar",
				NoTelp: "082828329292",
			},
		})
		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
	})

	t.Run("add-contact-failed-2", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		req := []model.ContactRequest{
			{
				Name: "Ardi",
				NoTelp: "082828329292",
			},
			{
				Name: "Amar",
				NoTelp: "082828329292",
			},
		}
	
		repoMock.On("Add", req).Return([]model.Contact{
			{
				Id: 1,
				Name: "Ardi",
				NoTelp: "082828329292",
			},
			{
				Id: 2,
				Name: "Amar",
				NoTelp: "082828329292",
			},
		}, fmt.Errorf("some error"))

		res, err := uc.Add([]model.ContactRequest{
			{
				Name: "Ardi",
				NoTelp: "082828329292",
			},
			{
				Name: "Amar",
				NoTelp: "082828329292",
			},
		})
		require.Error(t, err)
		require.Equal(t, http.StatusInternalServerError, res.Status)
	})

	t.Run("update-contact-success", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		id := 2
		req := model.ContactRequest{
			Name: "Ardi",
			NoTelp: "082828329292",
		}
		repoMock.On("Update", id, req).Return(nil)

		idStr := "2"
		res, err := uc.Update(idStr, req)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "Updated", res.Message)
	})

	t.Run("update-contact-fail-1", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		id := 2
		req := model.ContactRequest{
			Name: "Ardi",
			NoTelp: "082828329292",
		}
		repoMock.On("Update", id, req).Return(nil)

		idStr := "-2"
		res, err := uc.Update(idStr, model.ContactRequest{
			Name: "Andi",
			NoTelp: "082828329292",
		})

		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
	})

	t.Run("update-contact-fail-2", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		id := 2
		req := model.ContactRequest{
			Name: "Ardi",
			NoTelp: "082828329292",
		}
		repoMock.On("Update", id, req).Return(nil)

		idStr := "2"
		res, err := uc.Update(idStr, model.ContactRequest{
			Name: "",
			NoTelp: "082828329292",
		})

		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
	})

	t.Run("update-contact-fail-3", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		id := 2
		req := model.ContactRequest{
			Name: "Andi",
			NoTelp: "082828329292",
		}
		repoMock.On("Update", id, req).Return(fmt.Errorf("some error"))

		idStr := "2"
		res, err := uc.Update(idStr, model.ContactRequest{
			Name: "Andi",
			NoTelp: "082828329292",
		})

		require.Error(t, err)
		require.Equal(t, http.StatusInternalServerError, res.Status)
	})

	t.Run("delete-contact-success", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		id := 2
		repoMock.On("Delete", id).Return(nil)

		idStr := "2"
		res, err := uc.Delete(idStr)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
		require.Equal(t, "Deleted", res.Message)
	})

	t.Run("delete-contact-failed-1", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		id := 2
		repoMock.On("Delete", id).Return(nil)

		idStr := "-2"
		res, err := uc.Delete(idStr)
		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, res.Status)
	})

	t.Run("delete-contact-failed-2", func(t *testing.T) {
		repoMock := mocks.NewRepoMock()
		uc := NewUseCase(repoMock)

		id := 2
		repoMock.On("Delete", id).Return(fmt.Errorf("some error"))

		idStr := "2"
		res, err := uc.Delete(idStr)
		require.Error(t, err)
		require.Equal(t, http.StatusInternalServerError, res.Status)
	})
}