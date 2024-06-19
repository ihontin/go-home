package models

//go:generate eastjson -all $GOFILE

// easyjson -all models.go
//команда терминала для генерации файла easyjson

type Whetherer struct {
	Records     []Record `json:"records"`
	Skip        int64    `json:"skip"`
	Limit       int64    `json:"limit"`
	TotalAmount int64    `json:"totalAmount"`
}

type Record struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Profile   Profile   `json:"profile"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	CreatedAt string    `json:"createdAt"`
	CreatedBy CreatedBy `json:"createdBy"`
}
type CreatedBy string

type Profile struct {
	Dob        string `json:"dob"`
	Avatar     string `json:"avatar"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	StaticData string `json:"staticData"`
}
