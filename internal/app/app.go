package app

import (
	"context"
	"database/sql"
	"fmt"

	"log"

	srv "github.com/atadzan/music-library-api/internal/app/server"
	"github.com/atadzan/music-library-api/internal/pkg/handler"
	"github.com/atadzan/music-library-api/internal/pkg/services"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/atadzan/music-library-api/internal/config"
	"github.com/atadzan/music-library-api/internal/pkg/repository"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func Init() {
	ctx := context.Background()
	appCfg, err := config.LoadConfig()
	if err != nil {
		log.Println(ctx, "can't load config. Err: ", err)
		return
	}

	dbPool, err := repository.NewPostgresPool(ctx, repository.Params{
		Username: appCfg.DbUsername,
		Password: appCfg.DbPassword,
		Host:     appCfg.DbHost,
		Port:     appCfg.DbPort,
		DbName:   appCfg.DbName,
		SslMode:  appCfg.DbSSLMode,
	})
	if err != nil {
		log.Println(ctx, "can't init postgres. Err: ", err)
		return
	}
	defer dbPool.Close()

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", appCfg.DbUsername, appCfg.DbPassword,
		appCfg.DbHost, appCfg.DbPort, appCfg.DbName, appCfg.DbSSLMode))
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		log.Println(err)
		return
	}
	m.Up() //

	if err != nil {
		log.Println(ctx, "can't init migration. Err: ", err)
		return
	}
	if err = m.Up(); err != nil {
		log.Println(ctx, "can't apply migration. Err: ", err)
	}

	repo := repository.New(dbPool)
	svc := services.NewService(repo)
	handlerx := handler.New(svc)
	routes := handlerx.InitRoutes()
	srv := srv.New(appCfg.HTTPPort, routes)

	if err = srv.Run(); err != nil {
		log.Fatalln("can't run server. Err:", err)
	}
}
