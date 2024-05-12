package rest

import (
	"fmt"
	"mail/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) getFoldersByAccount(c *gin.Context) {
	h.corsMiddleware(c)
	var body map[string]types.UserInfo

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, error)
		return
	}

	accountId := body["user_info"].ACCOUNT_ID

	folders, err := h.services.Folder.GetFoldersByAccount(accountId)
	
	if err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, error)
		return
	}

	c.JSON(200, folders)
}

func (h *Handler) createFolder(c *gin.Context) {
	h.corsMiddleware(c)
	var data types.FolderBody

	if err := c.ShouldBindWith(&data, binding.JSON); err != nil {
		error := fmt.Sprintf(`err : %v`, err)
		c.JSON(400, error)
		return
	}

	newFolder, err := h.services.Folder.CreateFolder(data)

	if err != nil {
		error := fmt.Sprintf(`err : %v`, err)
		c.JSON(400, error)
		return
	}

	c.JSON(200, map[string]interface{}{
		"data": newFolder,
	})
}

func (h *Handler) updateFolder(c *gin.Context) {
	h.corsMiddleware(c)
	var bodyData types.FolderBody
	
	folderId := c.Param("id")

	if err := c.ShouldBindWith(&bodyData, binding.JSON); err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, error)
		return
	}
	
	updatedFolder, err := h.services.Folder.UpdateFolder(bodyData, folderId)
	
	if err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, error)
		return
	}

	c.JSON(200, map[string]interface{}{
		"data": updatedFolder,
	})
}

func (h *Handler) deleteFolder(c *gin.Context) {
	h.corsMiddleware(c)
	
	var body types.FolderBody
	
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		error := fmt.Sprintf(`err: %v`, err)
		c.JSON(400, error)
		return
	}

	folderId := c.Param("id")
	err := h.services.Folder.DeleteFolder(body, folderId)

	if err != nil {
		c.JSON(400, map[string]interface{}{
			"message": fmt.Sprintf(`err: %v`, err),
			"success":false,
		})
	}

	c.JSON(200, map[string]interface{}{
		"message": "",
		"success":true,
	})
}