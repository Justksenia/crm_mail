package rest

import (
	"fmt"
	"log"
	"mail/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) getMailsBox(c *gin.Context) {
	h.corsMiddleware(c)

	boxType := c.Param("type")

	box, err := h.services.GetMailsBox(boxType)

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, box)

}

func (h *Handler) updateMailBox(c *gin.Context) {
	h.corsMiddleware(c)
	var body types.BodyBoxUpdate

	boxId := c.Param("id")

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, error)
		return
	}

	updatedBox, err := h.services.UpdateMailBox(boxId, body)

	if err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, error)
		return
	}

	c.JSON(200, map[string]interface{}{
		"data": updatedBox,
	})

}

func (h *Handler) createBox(c *gin.Context) {

	h.corsMiddleware(c)

	var body types.CreateBoxBody

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, error)
		return
	}
	
	createdBox, err := h.services.CreateMailBox(body)

	if err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, error)
	}

	c.JSON(201, map[string]interface{}{
		"data": createdBox,
	})
}

func(h *Handler) getAllMailBox(c *gin.Context) {
	h.corsMiddleware(c)

	var userInfo map[string]types.UserInfo
	
	if err := c.ShouldBindWith(&userInfo, binding.JSON); err != nil {
		error := fmt.Sprintf(`err json: %v`, err)
		c.JSON(400, error)	
		return
	}

	accountId := userInfo["user_info"].ACCOUNT_ID
	
	folderType := c.Param("type") // mail_folders - родные папки из ящика; account_folders - кастомные папки привязанные к аккаунту
	switch folderType {
		case "mail_folders":
			result, err := h.services.GetAccountMailFolder(accountId)
			log.Println(result)
			if err != nil {
				error := fmt.Sprintf(`err result: %v`, err)
				c.JSON(400, error)
				return
			}
			c.JSON(200, result)	
			return
			
		case "account_folders":
			result, err := h.services.GetAccountFolder(accountId)
			if err != nil {
				error := fmt.Sprintf(`err: %v`, err)
				c.JSON(400, error)
				return
			}
			c.JSON(200, result)	
			return		
	}
}

func (h *Handler) checkConnetcMailBox(c *gin.Context) {
	var data types.ConnectMailBox 

	if err := c.ShouldBindWith(&data, binding.JSON); err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, map[string]interface{}{
			"message": error,
			"result":false,

		})
		return
	}

	message, result := h.services.CheckConnect(data)
	c.JSON(200, map[string]interface{}{
		"message": message,
		"result": result,

	})
}