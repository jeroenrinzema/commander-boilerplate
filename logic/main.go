package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander-boilerplate/logic/controllers"
	"github.com/jeroenrinzema/commander-boilerplate/logic/groups"
	"github.com/jeroenrinzema/commander-boilerplate/logic/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")

	options := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, database, password)
	db, err := gorm.Open("postgres", options)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.CartModal{})

	config := commander.NewConfig()
	config.Brokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	config.Group = os.Getenv("KAFKA_GROUP")
	config.Kafka = &kafka.ConfigMap{
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
	}

	config.AddGroups(groups.Cart)

	client, err := commander.New(config)
	if err != nil {
		panic(err)
	}

	groups.Cart.CommandHandle("Create", &controllers.CreateCommand{Database: db})
	groups.Cart.CommandHandle("Purchase", &controllers.PurchaseCommand{Database: db})

	go client.Consume()
	client.CloseOnSIGTERM()
}
