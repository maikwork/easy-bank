package repository

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v5"
	"github.com/maikwork/balanceUserAvito/internall/model"
)

const (
	errNoRowsInResultSet = "no rows in result set"
)

type PSQL struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewPSQL(db *pgx.Conn) PSQL {
	return PSQL{
		db:  db,
		ctx: context.Background(),
	}
}

func (p PSQL) Save(id uint64, value int64) {
	var money int64
	var query string

	query = fmt.Sprintf(GetBalanceFromUserByID, id)
	row := p.db.QueryRow(p.ctx, query)
	err := row.Scan(&money)
	if err != nil {
		if err.Error() == errNoRowsInResultSet {
			query := fmt.Sprintf(CreateUser, id, value)
			p.db.Exec(p.ctx, query)
			return
		}

		log.WithError(err).Info("can't scan row in db")
		return
	}

	if money+value < 0 {
		err := fmt.Errorf("Can't make a transfer operation, not enough funds")
		log.WithError(err).Info("недостаточно средств")
		return
	}

	balance := money + value
	query = fmt.Sprintf(UpdateBalanceToUserByID, balance, id)
	p.db.Exec(p.ctx, query)
}

func (p PSQL) Find(id uint64) model.User {
	var user model.User
	var query string

	query = fmt.Sprintf(GetAllFromUserByID, id)
	row := p.db.QueryRow(p.ctx, query)
	err := row.Scan(&user.ID, &user.Balance)
	if err != nil {
		log.WithError(err).Info("не может прочитать db")
	}

	query = fmt.Sprintf("select to_id, amount from transactions where from_id=%v", id)
	arr, err := p.db.Query(p.ctx, query)

	for arr.Next() {
		tmpTrans := model.Transaction{}
		arr.Scan(&tmpTrans.ID, &tmpTrans.Cash)
		user.Transactions = append(user.Transactions, tmpTrans)
	}

	return user
}

func (p PSQL) Transfer(from, to uint64, value int64) {
	var money int64
	var query string

	query = fmt.Sprintf(GetBalanceFromUserByID, from)
	row := p.db.QueryRow(p.ctx, query)
	if err := row.Scan(&money); err != nil {
		log.WithError(err).Info("не может прочитать db")
		return
	}

	if money-value < 0 {
		err := fmt.Errorf("Can't make a transfer operation, not enough funds")
		log.WithError(err).Info("недостаточно средств")
		return
	}
	query = fmt.Sprintf(UpdateBalanceToUserByID, money-value, from)
	p.db.Exec(p.ctx, query)

	query = fmt.Sprintf(GetBalanceFromUserByID, to)
	row = p.db.QueryRow(p.ctx, query)
	if err := row.Scan(&money); err != nil {
		log.WithError(err).Info("не может прочитать db")
		return
	}

	balance := money + value

	query = fmt.Sprintf(UpdateBalanceToUserByID, balance, to)
	p.db.Exec(p.ctx, query)

	query = fmt.Sprintf("insert into transactions(from_id, to_id, amount) values(%v, %v, %v)", from, to, value)
	p.db.Exec(p.ctx, query)
}
