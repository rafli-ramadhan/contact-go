package repository

import (
	"testing"

	"contact-go/model"
	"github.com/stretchr/testify/require"
)

func TestRepoJson(t *testing.T) {
	t.Run("test get last id failed", func(t *testing.T) {
		repo := NewContactJsonRepository("")

		_, err := repo.getLastID()
		require.Error(t, err)
	})

	t.Run("test get list contact success", func(t *testing.T) {
		repo := NewContactJsonRepository("../data/contact.json")

		res, err := repo.List()
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test get list contact failed", func(t *testing.T) {
		repo := NewContactJsonRepository("")

		res, err := repo.List()
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test create new contact success", func(t *testing.T) {
		repo := NewContactJsonRepository("../data/contact.json")

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

		res, err := repo.Add(req)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("test create new contact failed", func(t *testing.T) {
		repo := NewContactJsonRepository("")

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

		res, err := repo.Add(req)
		require.Error(t, err)
		require.Empty(t, res)
	})

	t.Run("test update contact success 1", func(t *testing.T) {
		repo := NewContactJsonRepository("../data/contact.json")

		id := 21
		req := model.ContactRequest{
			Name: "",
			NoTelp: "",
		}

		err := repo.Update(id, req)
		require.NoError(t, err)
	})

	t.Run("test update contact success 2", func(t *testing.T) {
		repo := NewContactJsonRepository("../data/contact.json")

		id := 21
		req := model.ContactRequest{
			Name: "Ardi",
			NoTelp: "082828329292",
		}

		err := repo.Update(id, req)
		require.NoError(t, err)
	})

	t.Run("test update contact failed 1", func(t *testing.T) {
		repo := NewContactJsonRepository("../data/contact.json")

		id := 25
		req := model.ContactRequest{
			Name: "Ardi",
			NoTelp: "082828329292",
		}

		err := repo.Update(id, req)
		require.Error(t, err)
	})

	t.Run("test update contact failed 2", func(t *testing.T) {
		repo := NewContactJsonRepository("")

		id := 21
		req := model.ContactRequest{
			Name: "Ardi",
			NoTelp: "082828329292",
		}

		err := repo.Update(id, req)
		require.Error(t, err)
	})

	t.Run("test delete contact success", func(t *testing.T) {
		repo := NewContactJsonRepository("../data/contact.json")

		id := 22
		err := repo.Delete(id)
		require.NoError(t, err)

		id = 23
		err = repo.Delete(id)
		require.NoError(t, err)
	})

	t.Run("test delete contact failed 1", func(t *testing.T) {
		repo := NewContactJsonRepository("../data/contact.json")

		id := 27
		err := repo.Delete(id)
		require.Error(t, err)
	})

	t.Run("test delete contact failed 2", func(t *testing.T) {
		repo := NewContactJsonRepository("")

		id := 27
		err := repo.Delete(id)
		require.Error(t, err)
	})
}