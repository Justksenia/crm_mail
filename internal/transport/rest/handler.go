package rest

import (
	"mail/internal/services"

	"github.com/gin-gonic/gin"

)

type Handler struct {
	services *services.Service
	// db *sqlx.DB
	// rdb *redis.Client
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	mail := router.Group("/mail")
	box := mail.Group("/box")
	folder := mail.Group("/folder")
	// x5 := router.Group("/x5")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	// mails.POST("/", h.getAllMessages) //все письма, без пагинации
	// mails.OPTIONS("/", h.optionsMessage)

	// mails.POST("/get_message/:id", h.getMessageId) // получение одного письма по ид
	// mails.OPTIONS("/get_message/:id", h.optionsMessage)

	// mails.PUT("/update_message/:id", h.updateMessage) // изменение поля у одного письма, хз нужно или нет
	// mails.OPTIONS("/update_message/:id", h.optionsMessage)

	// mails.POST("/count_messages/not_read", h.getCountMessagesNotRead) // все письма не прочитанные
	// mails.OPTIONS("/count_messages/not_read", h.optionsMessage)

	// mails.POST("/get_messages/del/:box_id", h.getDeleteMessages) //получение удаленных писем по ид ящика было mail/del/:id
	// mails.OPTIONS("/get_messages/del/:box_id", h.optionsMessage)

	// mails.POST("/get_messages/box/:id", h.getMessagesByBoxId) // получение все писем по ид ящика
	// mails.OPTIONS("/get_messages/box/:id", h.optionsMessage)

	// mails.POST("/get_messages/folder/:id", h.getMessagesByFolderId) // получение всех писем по ид папки
	// mails.OPTIONS("/get_messages/folder/:id", h.optionsMessage)

	// mails.POST("/create_mail", h.createMessage) // создание сообщение, приходит из java скрипта было read_mail нет в тест плане
	// mails.OPTIONS("/create_mail", h.optionsMessage)

	// mails.POST("/send_message", h.sendMessage) // отправка письма 
	// mails.OPTIONS("/send_message", h.optionsMessage)

	// mails.POST("/get_send_messages/box/:id", h.getSendMessagesByBoxId) // получение отправленных сообщений по ид ящика /send_messages/:box_id
	// mails.OPTIONS("/get_send_messages/box/:id", h.optionsMessage)

	// mails.POST("/get_send_message/:id", h.getSendMessage) //полуечние отправленного письма
	// mails.OPTIONS("/get_send_message/:id", h.optionsMessage)

	// mails.POST("/update_messages", h.messagesFilter) // фильтрация писем /filter/mail
	// mails.OPTIONS("/update_messages", h.optionsMessage)

	
	router.OPTIONS("mail/box/:type", h.optionsMessage)
	box.POST("/:type", h.getAllMailBox) //все ящики type - mail_folders - родные папки из ящика; account_folders - кастомные папки привязанные к аккаунту

	box.PUT("/update/:id", h.updateMailBox) // изменение ящика
	router.OPTIONS("mail/box/update/:id", h.optionsMessage)

	box.POST("/create", h.createBox) // создание ящика box/new
	router.OPTIONS("mail/box/create", h.optionsMessage)

	box.GET("monitoring/:type", h.getMailsBox) // мониторинг ящиков из java  :active or not_active
	router.OPTIONS("mail/box/monitoring/:type", h.optionsMessage)

	folder.POST("/get_folders", h.getFoldersByAccount) //все папки
	router.OPTIONS("mail/folder/get_folders", h.optionsMessage)

	folder.POST("/create_folder", h.createFolder) // создание папки
	router.OPTIONS("/mail/folder/create_folder", h.optionsMessage)

	folder.PUT("/update_folder/:id", h.updateFolder) // изменение папки по ид
	folder.OPTIONS("/update_folder/:id", h.optionsMessage)

	folder.DELETE("/folder/:id", h.deleteFolder) //удаление папки по ид
	folder.OPTIONS("/folder/:id", h.optionsMessage)

	// mails.GET("/tags", h.getAllTags) // получение списка тэгов
	// mails.OPTIONS("/tags", h.optionsMessage)

	// mails.POST("/count/:type", h.countFilter)
	// // подсчет количества тэгов или папок по фильтру  : tags or :folders  body-> {"props":"box", "tags":"1,2,3", "props_id":"1", "user_info": {"account_id":25, "permission_id":1}} props: all, is_read, is_main, attachment, is_favorites, box, folder. Если выбираем box/ folder, то добавляем props_id:"number - ид папки или ящика, tags не обязательный
	// mails.OPTIONS("/count/:type", h.optionsMessage)

	// mails.POST("/broadcast", h.setBroadcast) // механизм трансляцими
	// mails.OPTIONS("/broadcast", h.optionsMessage)

	// mails.POST("/get_broadcast", h.getBroadcast) // 
	// mails.OPTIONS("/get_broadcast", h.optionsMessage)

	// mails.GET("/broadcast/:id", h.getBroadcastId) // проверка кому транслировали, для тестов
	// mails.OPTIONS("/broadcast/:id", h.optionsMessage)

	// mails.POST("/box/create", h.addMailBox)  // добавление нового ящика
	// mails.OPTIONS("/box/create", h.optionsMessage)
	
	box.POST("/box/check_connect", h.checkConnetcMailBox)  // проверка подключения ящика
	router.OPTIONS("mail/box/check_connect", h.optionsMessage)

	// mails.POST("/userbox", h.createUserBox)
	// mails.OPTIONS("/userbox", h.optionsMessage)

	// mails.GET("/redis", h.testRedis)

	// x5.OPTIONS("/code", h.optionsMessage)
	// x5.GET("/code", h.getCodeX5Portal)

	return router
}
