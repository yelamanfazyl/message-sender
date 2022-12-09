package startbot

import (
	"upgradeFinal/startbot/cmd/bot"
)

func StartBot(upgradeBot bot.UpgradeBot) {
	upgradeBot.Bot.Handle("/start", upgradeBot.StartHandler)
	upgradeBot.Bot.Handle("/hello", upgradeBot.StartHandler)
	upgradeBot.Bot.Handle("/info", upgradeBot.InfoHandler)
	upgradeBot.Bot.Handle("/delete", upgradeBot.DeleteHandler)

	upgradeBot.Bot.Start()
}
