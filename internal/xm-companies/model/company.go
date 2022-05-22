package model

type Company struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}
