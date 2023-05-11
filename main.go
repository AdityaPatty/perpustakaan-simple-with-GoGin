package main

import (
	"log"

	"perpustakaan/handlerbuku"
	"perpustakaan/repositorybuku"
	"perpustakaan/servicebuku"

	"github.com/gin-gonic/gin"
)

func main() {
	// separate the code from the 'main' function so we can test it.
	// all code that available in main function were not testable
	Run()
}

func Run() {
	gin.SetMode(gin.ReleaseMode)

	konek := gin.New()
	konek.Use(gin.Recovery())

	dbPool, _, err := servicebuku.NewDBPool(servicebuku.DatabaseConfig{
		Username: "postgres",
		Password: "Patty12345",
		Hostname: "localhost",
		Port:     "5432",
		DBName:   "perbukuan",
	})

	defer dbPool.Close()

	if err != nil {
		log.Fatalf("unexpected error while tried to connect to database: %v\n", err)
	}

	bukDB := repositorybuku.NewDatabase(dbPool)
	bukService := servicebuku.NewBukuService(bukDB)
	bukAPI := handlerbuku.NewBukuHandler(bukService)

	v1 := konek.Group("/v1")

	accRouter := v1.Group("/buku")
	accRouter.POST("/", bukAPI.PerpustakaanCreateHandler)
	accRouter.PUT("/:id", bukAPI.PerpustakaanUpdateHandler)
	accRouter.DELETE("/:id", bukAPI.PerpustakaanDeleteHandler)
	accRouter.GET("/:id", bukAPI.PerpustakaanGetHandler)
	accRouter.GET("/", bukAPI.PerpustakaanGetsHandler)

	// book app group api endpoint : http://domainname.com/v1/book
	// accRouter = v1.Group("/book")
	// accRouter.POST("/", accAPI.BookCreateHandler)
	// accRouter.PUT("/:id", accAPI.BookUpdateHandler)
	// accRouter.DELETE("/:id", accAPI.BookDeleteHandler)
	// accRouter.GET("/:id", accAPI.BookGetHandler)
	// accRouter.GET("/", accAPI.BookGetsHandler)

	//run the server
	log.Fatalf("%v", konek.Run())
}
