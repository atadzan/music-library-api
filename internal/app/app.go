package app

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"log"

	"github.com/atadzan/music-library-api/internal/config"
	"github.com/atadzan/music-library-api/internal/pkg/repository"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

/* Init TODO
    1. Обогащенную информацию положить в БД postgres
		(структура БД должна быть создана путем миграций при старте сервиса)
	2. Покрыть код debug- и info-логами
	3. Вынести конфигурационные данные в .env-файл
	4. Сгенерировать сваггер на реализованное АПИ
*/

func Init() {
	ctx := context.Background()
	// TODO init config from.env file
	// TODO db init
	// TODO apply migration
	appCfg, err := config.LoadConfig()
	if err != nil {
		log.Println(ctx, "can't load config. Err: ", err)
		return
	}
	log.Println(appCfg)
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
	if err := m.Up(); err != nil {
		log.Println(ctx, "can't apply migration. Err: ", err)
		return
	}

	//if err = srv.Run(":"+appConfig.HTTP.Port, handlers.InitRoutes()); err != nil {
	//	errorx.PrintDetailedError(err)
	//	return err
	//}

}
