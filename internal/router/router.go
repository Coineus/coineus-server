package router

import (
	"os"

	"github.com/coineus/coineus-server/internal/api"
	"github.com/coineus/coineus-server/internal/storage"
	"github.com/gofiber/fiber/v2"
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

	//api healthcheck
	// apiRouter.Get("/healthcheck")

	//api auth
	apiRouter.Post("/register", api.RegisterHandler(*r.store))
	apiRouter.Post("/login", api.LoginHandler(*r.store))

	r.router.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SIGNING_KEY")),
	}))
	// //Recent Buy Operations
	// apiRouter.Get("/operations/get/id::id")
	// apiRouter.Get("/operations/getall")
	// apiRouter.Post("/operations/add")
	// apiRouter.Post("/operations/delete")
	// apiRouter.Post("/operations/update")

	// //Archived Buy Operations
	// apiRouter.Get("/archivedoperations/getall")
	// apiRouter.Get("/archivedOperations/get/id::id")
	// apiRouter.Post("/archivedoperations/add")
	// apiRouter.Post("/archivedOperations/delete")

	// //Wallet
	// apiRouter.Get("/wallets/getall")
	// apiRouter.Get("/wallets/getbyid")
	// apiRouter.Post("/wallets/add")
	// apiRouter.Post("/wallets/delete")
	// //update eklenecek

	// //Wallet Operations
	// apiRouter.Get("/walletoperations/getall")
	// apiRouter.Post("/walletoperations/add")
	// apiRouter.Post("/walletoperations/delete")

	// //User Operations
	// apiRouter.Get("/users/getcurrentuser")
	// apiRouter.Post("/users/update")
	// apiRouter.Post("/users/delete")

	//Middlewares
	//r.router.Use(LoggerMiddleware)
}
