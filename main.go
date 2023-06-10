package main

import (
	"log"
	"shopee-integration-service/src/service"
)

func main() {
	log.Printf("Application is running")

	productService := service.NewProductService()
	productService.ConsumeProductMessages()
}
