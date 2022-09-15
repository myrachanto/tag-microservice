package tag

import (
	httperrors "github.com/myrachanto/erroring"
)

var (
	TagService TagServiceInterface = &tagService{}
)

type TagServiceInterface interface {
	Create(seo *Tag) (*Tag, httperrors.HttpErr)
	GetOne(id string) (*Tag, httperrors.HttpErr)
	GetAll() ([]*Tag, httperrors.HttpErr)
	Featured(code string, status bool) httperrors.HttpErr
	Update(id string, tag *Tag) (*Tag, httperrors.HttpErr)
	Delete(id string) (string, httperrors.HttpErr)
}
type tagService struct {
	repo TagRepoInterface
}

func NewtagService(repository TagRepoInterface) TagServiceInterface {
	return &tagService{
		repository,
	}
}

func (service *tagService) Create(tag *Tag) (*Tag, httperrors.HttpErr) {
	tag, err1 := service.repo.Create(tag)
	return tag, err1

}

func (service *tagService) Featured(code string, status bool) httperrors.HttpErr {
	err1 := service.repo.Featured(code, status)
	return err1
}
func (service *tagService) GetOne(id string) (*Tag, httperrors.HttpErr) {
	tag, err1 := service.repo.GetOne(id)
	return tag, err1
}

func (service *tagService) GetAll() ([]*Tag, httperrors.HttpErr) {
	tags, err := service.repo.GetAll()
	return tags, err
}

func (service *tagService) Update(id string, tag *Tag) (*Tag, httperrors.HttpErr) {
	tag, err1 := service.repo.Update(id, tag)
	return tag, err1
}
func (service *tagService) Delete(id string) (string, httperrors.HttpErr) {
	success, failure := service.repo.Delete(id)
	return success, failure
}
