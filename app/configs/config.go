package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	JWT_SECRET string
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
	// JWT_SECRET  string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	// inisialisasi variabel dg type struct AppConfig
	app := AppConfig{}
	isRead := true

	// proses mencari & membaca environment var dg key tertentu
	// if val, found := os.LookupEnv("JWT_KEY"); found {
	// 	app.JWT_SECRET = val
	// 	isRead = false
	// }
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("JWTSECRET"); found {
		app.DB_NAME = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		// app.JWT_SECRET = viper.Get("JWT_KEY").(string)
		app.DB_USERNAME = viper.Get("DBUSER").(string)
		app.DB_PASSWORD = viper.Get("DBPASS").(string)
		app.DB_HOSTNAME = viper.Get("DBHOST").(string)
		app.DB_PORT, _ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DB_NAME = viper.Get("DBNAME").(string)
		JWT_SECRET = viper.Get("JWTSECRET").(string)
	}

	return &app
}
