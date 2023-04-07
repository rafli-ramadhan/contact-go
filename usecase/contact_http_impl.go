package usecase

import (
	"contact-go/model"
	"contact-go/repository"
	"errors"
	"net/http"
)

type usecase struct {
	repo repository.ContactRepositorier
}

func NewUseCase(repository repository.ContactRepositorier) *usecase {
	return &usecase{
		repo: repository,
	}
}

func (uc *usecase) List() (model.ContactResponse, error) {
	res, err := uc.repo.List()
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data: 	 nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data: 	 res,
	}, nil
}

func (uc *usecase) Add(req []model.ContactRequest) (model.ContactResponse, error) {
	for _, v := range req {
		if v.Name == "" || v.NoTelp == "" {
			return model.ContactResponse{
				Status:  http.StatusBadRequest,
				Message: "Bad request",
				Data: 	 nil,
			}, errors.New("name or no telp should not be empty")
		}
	}

	res, err := uc.repo.Add(req)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data: 	 nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusCreated,
		Message: "Created",
		Data: 	 res,
	}, nil
}

func (uc *usecase) Update(id int, req model.ContactRequest) (model.ContactResponse, error) {
	if req.Name == "" || req.NoTelp == "" {
		return model.ContactResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad request",
			Data: 	 nil,
		}, errors.New("name or no telp should not be empty")
	}

	err := uc.repo.Update(id, req)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data: 	 nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "Updated",
		Data: 	 nil,
	}, nil
}

func (uc *usecase) Delete(id int) (model.ContactResponse, error) {
	err := uc.repo.Delete(id)
	if err != nil {
		return model.ContactResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Data: 	 nil,
		}, err
	}
	return model.ContactResponse{
		Status:  http.StatusOK,
		Message: "Deleted",
		Data: 	 nil,
	}, nil
}
