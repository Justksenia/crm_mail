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
	GetAccountFolder(accountId int)([]types.Folder, error)
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