package config

import (
	"os"
	"time"
)

type Config struct {
	APP         string
	Environment string
	LogLevel    string
	RPCPort     string
	RabbitMQURL string
	Context     struct {
		Timeout string
	}
	HotelUrl   string
	UserUrl    string
	BookingUrl string
	Token      struct {
		Secret     string
		AccessTTL  time.Duration
		RefreshTTL time.Duration
	}
	RedisURL string
	DB       struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
	}
	NotificationUrl string

	MessageBrokerUses struct {
		URL          string
		Topic        string
		TopicBooking string
		Keys         struct {
			Create         []byte
			Update         []byte
			Delete         []byte
			UpdateEmail    []byte
			UpdatePassword []byte
		}
		KeysBooking struct {
			CreateOrder  []byte
			AddWaitGroup []byte
		}
	}
}

func Token() string {
	c := Config{}
	c.Token.Secret = getEnv("TOKEN_SECRET", "token_secret")
	return c.Token.Secret
}

func New() *Config {
	var config Config

	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "prod")
	config.RPCPort = getEnv("RPC_PORT", "api_gateway:9006")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	config.Token.Secret = getEnv("TOKEN_SECRET", "D1YORTOP4EEK")
	accessTTl, err := time.ParseDuration(getEnv("TOKEN_ACCESS_TTL", "1h"))
	if err != nil {
		return nil
	}
	refreshTTL, err := time.ParseDuration(getEnv("TOKEN_REFRESH_TTL", "24h"))
	if err != nil {
		return nil
	}
	config.Token.AccessTTL = accessTTl
	config.Token.RefreshTTL = refreshTTL

	config.RabbitMQURL = getEnv("RabbitMQ_URL", "amqp://guest:guest@localhost:5672/")

	config.UserUrl = getEnv("User_URL", "user_service_container:9000")
	config.NotificationUrl = getEnv("Notification_URL", "notification_service_container:9001")
	config.BookingUrl = getEnv("Booking_URL", "booking_service_container:9003")
	config.HotelUrl = getEnv("Hotel_URL", "hotel_service_container:9004")
	config.RedisURL = getEnv("REDIS_URL", "redis:6379")

	config.MessageBrokerUses.URL = getEnv("KAFKA_URL", "broker:29092")
	config.MessageBrokerUses.Topic = getEnv("MESSAGE_BROKER_USE_TOKEN", "USER_SERVICE")
	config.MessageBrokerUses.Keys.Create = []byte(getEnv("MESSAGE_BROKER_USE_KEY", "CREATE"))
	config.MessageBrokerUses.Keys.Delete = []byte(getEnv("MESSAGE_BROKER_USE_KEY", "DELETE"))
	config.MessageBrokerUses.Keys.Update = []byte(getEnv("MESSAGE_BROKER_USE_KEY", "UPDATE"))
	config.MessageBrokerUses.Keys.UpdateEmail = []byte(getEnv("MESSAGE_BROKER_USE_KEY", "UPDATE_EMAIL"))
	config.MessageBrokerUses.Keys.UpdatePassword = []byte(getEnv("MESSAGE_BROKER_USE_KEY", "UPDATE_PASSWORD"))
	config.MessageBrokerUses.TopicBooking = getEnv("MESSAGE_BROKER_USE_TOKEN", "USER_SERVICE")
	config.MessageBrokerUses.KeysBooking.CreateOrder = []byte(getEnv("MESSAGE_BROKER_USE_KEY", "CREATE_ORDER"))
	config.MessageBrokerUses.KeysBooking.AddWaitGroup = []byte(getEnv("MESSAGE_BROKER_USE_KEY", "ADD_WAIT"))
	return &config
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
