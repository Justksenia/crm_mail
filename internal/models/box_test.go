package models

import (
	"log"
	"mail/internal/types"
	"testing"

	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestCreateMailBox(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	r := NewModel(db)

	// data := types.CreateBoxBody {
	// 	Body: types.BoxData{
	// 		Name:"test",
	// 		Color: "black",
	// 		AccountJson: "[1,2,3]",
	// 	},
	// 	User_Info: types.UserInfo{
	// 		ACCOUNT_ID: 1,
	// 	},
	// }


	//type mockBehavier func(mockData types.CreateBoxBody)
	type mockBehavier func(args types.CreateBoxBody)

	testCase := []struct{
		name string
		mockBehavier mockBehavier
		data types.CreateBoxBody
		result []types.Box
		wantError bool
	}{
		{
			name:"ok",
			data: types.CreateBoxBody {
				Body: types.BoxData{

						Name:"test",
						Color: "black",
						AccountJson: "[1,2,3]",
					},
					User_Info: types.UserInfo{
						ACCOUNT_ID: 1,
					},

			
			},
			result: []types.Box{
				{
					Id: 13,
					Name: "test",
					Color: "black",
					AccountJson: "[1,2,3]",
				},
			},
			mockBehavier: func(data){
				mock.ExpectQuery("INSERT INTO public.mail_box").WithArgs( )
			},
			wantError: false,
		},
	}


}