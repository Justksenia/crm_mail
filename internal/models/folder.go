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

func (f *FolderPostgres) GetAccountFolders(accountId int) (folders []types.Folder, err error) {

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

func (f *FolderPostgres) GetFoldersByAccount(accountId int) (folders []types.Folder, err error) {

	q := fmt.Sprintf(`SELECT * FROM public.folders WHERE account_id = %d ORDER BY id ASC`, accountId)
	err = f.db.Select(&folders, q)

	if len(folders) == 0 {
		folders = []types.Folder{}
	}

	return

}

func (f *FolderPostgres) CreateFolder(data types.FolderBody)(folder types.Folder, err error) {

	body := data.Body
	accountId := data.User_Info.ACCOUNT_ID
	var folderId int

	q := fmt.Sprintf(`INSERT INTO public.folders (name, account_id, color) values ('%s', %d, '%s') RETURNING id`, body.Name, accountId, body.Color)
	err = f.db.Select(&folderId, q) 

	if err!= nil {
		return
	}
	
	if (folderId > 0) {
		q = fmt.Sprintf(`SELECT * FROM public.folders WHERE id = %d`, folderId)
		err = f.db.Select(&folder, q);
	} else {
		folder = types.Folder{}
	}

	return
}

func (f *FolderPostgres) UpdateFolder(data types.FolderBody, folderId string)(updatedFolder types.Folder, err error) {

	var q string
	var updatedFolderId int
	
	if (data.Body.Color != "" && data.Body.Name == "") {
		q = fmt.Sprintf(`UPDATE folders SET color = '%s'`,data.Body.Color)
	} else if (data.Body.Name != "" && data.Body.Color == "") {
		q = fmt.Sprintf(`UPDATE folders SET name = '%s'`,data.Body.Name)
	} else {
		q = fmt.Sprintf(`UPDATE folders SET name = '%s', color = '%s'`,data.Body.Name, data.Body.Color)
	}
	q += fmt.Sprintf(` WHERE id = %s AND account_id = %d RETURNING id`, folderId, data.User_Info.ACCOUNT_ID)

	err = f.db.Select(&updatedFolderId, q)

	if err != nil {
		return 
	}
	
	if (updatedFolderId > 0) {
		q = `SELECT * FROM folders WHERE id = $1;`
		err = f.db.Select(&updatedFolder, q, folderId);
	} else {
		updatedFolder = types.Folder{}
	}

	return
}

func (f *FolderPostgres) DeleteFolder(data types.FolderBody, folderId string)(err error) {
	var deleteFolder []types.Folder
	var deleteFolderId int

	q := fmt.Sprintf(`SELECT * FROM folders WHERE id = %s AND account_id = %d `, folderId, data.User_Info.ACCOUNT_ID)
	err = f.db.Select(&deleteFolder, q)

	if err != nil {
		return
	}

	if (len(deleteFolder) > 0) {
		q := fmt.Sprintf(`INSERT INTO folders_archive (name, color, account_id, previous_id) values ('%s', '%s', %d, %d) RETURNING id`, deleteFolder[0].Name, deleteFolder[0].Color, data.User_Info.ACCOUNT_ID, deleteFolder[0].Id)
		err = f.db.QueryRow(q).Scan(&deleteFolderId)

		if err != nil {
			return
		}

		q = fmt.Sprintf(`DELETE FROM public.folders WHERE id = %s`, folderId)
		_, err = f.db.Exec(q)

		if err != nil {
			return
		}
	
		q = fmt.Sprintf (
			`UPDATE db_messages SET folders = REPLACE(folders,' ','') WHERE del = false;
			UPDATE db_messages SET folders = REPLACE(folders,'"%s",','') WHERE del = false;
			UPDATE db_messages SET folders = REPLACE(folders,',"%s"','') WHERE del = false;
			UPDATE db_messages SET folders = REPLACE(folders,'"%s"','') WHERE del = false;
			UPDATE db_messages SET folders = REPLACE(folders,', "%s"','') WHERE del = false RETURNING id; `, folderId, folderId, folderId, folderId)
		
		_, err = f.db.Exec(q)

		if err != nil {
			return
		}
	}
	
	return
}