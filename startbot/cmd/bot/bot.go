package bot

import (
	"log"
	"time"
	"strconv"
	"upgradeFinal/startbot/cmd/bot/models"

	"gopkg.in/telebot.v3"
)

type UpgradeBot struct {
	Bot   *telebot.Bot
	Users *models.UserModel
}

func (bot *UpgradeBot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Name:       ctx.Sender().Username,
		TelegramId: ctx.Chat().ID,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
		ChatId:     ctx.Chat().ID,
		MessageAmount: 0,
	}

	existUser, err := bot.Users.FindOne(ctx.Chat().ID)

	if err != nil {
		log.Printf("Ошибка получения пользователя %v", err)
	}

	if existUser == nil {
		err := bot.Users.Create(newUser)

		if err != nil {
			log.Printf("Ошибка создания пользователя %v", err)
		}
	}

	return ctx.Send("Привет, " + ctx.Sender().FirstName)
}

func (bot *UpgradeBot) SendAll(message string) error {
	users, err := bot.Users.FindAll()

	if err != nil {
		return err
	}

	for _, user := range users {
		user.MessageAmount++
		err := bot.Users.Update(user)
		if err != nil {
			return err
		}
		bot.Bot.Send(&telebot.User{ID: user.ChatId}, message)
	}

	return nil
}

func (bot *UpgradeBot) InfoHandler(ctx telebot.Context) error {
	log.Println("Info Handler called")

	user, err := bot.Users.FindOne(ctx.Chat().ID)

	if err != nil {
		return err
	}


	return ctx.Send("Ваш никнейм: " + user.Name + "\nВаше имя: " + user.FirstName + " " + user.LastName + "\nКоличество сообщений: " + strconv.Itoa(user.MessageAmount) + "\nДата подписки: " + user.CreatedAt.String())
}

func (bot *UpgradeBot) DeleteHandler(ctx telebot.Context) error {
	log.Println("Delete Handler called")

	user, err := bot.Users.FindOne(ctx.Chat().ID)

	if err != nil {
		return err
	}

	err = bot.Users.Delete(*user)

	if err != nil {
		return err
	}

	return ctx.Send("Вы успешно отписались от рассылки")
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return b
}