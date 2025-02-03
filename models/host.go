package models

type Host struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IP       string `json:"ip"`
	OS       string `json:"os"`
	Location string `json:"location"`
}