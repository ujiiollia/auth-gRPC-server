package main

import "app/internal/config"

func main() {
	//инициализирован конфиг
	cfg := config.MustLoad()
	_ = cfg
	//todo: инициализировать логгер
	//todo: инициализировать логику
	//todo: запустить сервер
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)
