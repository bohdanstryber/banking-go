package app

import (
	"fmt"
	"github.com/bohdanstryber/banking-go/config"
	"github.com/bohdanstryber/banking-go/domain"
	"github.com/bohdanstryber/banking-go/service"
	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

var cnfg config.Config

func Start() {
	err := cleanenv.ReadConfig(".env", &cnfg)
	if err != nil {
		panic("Config file is not defined")
	}

	router := mux.NewRouter()

	dbClient := getDbClient()

	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDB)}

	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")

	router.
		HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")

	router.
		HandleFunc("/customers/{id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")

	router.
		HandleFunc("/customers/{id:[0-9]+}/account/{account_id:[0-9]+}", ah.NewTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	http.ListenAndServe(cnfg.AppUrl, router)
}

func getDbClient() *sqlx.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cnfg.DbUser,
		cnfg.DbPassword,
		cnfg.DbAddress,
		cnfg.DbPort,
		cnfg.DbName)
	client, err := sqlx.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
