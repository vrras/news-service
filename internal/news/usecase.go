package news

import (
	"context"
)

type UseCase interface {
	FindAll(context context.Context, topicKeyword string, statusKeyword string) ([]News, error)
	FindByID(context context.Context, id int) (News, error)
	Add(context context.Context, news News) (News, error)
	Update(context context.Context, editedNews News) (News, error)
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
func (us *useCase) FindAll(context context.Context, topicKeyword string, statusKeyword string) (res []News, err error) {
	res, err = us.repo.GetBySpec(context, FilterSpec{Topic: topicKeyword, Status: statusKeyword})

	return res, err
}

//FindById ...
func (us *useCase) FindByID(context context.Context, id int) (res News, err error) {
	res, err = us.repo.GetByID(context, id)

	return res, err
}

//Add ...
func (us *useCase) Add(context context.Context, newNews News) (res News, err error) {
	res, err = us.repo.Upsert(context, newNews)

	return res, err
}

//Update ...
func (us *useCase) Update(context context.Context, editedNews News) (res News, err error) {
	res, err = us.repo.Upsert(context, editedNews)

	return res, err
}

//Remove ...
func (us *useCase) Delete(context context.Context, id int) (err error) {
	err = us.repo.DeleteByID(context, id)

	return err
}
