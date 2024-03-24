package models

import (
	"fmt"
	"log"
	"mail/internal/types"

	"github.com/jmoiron/sqlx"
)

type BoxPostgres struct {
	db *sqlx.DB
}

func NewBoxPostgres(db *sqlx.DB) *BoxPostgres {
	return &BoxPostgres{db:db}
}

func (b *BoxPostgres) GetMailsBox(typeBox string) (box []types.FullBox, err error) {

	isActive := false

	if typeBox == "active" {
		isActive = true
	}

	q := fmt.Sprintf(`SELECT * FROM mail_box WHERE is_active = true AND is_connect = %t`, isActive)

	err = b.db.Select(&box, q)

	if err != nil {
		return
	}

	if len(box) == 0 {
		box = []types.FullBox{}
		return
	}

	return 

}

func (b *BoxPostgres) UpdateMailBox(boxId string, body types.BodyBoxUpdate) (updatedMailBox []types.Box, err error) {

	var updatedBoxId int
	var column string
	var val any

	for key, value := range body.Body {
		column = key
		val = value
	}

	sql := fmt.Sprintf(`UPDATE mail_box SET %s = $1 WHERE id = %s RETURNING id;`, column, boxId)
	err = b.db.QueryRow(sql, val).Scan(&updatedBoxId)

	if err != nil {
		return
	}

	if updatedBoxId > 0 {
		q := fmt.Sprintf(`SELECT id, name, color, account_json FROM public.mail_box WHERE id = %d`, updatedBoxId)
		err = b.db.Select(&updatedMailBox, q)

		if err != nil {
			return
		}
	} else {
		updatedMailBox = []types.Box{}
	}

	return
	
}

func (b *BoxPostgres) CreateMailBox(body types.CreateBoxBody) (createdBox []types.Box, err error) {

	var boxId int

	q := fmt.Sprintf(`INSERT INTO public.mail_box (name, color, account_json) values ('%s', '%s', '%s') RETURNING id;`, body.Body.Name, body.Body.Color, body.Body.AccountJson)
	err = b.db.QueryRow(q).Scan(&boxId)

	if err != nil {
		return
	}

	if boxId > 0 {
		queryNewBox := `SELECT id, name, color, account_json FROM public.mail_box WHERE id = $1;`
		err = b.db.Select(&createdBox, queryNewBox, boxId)
		if err != nil {
			return
		}
	} else {
		createdBox = []types.Box{}
		return
	}

	return
}

func (b *BoxPostgres) GetAccountMailFolder(accountId int) (boxWithFolders []types.BoxWithFolders, err error) {
	
	// var boxs []types.Box
	var folders []types.Folder
	var foldersBox []types.Folder

	sql := fmt.Sprintf(`SELECT id, name, color FROM mail_box WHERE account_json::jsonb @> '%d' AND is_active = true ORDER BY id ASC`, accountId) 
	err = b.db.Select(&boxWithFolders, sql)
	
	if err != nil {
		return
	}

	sql = fmt.Sprintf(`SELECT id, name, color, box_id FROM folders WHERE account_id = %d AND box_id != 0 ORDER BY box_id ASC`, accountId) 

	err = b.db.Select(&folders, sql)

	log.Println(boxWithFolders)
	
	if err != nil {
		log.Println(err)
		return   
	}

	for i:=0; i< len(boxWithFolders); i++ {
	  for j:=0; j< len(folders); j++ {
		if boxWithFolders[i].Id == folders[j].Box_Id {
			foldersBox = append(foldersBox, folders[j])
		}
	}
	boxWithFolders[i].Folders = foldersBox
	foldersBox = nil
}
return

}