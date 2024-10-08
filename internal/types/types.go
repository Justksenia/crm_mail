package types

type UserInfo struct {
	ACCOUNT_ID    int 
	PERMISSION_ID int
}

type FullBox struct {
	Id          int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Color    string `db:"color" json:"color"`
	AccountJson string   `db:"account_json" json:"account_json"`
	Active bool `json:"active" db:"is_active"`
	Host string `json:"host"`
	Login string `json:"login"`
	Password string `json:"password"`
	Connect bool `db:"is_connect" json:"connect"`
	OldMail bool `db:"is_download_old_mail" json:"is_download_old_mail"`
	SmtpHost string `db:"smtp_host" json:"smtp_host"`
}

type Box struct {
	Id          int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Color    string `db:"color" json:"color"`
	AccountJson string   `db:"account_json" json:"account_json"`
}

type BoxWithFolders struct {
	Box
	Folders []Folder `json:"folders"`
}

type Folder struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Color string `json:"color" db:"color"`
	Box_Id int 	`json:"box_id" db:"box_id"`
}

type BodyBoxUpdate struct {
	User_Info UserInfo 
	Body map[string]any
}

type CreateBoxBody struct {
	Body BoxData	`json:"data"`
	User_Info UserInfo `json:"user_info"`
}

type BoxData struct {
	Name     string  `json:"name"`
	Color    string	`json:"color"`
	AccountJson string `json:"account_json"`
}

type ConnectMailBox struct {
	Host     string
	Login    string
	Password string
	Username string
	Id       int
}

type FolderBody struct {
	Body FolderBodyData
	User_Info UserInfo
}

type FolderBodyData struct {
	Name      string
	Color     string
}

var (
	Pa int
	Ra string
	s string
)