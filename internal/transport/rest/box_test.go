package rest

import (
	"bytes"
	"mail/internal/services"
	mock_services "mail/internal/services/mocks"
	"mail/internal/types"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	//"github.com/golang/mock/gomock"
)

func TestHandler_createBox(t *testing.T) {
	type mockBehavior func (m *mock_services.MockBox, t types.CreateBoxBody, b []types.Box)
	testConnecting := []struct{
		name string  //test name
		inputBody string //input data in json from client
		connectData types.CreateBoxBody // type  response body
		mockBehavior mockBehavior
		statusCode int // expecred status code
		requestBody string //expected request body json
		returnData []types.Box
	}{
		{
			name: "Ok",
			// то что приходит с клиента
			inputBody: `{"user_info":{"ACCOUNT_ID":1}, "data":{"color":"red", "account_json":"[1,2]"}}`,
			// структура того что  приходит с клиента
			connectData: types.CreateBoxBody {
				User_Info: types.UserInfo{
					ACCOUNT_ID: 1,
				},
				Body: types.BoxData{
					Color: "red",
					AccountJson: "[1,2]",
				},
			},
			// структура возвращаемых данных от сервиса 
			returnData: []types.Box{
				{
					Color: "red",
					AccountJson: "[1,2]",
					Name:"",
					Id:15,
				},
			},
			mockBehavior: func(conn *mock_services.MockBox, data types.CreateBoxBody, d []types.Box) {
				conn.EXPECT().CreateMailBox(data).Return(d,nil)
			},
			statusCode: 201,
			requestBody: `{"success":true}`,
		},
	}
	for _,value := range testConnecting{
		t.Run(value.name, func(t *testing.T){
			 c:= gomock.NewController(t)
			 defer c.Finish()
			 box := mock_services.NewMockBox(c)

			value.mockBehavior(box, value.connectData, value.returnData)
			 res := &services.Service{Box:box}
			 handler := NewHandler(res)
			 r := gin.New()
			 r.POST("/mail/box/create", handler.createBox)
			 w:= httptest.NewRecorder()
			 req := httptest.NewRequest("POST", "/mail/box/create", bytes.NewBufferString(value.inputBody) )
			 r.ServeHTTP(w,req)
			 assert.Equal(t, w.Code, value.statusCode)
			 assert.Equal(t, w.Body.String(), value.requestBody)
			
		})
	}
}
func TestHandler_checkConnectMailBox(t *testing.T) {
	
	type mockBehavior func(s *mock_services.MockBox, t types.ConnectMailBox)
	testConnecting := []struct{
		name string
		inputBody string
		connectData types.ConnectMailBox
		mockBehavier mockBehavior
		statusCode int
		expectedRequestBody string

	}{
		{
			name: "ok",
			inputBody: `{"host":"smtp.yandex.ru", "login":"fokina", "password":"1234", "username":"username","id":12 }`,
			connectData: types.ConnectMailBox{
				Host: "smtp.yandex.ru" ,
				Login: "fokina",
				Password: "1234",
				Username: "username",
				Id: 12,

			},
			mockBehavier: func(conn *mock_services.MockBox, data types.ConnectMailBox){
				conn.EXPECT().CheckConnect(data).Return("можно коннектиться", true)
			},
			statusCode: 200,
			expectedRequestBody: "1",

		},
	}
	for _,value := range testConnecting{
		t.Run(value.name, func(t *testing.T){
			 c:= gomock.NewController(t)
			 defer c.Finish()
			 box := mock_services.NewMockBox(c)

			 value.mockBehavier(box, value.connectData)
			 res := &services.Service{Box:box}
			 handler := NewHandler(res)
			 r := gin.New()
			 r.POST("/mail/box/check_connect", handler.checkConnetcMailBox)
			 w:= httptest.NewRecorder()
			 req := httptest.NewRequest("POST", "/mail/box/check_connect", bytes.NewBufferString(value.inputBody) )
			 r.ServeHTTP(w,req)
			 assert.Equal(t, w.Code, value.statusCode)
			 assert.Equal(t, w.Body.String(), value.expectedRequestBody)
			
		})
	}

}