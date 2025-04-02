package main

import (
  handler "key-haven-back/internal/http/handler"
  "key-haven-back/internal/http/router"
  "log"
  "os"
  "os/exec"

  "key-haven-back/config"
  _ "key-haven-back/docs"
  "key-haven-back/internal/http"
  "key-haven-back/internal/infra/database"
  "key-haven-back/internal/repository"
  "key-haven-back/internal/service"
  "key-haven-back/pkg/docs"

  "github.com/joho/godotenv"
  "go.mongodb.org/mongo-driver/mongo"
  "go.uber.org/fx"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Println("Warning: Error loading .env file")
  }

  cmd := exec.Command("make", "swag")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    log.Printf("Error running make swag: %v", err)
  }
  
  fxopts := []fx.Option{
    // config
    fx.Provide(config.NewConfig),

    // database
    fx.Provide(
      database.NewMongoDBClient,
    ),

    // repository
    fx.Provide(
      func(client database.MongoDBClient) *mongo.Database {
        return client.Database("key-haven")
      },
      repository.NewUserRepository,
      repository.NewVaultRepository,
      repository.NewCredentialRepository,
    ),

    // service
    fx.Provide(
      service.NewUserService,
      service.NewAuthService,
      service.NewVaultService,
      service.NewCredentialService,
    ),

    // http
    fx.Provide(http.NewServer),
    fx.Invoke(http.StartServer),

    // handler
    fx.Provide(
      handler.NewAuthHandler,
      handler.NewVaultHandler,
      handler.NewCredentialHandler,
    ),

    // router
    fx.Provide(
      router.RegisterRoutesFuncProvider,
      router.RegisterSwaggerRoutesFuncProvider,
    ),

    // docs
    fx.Provide(
      docs.RegisterDocsRouterFuncProvider,
    ),
  }

  app := fx.New(fxopts...)
  app.Run()
}
