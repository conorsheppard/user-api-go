package api

import (
	"database/sql"
	"github.com/conorsheppard/user-api-go/internal/api/controller"
	"github.com/conorsheppard/user-api-go/internal/api/service/impl"
	db "github.com/conorsheppard/user-api-go/internal/db/sqlc"
	"github.com/conorsheppard/user-api-go/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

var (
	conn             = setUpDB()
	store            = db.NewStore(conn)
	healthService    = impl.NewHealthService()
	healthController = controller.NewHealthController(healthService)
	userService      = impl.NewUserService(store)
	userController   = controller.NewUserController(userService)
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	server := &Server{router: router}
	setUpRoutes(router)
	server.router = router

	return server
}

func setUpRoutes(router *gin.Engine) {
	healthController.SetUpHealthRoute(router)
	userController.SetUpUserRoutes(router)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func setUpDB() *sql.DB {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	return conn
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")
}
