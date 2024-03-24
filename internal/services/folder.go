package services

import (
	"mail/internal/models"
	"mail/internal/types"
)

type FolderService struct {
	model models.Model
}


func NewFolderService(model models.Model) *FolderService {
	return &FolderService{model: model}
}

func (f *FolderService) GetAccountFolder(accountId int)(folders []types.Folder, err error){
	folders, err = f.model.GetAccountFolder(accountId)
	return
}