package repository

import "github.com/maikwork/balanceUserAvito/internall/model"

type UserStore interface {
	Save(uint64, int64)
	Find(uint64) model.User
	Transfer(uint64, uint64, int64)
}
