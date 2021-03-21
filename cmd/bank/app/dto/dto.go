package dto

import "time"

type CardDTO struct {
	Id int64 `json:"id"`
	Number string `json:"number"`
	Balance int64 `json:"balance"`
	Issuer string `json:"issuer"`
	OwnerId int64 `json:"owner_id"`
	Status string `json:"status"`
}

type TransactionsDTO struct {
	Id int64 `json:"id"`
	CardId	int64 `json:"card_id"`
	Sum	int64 `json:"sum"`
	MccId	int64 `json:"mcc_id"`
	Description	string `json:"description"`
	Status	string `json:"status"`
	Created	time.Time	`json:"created"`
}