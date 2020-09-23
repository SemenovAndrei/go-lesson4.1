package dto

type CardDTO struct {
	Id     int64  `json:"id"`
	UserId string `json:"userId"`
	Number int64  `json:"number"`
	Type   string `json:"type"`
	System string `json:"system"`
}

type CardErrDTO struct {
	Err string `json:"error"`
}