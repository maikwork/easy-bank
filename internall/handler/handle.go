package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/maikwork/balanceUserAvito/internall/model"
	"github.com/maikwork/balanceUserAvito/internall/repository"
	"github.com/maikwork/balanceUserAvito/pkg"
	log "github.com/sirupsen/logrus"
)

const (
	TRANSACTION = "transaction"
	TRANSFER    = "transfer"
	CURRENT     = "current"
	ADDED       = "added"
	WITHDRAWAL  = "withdrawal"
)

type urlHandle struct {
	id    string
	query string
	args  url.Values
}

type UserHandler struct {
	store repository.UserStore
	url   urlHandle
}

func NewHandeler(store repository.UserStore) UserHandler {
	return UserHandler{
		store: store,
	}
}

func (h UserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	var resp model.Response
	var data []byte

	query, id, err := getTypeQueryAndID(r.URL.Path)
	if err != nil {
		return
	}

	h.url = urlHandle{
		id:    id,
		query: query,
		args:  r.URL.Query(),
	}

	switch query {
	case TRANSACTION:
		resp = h.getAllTransactions()
		// data, _ = json.Marshal(trans)
	case TRANSFER:
		resp = h.transfer()
	case CURRENT:
		resp = h.current()
		// data, _ = json.Marshal(amount)
	case ADDED:
		resp = h.added()
	case WITHDRAWAL:
		resp = h.withdrawal()
	default:
		data = []byte(h.notfound())
		return
	}

	data, err = json.Marshal(resp)
	if err != nil {
		log.WithError(err).Info("can't marshal response")
	}

	w.Write(data)
}

// GET /user/:id/transaction
func (h UserHandler) getAllTransactions() model.Response {
	var data interface{}

	id, err := strconv.Atoi(h.url.id)
	if err != nil {
		data = model.ErrorHTTP{
			Err:  err.Error(),
			Desc: "wrang is id of user",
		}
		return model.Response{Success: false, Data: data}
	}
	data = h.store.Find(uint64(id))

	result := model.Response{
		Success: true,
		Data:    data,
	}

	return result
}

// GET /user/:id/transgre?to={id}&m={amount}
func (h UserHandler) transfer() model.Response {
	var data interface{}

	id, err := strconv.Atoi(h.url.id)
	if err != nil {
		data = model.ErrorHTTP{
			Err:  err.Error(),
			Desc: "wrang is id of user",
		}
		return model.Response{Success: false, Data: data}
	}

	toID, err := strconv.Atoi(h.url.args.Get("to"))
	if err != nil {
		data = model.ErrorHTTP{
			Err:  err.Error(),
			Desc: "wrang is 'to' of user",
		}
		return model.Response{Success: false, Data: data}
	}

	m, err := strconv.Atoi(h.url.args.Get("m"))
	if err != nil {
		data = model.ErrorHTTP{
			Err:  err.Error(),
			Desc: "wrang is 'm' of user",
		}
		return model.Response{Success: false, Data: data}
	}
	h.store.Transfer(uint64(id), uint64(toID), int64(m))

	result := model.Response{
		Success: true,
		Data:    data,
	}
	return result
}

// GET /user/:id/current -> текущий баланс в RUB
// GET /user/:id/current?currency={EUR/USD...} -> баланс в заданной валюте
func (h UserHandler) current() model.Response {
	var data interface{}

	id, err := strconv.Atoi(h.url.id)
	if err != nil {
		data = model.ErrorHTTP{
			Err:  err.Error(),
			Desc: "wrang is id of user",
		}
		return model.Response{Success: false, Data: data}
	}

	user := h.store.Find(uint64(id))
	data = fmt.Sprint(user.Balance)

	if h.url.args.Has("currency") {
		c := h.url.args.Get("currency")
		data = pkg.Convert("xOS3mOPbf1jtyL7hH9E6Bd6MNYznWiJQ", "RUB", c, data.(string))
	}

	result := model.Response{
		Success: true,
		Data:    data,
	}
	return result
}

// GET /user/:id/added?m={amount}
func (h UserHandler) added() model.Response {
	var data interface{}

	id, err := strconv.Atoi(h.url.id)
	if err != nil {
		data = model.ErrorHTTP{
			Err:  err.Error(),
			Desc: "wrang is id of user",
		}
		return model.Response{Success: false, Data: data}
	}

	value, _ := strconv.Atoi(h.url.args.Get("m"))
	if err != nil {
		data = model.ErrorHTTP{
			Err:  err.Error(),
			Desc: "wrang is 'm' of user",
		}
		return model.Response{Success: false, Data: data}
	}

	h.store.Save(uint64(id), int64(value))

	result := model.Response{
		Success: true,
		Data:    data,
	}
	return result
}

// GET /user/:id/withdrawal?m={amount}
func (h UserHandler) withdrawal() model.Response {
	var data interface{}

	id, err := strconv.Atoi(h.url.id)
	if err != nil {
		data = model.ErrorHTTP{
			Err:  err.Error(),
			Desc: "wrang is id of user",
		}
		return model.Response{Success: false, Data: data}
	}
	value, err := strconv.Atoi(h.url.args.Get("m"))
	if err != nil {
		data = model.ErrorHTTP{
			Err:  err.Error(),
			Desc: "wrang is 'm' of user",
		}
		return model.Response{Success: false, Data: data}
	}
	h.store.Save(uint64(id), int64(-value))

	result := model.Response{
		Success: true,
		Data:    data,
	}
	return result
}

func (h UserHandler) notfound() string {
	return "Not FOUND!!!"
}

func getTypeQueryAndID(path string) (string, string, error) {
	var err error
	var p string
	var id string
	parr := strings.Split(path, "/")
	fmt.Println(parr)
	l := len(parr)

	if l-2 < 0 {
		err = fmt.Errorf("Don't have id and type query")
		return p, id, err
	}

	p = parr[len(parr)-1]
	id = parr[len(parr)-2]

	return p, id, err
}
