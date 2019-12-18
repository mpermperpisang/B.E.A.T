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

func readMember() {
	reader, _ := os.Open("./data/beat.yml")
	buf, _ := ioutil.ReadAll(reader)

	yaml.Unmarshal(buf, &beat)
}

/*Mention : Fungsi untuk mention member BEAT*/
func Mention(bot *telebot.Bot, m *telebot.Message) {
	var contentMessage string

	readMember()

	initial := "Dicariin nih, wankawan~~~ colek \n"

	switch strings.ToUpper(m.Text) {
	case "OI (A)TE":
		contentMessage = initial + beat.Members.ATE + beat.Members.TE
	case "OI TE":
		contentMessage = initial + beat.Members.TE
	case "OI ATE":
		contentMessage = initial + beat.Members.ATE
	case "OI ATA":
		contentMessage = initial + beat.Members.ATA
	case "OI TEM":
		contentMessage = initial + beat.Members.TEM
	case "OI ALL":
		contentMessage = initial + beat.Members.TEM + beat.Members.ATA + beat.Members.TE + beat.Members.ATE
	}

	bot.Reply(m, contentMessage, telebot.ModeHTML)
}
