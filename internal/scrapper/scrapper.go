package scrapper

import "image-scrapper/internal/config"

var appConfig *config.AppConfig

func Init(config *config.AppConfig) {
	appConfig = config
}

func Run() {

}
