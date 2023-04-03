package model

type User struct {
	FIO      string `json:"fio"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
