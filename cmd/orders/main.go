package main

import (
	"fmt"

	"example.com/orders/internal/config"
	"example.com/orders/internal/handler"
	"example.com/orders/internal/repository"
	"example.com/orders/internal/service"
)

func main() {
	repo := repository.NewMemory()
	orders := service.NewService(repo)
	newHandler := handler.NewHandler(orders)

	address := config.GetHTTP_ADDR()
	order, err := newHandler.Create("Покупатель")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Printf("сервис слушал бы %s, создан заказ: %+v\n", address, order)
}
