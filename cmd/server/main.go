package main

import (
	"fmt"
	"net/http"

	"github.com/AinuroV/az.auth/internal/health"
)

const (
	serviceName = "az.auth"
	version     = "1.0.0"
)

func main() {
	
	healthHandler := health.NewHandler(serviceName, version)

	
	http.HandleFunc("/health", healthHandler.ServeHTTP)

	fmt.Println("Сервер запущен на 80 порту")
	fmt.Println("Проверьте здоровье: http://localhost:8080/health")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}