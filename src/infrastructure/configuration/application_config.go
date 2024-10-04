package configuration

import (
	"apiblog/src/external/routes"
	"apiblog/src/infrastructure/commons/logger"
	"apiblog/src/infrastructure/database"
	"apiblog/src/infrastructure/interceptor"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	Porta   = 8080
	Context = ""
)

func StartConfiguration() {
	loadLogger()
	loadEnv()
	loadDatabase()
	loadServer()
}

func loadEnv() {
	if erro := godotenv.Load(); erro != nil {
		panic("Error ao carregar as variáveis de ambiente!")
	}
	log.Printf("Variáveis de ambiente carregadas com sucesso!")
}

func registerRoutes(routers *mux.Router) *mux.Router {
	rotas := routes.PublicationRoutes

	for _, rota := range rotas {
		routers.HandleFunc(rota.URI, interceptor.Logger(rota.Func)).Methods(rota.Metodo)
	}
	log.Print("Rotas HTTP Registradas com sucesso!")
	return routers
}

func loadLogger() {
	custom_log := slog.New(logger.NewHandler(nil))
	slog.SetDefault(custom_log)
	log.Print("Logger Carregado com sucesso!")
}

func loadServer() {
	routers := registerRoutes(mux.NewRouter())
	log.Printf("Servidor iniciado na porta %d", Porta)
	if erro := http.ListenAndServe(fmt.Sprintf(":%d", Porta), routers); erro != nil {
		panic(fmt.Sprintf("Error ao iniciar servidor %s", erro.Error()))
	}
}

func loadDatabase() {
	db_name, _ := os.LookupEnv("DATABASE_NAME")
	db_host := os.Getenv("DATABASE_HOST")
	db_port := os.Getenv("DATABASE_PORT")
	db_user := os.Getenv("DATABASE_USER")
	db_password := os.Getenv("DATABASE_PASSWORD")

	database.SetDatabaseEnv(db_name, db_host, db_port, db_user, db_password)
	if _, erro := database.Conectar(); erro != nil {
		panic("Não foi possivel conectar-se com o banco de dados")
	}

	log.Println("Conexão com Banco de dados Estabelecida")

	database.InitalStrucuture()
}
