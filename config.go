package main

var config Config

type Config struct {
	ContentMinLength        int
	MaxSleepTimeMinutes     int
	MaxRetryConnectAttempts int
}

func init() {
	config = Config{
		ContentMinLength:        2,
		MaxSleepTimeMinutes:     60,
		MaxRetryConnectAttempts: 50,
	}
}
