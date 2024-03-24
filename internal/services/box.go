package services

import (
	"fmt"
	"mail/internal/models"
	"mail/internal/types"

	"github.com/emersion/go-imap/client"
)

type BoxService struct {
	model models.Model
}

func NewBoxService(model models.Model) *BoxService{
	return &BoxService{model:model}
}

func (b *BoxService) GetMailsBox(typeBox string) (box []types.FullBox, err error) {
	box, err = b.model.GetMailsBox(typeBox)
	return
}

func (b *BoxService) UpdateMailBox(boxId string, body types.BodyBoxUpdate) (updatedBox []types.Box, err error) {
	updatedBox, err = b.model.UpdateMailBox(boxId, body)
	return
}

func (b *BoxService) CreateMailBox(body types.CreateBoxBody)(createdBox []types.Box, err error){
	createdBox, err = b.model.CreateMailBox(body)
	return
}

func (b *BoxService) GetAccountMailFolder(accountId int)(boxsWithFolders []types.BoxWithFolders, err error){
	boxsWithFolders, err = b.model.GetAccountMailFolder(accountId)
	return
}

func (b *BoxService) CheckConnect(data types.ConnectMailBox)(message string, result bool){


	cl, err := client.DialTLS(data.Host+":993", nil)

	if err != nil {
		message = fmt.Sprintf(`ошибка подключения: %v`, err)	
		result = false
		return 
	}

	if err = cl.Login(data.Login, data.Password); err != nil {
		result = false
		message = fmt.Sprintf("ошибка логина/пароля: %v", err)	
		return
	} else {
		result = true
		message = "можно коннектиться"
	}
	return
}