package initializers

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBUserName     string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	TimeZone       string `mapstructure:"TIME_ZONE"`

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	FrontendURL string `mapstructure:"FRONTEND_URL"`

	EnableEmailVerification bool `mapstructure:"ENABLE_EMAIL_VERIFICATION"`

	MailHost     string `mapstructure:"MAIL_HOST"`
	MailPort     string `mapstructure:"MAIL_PORT"`
	MailUsername string `mapstructure:"MAIL_USERNAME"`
	MailPassword string `mapstructure:"MAIL_PASSWORD"`
	MailFrom     string `mapstructure:"MAIL_FROM"`
}

func GetEnvConfig() *Config {
	godotenv.Load(".env")

	enableEmailVerification, _ := strconv.ParseBool(os.Getenv("ENABLE_EMAIL_VERIFICATION"))

	return &Config{
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUserPassword: os.Getenv("DB_PASS"),
		DBUserName:     os.Getenv("DB_USER"),
		// DBSSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
		TimeZone: os.Getenv("TIME_ZONE"),

		EnableEmailVerification: enableEmailVerification,

		FrontendURL: os.Getenv("FRONTEND_URL"),

		MailHost:     os.Getenv("MAIL_HOST"),
		MailPort:     os.Getenv("MAIL_PORT"),
		MailUsername: os.Getenv("MAIL_USERNAME"),
		MailPassword: os.Getenv("MAIL_PASSWORD"),
		MailFrom:     os.Getenv("MAIL_FROM"),
	}
}
