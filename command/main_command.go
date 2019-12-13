package command

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/tucnak/telebot.v2"
	yaml "gopkg.in/yaml.v2"
)

var data string
var beat BEAT

/*BEAT : kumpulan variabel untuk file yaml*/
type BEAT struct {
	Members struct {
		TE  string `yaml:"TE"`
		ATE string `yaml:"ATE"`
		ATA string `yaml:"ATA"`
		TEM string `yaml:"TEM"`
	} `yaml:"Members"`
}

/**/
func readMember() {
	reader, _ := os.Open("./data/beat.yml")
	buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &beat)
}

/*Main : Kumpulan command bot*/
func Main(bot *telebot.Bot, m *telebot.Message) {
	var contentMessage string

	readMember()

	switch strings.ToUpper(m.Text) {
	case "OI (A)TE":
		contentMessage = beat.Members.ATE + " " + beat.Members.TE
	case "OI TE":
		contentMessage = beat.Members.TE
	case "OI ATE":
		contentMessage = beat.Members.ATE
	case "OI ATA":
		contentMessage = beat.Members.ATA
	case "OI TEM":
		contentMessage = beat.Members.TEM
	}

	bot.Reply(m, contentMessage, telebot.ModeHTML)
}
