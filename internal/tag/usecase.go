package tag

import (
	"context"
)

type UseCase interface {
	FindAll(context context.Context) ([]Tag, error)
	FindByID(context context.Context, id int) (Tag, error)
	Add(context context.Context, tag Tag) (Tag, error)
	Update(context context.Context, editedTag Tag) (Tag, error)
	Delete(context context.Context, id int) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

//FindAll ...
func (us *useCase) FindAll(context context.Context) (res []Tag, err error) {
	res, err = us.repo.GetAll(context)

	return res, err
}

//FindById ...
func (us *useCase) FindByID(context context.Context, id int) (res Tag, err error) {
	res, err = us.repo.GetByID(context, id)

	return res, err
}

//Add ...
func (us *useCase) Add(context context.Context, newTag Tag) (res Tag, err error) {
	res, err = us.repo.Upsert(context, newTag)

	return res, err
}

//Update ...
func (us *useCase) Update(context context.Context, editedTag Tag) (res Tag, err error) {
	res, err = us.repo.Upsert(context, editedTag)

	return res, err
}

//Remove ...
func (us *useCase) Delete(context context.Context, id int) (err error) {
	err = us.repo.DeleteByID(context, id)

	return err
}
