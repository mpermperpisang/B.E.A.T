package command

import (
	"gopkg.in/tucnak/telebot.v2"
)

/*Main : Kumpulan command bot*/
func Main(bot *telebot.Bot, m *telebot.Message) {
	Mention(bot, m)
}
