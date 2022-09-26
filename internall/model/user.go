package model

type User struct {
	ID           uint64 `json:"id"`
	Balance      uint64 `json:"balance"`
	Transactions []Transaction
}

type Transaction struct {
	ID   uint `json:"id"`
	Cash uint `json:"cash"`
}
