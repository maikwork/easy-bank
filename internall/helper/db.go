package helper

import (
	"fmt"

	"github.com/maikwork/balanceUserAvito/internall/model"
)

func GetDSN(d *model.DBSetting) string {
	dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v", d.Type, d.Username,
		d.Password, d.Host,
		d.Port, d.DBName)
	return dsn
}
