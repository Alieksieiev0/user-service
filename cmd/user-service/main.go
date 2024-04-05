package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Alieksieiev0/user-service/internal/database"
	"github.com/Alieksieiev0/user-service/internal/services"
	"github.com/Alieksieiev0/user-service/internal/transport/grpc"
	"github.com/Alieksieiev0/user-service/internal/transport/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(111111)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Setup(db)
	if err != nil {
		log.Fatal(err)
	}

	var (
		r = flag.String("rest", ":3000", "listen address of json server")
		g = flag.String("grpc", ":4000", "listen address of grpc grpc")
	)
	service := services.NewUserService(db)
	app := fiber.New()

	grpcServer := grpc.NewServer()
	go grpcServer.Start(*g, service)
	/*
		if err != nil {
			log.Fatal(err)
		}
	*/

	restServer := rest.NewServer(app)
	err = restServer.Start(*r, service)
	if err != nil {
		log.Fatal(err)
	}

	/*
		client, err := grpc.NewGRPCClient(*g)
		go func() {
			time.Sleep(3 * time.Second)
			client.GetByUsername(context.Background(), &proto.UsernameRequest{Username: "test"})
		}()
	*/
}
