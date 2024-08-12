package services

import (
	"mail/internal/models"
	"mail/internal/types"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Box interface {
	GetMailsBox(typeBox string) ([]types.FullBox, error)
	UpdateMailBox(boxId string, body types.BodyBoxUpdate)([]types.Box, error)
	CreateMailBox(body types.CreateBoxBody)([]types.Box, error)
	GetAccountMailFolder(accountId int)([]types.BoxWithFolders, error)
	CheckConnect (types.ConnectMailBox)(string, bool)
}

type Folder interface {
	GetAccountFolder(accountId int)([]types.Folder, error)
	GetFoldersByAccount(accountId int)([]types.Folder, error)
	CreateFolder(types.FolderBody)(types.Folder, error)
	UpdateFolder(types.FolderBody, string)(types.Folder, error) 
	DeleteFolder(types.FolderBody, string)(error)
}


type Service struct {
	Box
	Folder
}

func NewService(m models.Model) *Service {
	return &Service{
		Box : NewBoxService(m),
		Folder : NewFolderService(m),
	}
}