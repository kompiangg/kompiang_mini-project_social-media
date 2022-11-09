package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database
	JWTConfig
	Server
	Admin
	Cloudinary
	UploadFolderName string
}

var config *Config

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[INFO] The .env file doesn't exist")
		log.Println("[INFO] Program will load environment variable value")
	}

	config = &Config{
		Database: Database{
			Username:     os.Getenv("DB_USERNAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			Hostname:     os.Getenv("DB_HOSTNAME"),
			Port:         os.Getenv("DB_PORT"),
			DatabaseName: os.Getenv("DB_NAME"),
		},
		JWTConfig: JWTConfig{
			JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		},
		Server: Server{
			Port: os.Getenv("APP_PORT"),
		},
		Admin: Admin{
			SecretToken: os.Getenv("ADMIN_SECRET_KEY"),
			Username:    os.Getenv("ADMIN_USERNAME"),
			Password:    os.Getenv("ADMIN_PASSWORD"),
		},
		Cloudinary: Cloudinary{
			APIKey:    os.Getenv("API_KEY"),
			APISecret: os.Getenv("API_SECRET"),
			CloudName: os.Getenv("CLOUD_NAME"),
		},
		UploadFolderName: "uploads/",
	}

	if config.Database.Username == "" {
		log.Fatal("[ERROR] Error while init database, database name cant be empty")
	}

	if config.Database.Password == "" {
		log.Fatal("[ERROR] Error while init database, database password cant be empty")
	}

	if config.Database.Hostname == "" {
		log.Fatal("[ERROR] Error while init database, database hostname cant be empty")
	}

	if config.Database.Port == "" {
		log.Fatal("[ERROR] Error while init database, database port cant be empty")
	}

	if config.Database.DatabaseName == "" {
		log.Fatal("[ERROR] Error while init database, database name cant be empty")
	}

	if config.Cloudinary.APISecret == "" {
		log.Fatal("[ERROR] Error while init cloudinary, api secret cant be empty")
	}

	if config.Cloudinary.APIKey == "" {
		log.Fatal("[ERROR] Error while init cloudinary, api key cant be empty")
	}

	if config.Cloudinary.CloudName == "" {
		log.Fatal("[ERROR] Error while init cloudinary, cloud name cant be empty")
	}

	if config.JWTConfig.JWTSecretKey == "" {
		log.Fatal("[ERROR] Error while init jwt config, jwt secret key cant be empty")
	}

	if config.JWTConfig.JWTSecretKey == "" {
		log.Fatal("[ERROR] Error while init jwt config, jwt secret key cant be empty")
	}

	if config.Server.Port == "" {
		log.Fatal("[ERROR] Error while init web server, app port cant be empty")
	}

	if config.Admin.SecretToken == "" {
		log.Fatal("[ERROR] Error while getting admin secret token, admin secret token cant be empty")
	}

	if config.Admin.Username == "" {
		config.Admin.Username = "root"
		log.Println("[DANGEROUS] Admin username is empty, it will set to default username and its dangerous")
	}

	if config.Admin.Username == "" {
		config.Admin.Username = "secret"
		log.Println("[DANGEROUS] Admin password is empty, it will set to default username and its dangerous")
	}
}

func GetConfig() *Config {
	if config == nil {
		initConfig()
	}
	return config
}
