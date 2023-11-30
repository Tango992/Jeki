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
	registerMailService := service.NewRegisterMail(mbChannel)
	
	go registerMailService.SendEmail(q)
	
	var forever chan struct{}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}