package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "go-stripe/internal/driver"
    "go-stripe/internal/models"
)

type application struct {
    config   Config
    infoLog  *log.Logger
    errorLog *log.Logger
    version  string
    DB       models.DBModel
}

func (app *application) serve() error {
    srv := &http.Server{
        Addr:              fmt.Sprintf(":%d", app.config.Application.Port),
        Handler:           app.routes(),
        IdleTimeout:       30 * time.Second,
        ReadTimeout:       10 * time.Second,
        ReadHeaderTimeout: 5 * time.Second,
        WriteTimeout:      5 * time.Second,
    }

    app.infoLog.Printf("Starting Back end server in %s mode on %s:%d\n",
        app.config.Application.Environment, app.config.Application.HostName, app.config.Application.Port)

    return srv.ListenAndServe()
}

func main() {

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // load ENV variables
    config, err := LoadEnv()
    if err != nil {
        errorLog.Fatal(err)
    }
    infoLog.Println("ENV variable parse: OK")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&tls=%s",
        config.Storage.MySQL.Username,
        config.Storage.MySQL.Password,
        config.Storage.MySQL.Host,
        config.Storage.MySQL.Port,
        config.Storage.MySQL.Name,
        config.Storage.MySQL.SSLMode,
    )

    conn, err := driver.OpenDB(dsn)
    if err != nil {
        errorLog.Fatal(err)
    }
    defer conn.Close()
    infoLog.Println("DB pinged and responding")

    app := &application{
        config:   config,
        infoLog:  infoLog,
        errorLog: errorLog,
        version:  config.Application.Version,
        DB:       models.DBModel{DB: conn},
    }

    err = app.serve()
    if err != nil {
        errorLog.Fatal(err)
    }
}
