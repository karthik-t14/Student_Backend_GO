package models

type Student struct {
	Id     int    `json:"id"`
	Name   string `json:"sname"`
	Age    int    `json:"age"`
	Branch string `json:"branch"`
}
