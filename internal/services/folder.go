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
	folders, err = f.model.GetAccountFolders(accountId)
	return
}

func (f *FolderService) GetFoldersByAccount(accountId int)(folders []types.Folder, err error){
	folders, err = f.model.GetFoldersByAccount(accountId)
	return
}

func (f *FolderService) CreateFolder(data types.FolderBody)(folder types.Folder, err error){
	folder, err = f.model.CreateFolder(data)
	return
}

func (f *FolderService) UpdateFolder(data types.FolderBody, folderId string)(folder types.Folder, err error) {
	folder, err = f.model.UpdateFolder(data, folderId)
	return
}

func (f *FolderService) DeleteFolder(data types.FolderBody, folderId string)(err error) {
	err = f.model.DeleteFolder(data, folderId)
	return
}