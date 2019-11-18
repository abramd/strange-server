package main

import (
	"github.com/abramd/strange-server/internal/config"
	"github.com/abramd/strange-server/internal/model"
	"github.com/abramd/strange-server/internal/process"
	"github.com/abramd/strange-server/internal/server"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	dsn := os.Getenv("DSN")
	time.Sleep(time.Second * 3)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if config.Cfg.WithMigrations {
		m, err := migrate.New("file:///migrations", dsn)
		if err != nil {
			panic(err)
		}
		err = m.Up()
		if err != nil {
			panic(err)
		}
	}

	m := model.NewManager(db)
	sourceMap, err := server.NewSourceMap(m)
	if err != nil {
		panic(err)
	}

	go process.UpdateProcessed(m, config.Cfg.EntityProcessDuration)
	go process.UpdateStateDeleted(m, config.Cfg.EntityDeleteDuration)

	r := mux.NewRouter()
	r.HandleFunc("/post", server.SourceValidationMW(sourceMap, server.PostHandler(m))).Methods(http.MethodPost)
	r.HandleFunc("/list", server.ListHandler(m)).Methods(http.MethodGet)
	r.HandleFunc("/source_list", server.SourceListHandler(m)).Methods(http.MethodGet)
	log.Println("server started")
	log.Fatal(http.ListenAndServe(config.Cfg.Port, r))
}
