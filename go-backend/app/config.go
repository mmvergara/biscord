package app

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	MongoDBAddr string
	PORT        uint16
}

func LoadConfig() Config {
	// Provide default values
	cfg := Config{
		MongoDBAddr: "mongodb://localhost:27017",
		PORT:        8080,
	}

	if MonogoDBAddr, exists := os.LookupEnv("MONGODB_ADDR"); exists {
		cfg.MongoDBAddr = MonogoDBAddr
	}

	if portStr, exists := os.LookupEnv("PORT"); exists {
		if PORT,err := strconv.ParseUint(portStr, 10, 16); err != nil {
			log.Fatalf("Invalid PORT value: %v", err)
		} else {
			cfg.PORT = uint16(PORT)
		}
	}

	return cfg
}