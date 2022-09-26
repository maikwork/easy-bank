package repository

const (
	GetBalanceFromUserByID  = "select balance from users where id=%v"
	GetAllFromUserByID      = "select * from users where id=%v"
	UpdateBalanceToUserByID = "update users set balance=%v where id=%v"
	CreateUser              = "insert into users values(%v, %v)"
)
