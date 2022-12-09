package main

import (
	"fmt"
	"flag"
	"log"
	"sync"
	"gorm.io/driver/sqlite"
	"github.com/BurntSushi/toml"
	"upgradeFinal/httpserver"
	"gorm.io/gorm"
	"upgradeFinal/startbot/cmd/bot"
	"upgradeFinal/startbot/cmd/bot/models"
	"upgradeFinal/startbot"	
)

type Config struct {
	Env      string
	BotToken string
	Dsn      string
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	db, err := gorm.Open(sqlite.Open(cfg.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}

	upgradeBot := bot.UpgradeBot{
		Bot:   bot.InitBot(cfg.BotToken),
		Users: &models.UserModel{Db: db},
	}

	var wg sync.WaitGroup

	server := server.Server{}

	wg.Add(1)
	
	go func(){
		defer wg.Done()
		fmt.Println("Starting server")
		server.StartServer(upgradeBot) 
	}()


	wg.Add(1)

	go func(){
		defer wg.Done()
		startbot.StartBot(upgradeBot)
	}()

	wg.Wait()
}