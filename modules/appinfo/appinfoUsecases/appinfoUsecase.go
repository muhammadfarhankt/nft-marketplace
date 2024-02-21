package appinfoUsecases

import (
	"github.com/muhammadfarhankt/nft-marketplace/modules/appinfo"
	"github.com/muhammadfarhankt/nft-marketplace/modules/appinfo/appinfoRepositories"
)

type IAppinfoUsecase interface {
	FindCategory(req *appinfo.CategoryFilter) ([]*appinfo.Category, error)
	InsertCategory(category []*appinfo.Category) error
	DeleteCategory(categoryId string) error
}

type appinfoUsecase struct {
	appinfoRepository appinfoRepositories.IAppinfoRepository
}

func AppinfoUsecase(appinfoRepository appinfoRepositories.IAppinfoRepository) IAppinfoUsecase {
	return &appinfoUsecase{
		appinfoRepository: appinfoRepository,
	}
}

func (u *appinfoUsecase) FindCategory(req *appinfo.CategoryFilter) ([]*appinfo.Category, error) {
	category, err := u.appinfoRepository.FindCategory(req)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (u *appinfoUsecase) InsertCategory(category []*appinfo.Category) error {
	err := u.appinfoRepository.InsertCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (u *appinfoUsecase) DeleteCategory(categoryId string) error {
	err := u.appinfoRepository.DeleteCategory(categoryId)
	if err != nil {
		return err
	}
	return nil
}
