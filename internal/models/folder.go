package models

import (
	"fmt"
	"mail/internal/types"

	"github.com/jmoiron/sqlx"
)

type FolderPostgres struct {
	db *sqlx.DB
}

func NewFolderPostgres(db *sqlx.DB) *FolderPostgres {
	return &FolderPostgres{db: db}
}

func (f *FolderPostgres) GetAccountFolder(accountId int) (folders []types.Folder, err error) {

	q := fmt.Sprintf(`SELECT id, name, color, box_id FROM folders WHERE account_id = %d AND box_id = 0 ORDER BY box_id ASC`, accountId )   
	err = f.db.Select(&folders, q)
	
	if err!= nil {
		return
	}

	if len(folders) == 0 {
		folders = []types.Folder{}
	}

	return
}