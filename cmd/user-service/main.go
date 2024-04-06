package main

import (
	"flag"
	"log"

	"github.com/Alieksieiev0/user-service/internal/database"
	"github.com/Alieksieiev0/user-service/internal/services"
	"github.com/Alieksieiev0/user-service/internal/transport/grpc"
	"github.com/Alieksieiev0/user-service/internal/transport/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		restServerAddr = flag.String("rest-server", ":3000", "listen address of rest server")
		grpcServerAddr = flag.String("grpc-server", ":4000", "listen address of grpc server")
		app            = fiber.New()
		g              = new(errgroup.Group)
	)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Start()
	if err != nil {
		log.Fatal(err)
	}

	service := services.NewUserService(db)

	grpcServer := grpc.NewServer()
	g.Go(func() error {
		return grpcServer.Start(*grpcServerAddr, service)
	})

	restServer := rest.NewServer(app)
	g.Go(func() error {
		return restServer.Start(*restServerAddr, service)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
