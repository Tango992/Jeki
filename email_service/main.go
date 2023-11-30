package main

import (
	"email-service/config"
	"email-service/service"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conn, mbChannel := config.InitMbChannel()
	defer conn.Close()

	q := config.InitMbQueue(mbChannel)
	mailService := service.NewMailService(mbChannel)
	
	go mailService.SendVerificationEmail(q)
	
	var forever chan struct{}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}