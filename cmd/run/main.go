package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/maikwork/balanceUserAvito/internall/db"
	"github.com/maikwork/balanceUserAvito/internall/helper"
	"github.com/maikwork/balanceUserAvito/internall/repository"
	"github.com/maikwork/balanceUserAvito/internall/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	path := flag.String("c", "cmd/config.yml", "Path to config")
	flag.Parse()

	ctx := context.Background()

	config := helper.ReadConfig(*path)
	log.Info("config: ", config)

	dsn := helper.GetDSN(&config.DBSetting)
	db := db.Connect(dsn)
	defer db.Close(ctx)

	store := repository.NewPSQL(db)

	s := server.New(store)
	fmt.Println("Sever run...")
	s.Run()
}
