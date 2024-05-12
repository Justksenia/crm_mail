package models

import (
	"mail/internal/types"

	"github.com/jmoiron/sqlx"
)

type Box interface {
	GetMailsBox(typeBox string) ( []types.FullBox, error)
	UpdateMailBox(boxId string, body types.BodyBoxUpdate)([]types.Box, error)
	CreateMailBox(body types.CreateBoxBody)([]types.Box, error)
	GetAccountMailFolder(accountId int)([]types.BoxWithFolders, error)
	
}

type Folder interface {
	GetAccountFolders(accountId int)([]types.Folder, error)
	GetFoldersByAccount(accountId int)([]types.Folder, error)
	CreateFolder(types.FolderBody)(types.Folder, error)
	UpdateFolder(types.FolderBody, string)(types.Folder, error) 
	DeleteFolder(types.FolderBody, string)(error)
}

type Model struct {
	Box
	Folder
}

func NewModel(db *sqlx.DB) *Model {

	return &Model{
		Box: NewBoxPostgres(db), 
		Folder: NewFolderPostgres(db), 
	}
}