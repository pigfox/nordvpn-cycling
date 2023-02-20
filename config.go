package main

var config Config

type Config struct {
	ContentMinLength    int
	MaxSleepTimeMinutes int
	MaxConnectAttempts  int
}

func init() {
	config = Config{
		ContentMinLength:    2,
		MaxSleepTimeMinutes: 60,
		MaxConnectAttempts:  50,
	}
}
