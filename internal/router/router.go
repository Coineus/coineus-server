package router

import (
	"os"
	"time"

	"github.com/coineus/coineus-server/internal/api"
	"github.com/coineus/coineus-server/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
)

type Router struct {
	router *fiber.App
	store  *storage.Database
}

func New(s *storage.Database) *Router {
	r := &Router{
		router: fiber.New(),
		store:  s,
	}
	r.initRoutes()
	return r
}

func (r *Router) ServeHTTP(port string) {
	r.router.Listen(":" + port)
}

func (r *Router) initRoutes() {
	//api grouping
	apiRouter := r.router.Group("")

	// Middlewares
	// CORS
	apiRouter.Use(cors.New())
	// Logger
	apiRouter.Use(logger.New(
		logger.Config{
			Format: time.Now().Format("2006-01-02 15:04:05") + " - [${ip}] ${method} ${status} | ${path}\n",
			Output: os.Stdout,
		},
	))

	//api auth
	apiRouter.Post("/register", api.RegisterHandler(r.store.Users))
	apiRouter.Post("/login", api.LoginHandler(r.store.Users))
	r.router.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SIGNING_KEY")),
	}))

	//api healthcheck
	apiRouter.Get("/healthcheck", api.HealtCheckHandler(*r.store))

	//Recent Buy Operations
	apiRouter.Get("/operations/get/:id", api.GetOperationByIdHandler(r.store.RecentOps))
	apiRouter.Get("/operations/getall", api.GetAllOperationsByUserIdHandler(r.store.RecentOps))
	apiRouter.Post("/operations/add", api.AddOperationHandler(r.store.RecentOps))
	apiRouter.Post("/operations/delete", api.DeleteOperationHandler(r.store.RecentOps))
	apiRouter.Post("/operations/update", api.UpdateOperationHandler(r.store.RecentOps))

	//Archived Buy Operations
	apiRouter.Get("/archivedoperations/getall", api.GetAllArchivedOperationsByUserIdHandler(r.store.ArchivedOps))
	apiRouter.Get("/archivedOperations/get/:id", api.AddArchivedOperationHandler(r.store.ArchivedOps))
	apiRouter.Post("/archivedoperations/add", api.AddArchivedOperationHandler(r.store.ArchivedOps))
	apiRouter.Post("/archivedOperations/delete", api.DeleteArchivedOperationHandler(r.store.ArchivedOps))

	//Wallet
	apiRouter.Get("/wallets/getall", api.GetAllWalletsByUserIdHandler(r.store.Wallets))
	apiRouter.Get("/wallets/getbyid", api.GetWalletByIdHandler(r.store.Wallets))
	apiRouter.Post("/wallets/add", api.AddWalletHandler(r.store.Wallets))
	apiRouter.Post("/wallets/delete", api.DeleteWalletHandler(r.store.Wallets, r.store.WalletOps))
	//update eklenecek

	//Wallet Operations
	apiRouter.Get("/walletoperations/getall", api.GetAllWalletOperationsByUserIdHandler(r.store.WalletOps))
	apiRouter.Post("/walletoperations/add", api.AddWalletOperationHandler(r.store.WalletOps))
	apiRouter.Post("/walletoperations/delete", api.DeleteWalletOperationHandler(r.store.WalletOps))

	//User Operations
	apiRouter.Get("/users/getcurrentuser", api.GetUserHandler(r.store.Users))
	apiRouter.Post("/users/update", api.UpdateUserHandler(r.store.Users))
	apiRouter.Post("/users/delete", api.DeleteUserHandler(r.store.Users))
}
