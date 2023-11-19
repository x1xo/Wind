package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		IP                 string
		Port               uint16
		LimiterMaxRequests int
		LimiterDuration  time.Duration
	}
	Discord struct {
		Token    string
		Guild_ID string
	}
	Database struct {
		Type string
		URL  string
	}
}

var config *Config

func GetConfig() *Config {
	return config
}

func ParseConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config = &Config{}

	portString := os.Getenv("SERVER_PORT")
	portInt, err := strconv.Atoi(portString)
	if err != nil {
		panic("Port is not a number")
	}
	if portInt < 1 || portInt > 65535 {
		panic("Port must be between 1 and 65535")
	}

	config.Server.Port = uint16(portInt)
	config.Server.IP = os.Getenv("SERVER_IP")

	limiter_max := os.Getenv("LIMITER_MAX_REQUESTS")
	limiter_max_num, err := strconv.Atoi(limiter_max)
	if err != nil {
		panic("LIMITER_MAX_REQUESTS must be a number")
	}
	
	config.Server.LimiterMaxRequests = limiter_max_num

	limiter_duration := os.Getenv("LIMITER_DURATION");
	limiter_duration_num, err := strconv.Atoi(limiter_duration);
	if err != nil {
		panic("LIMITER_EXPIRATION must be a duration in miliseconds")
	}

	config.Server.LimiterDuration = time.Duration(limiter_duration_num)*time.Millisecond

	config.Discord.Token = os.Getenv("DISCORD_TOKEN")
	config.Discord.Guild_ID = os.Getenv("DISCORD_GUILD_ID")
	//Database

	dbType := os.Getenv("DATABASE_TYPE")
	if dbType != "redis" && dbType != "sqlite" {
		panic("Invalid database type. (Choose between `sqlite` and `redis`)")
	}

	config.Database.Type = dbType
	config.Database.URL = os.Getenv("DATABASE_URL")

	return config, nil
}
