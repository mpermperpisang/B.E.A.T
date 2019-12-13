all: package

package:
	@go get -u gopkg.in/tucnak/telebot.v2
	@go get github.com/joho/godotenv
	@go get gopkg.in/yaml.v2
	@echo "Package installed"
