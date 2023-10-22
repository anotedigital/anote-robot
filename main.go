package main

import (
	"log"

	"github.com/anonutopia/gowaves"
	macaron "gopkg.in/macaron.v1"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

var conf *Config

var bot *telebot.Bot

var bot2 *telebot.Bot

var anc *gowaves.WavesNodeClient

var db *gorm.DB

var m *macaron.Macaron

var monitor *Monitor

var pc *PriceClient

var cch *Cache

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	m = initMacaron()

	bot = initTelegramBot()

	bot2 = initTelegramBot2()

	initCommands()

	anc = initAnote()

	db = initDb()

	monitor = initMonitor()

	pc = initPriceClient()

	initAnoteToday()

	cch = initCache()

	log.Println("AnoteRobot started.")

	go bot.Start()

	p := int64(20)
	dataTransaction2("%s__priceAnote", nil, &p, nil)

	bot2.Start()
}
