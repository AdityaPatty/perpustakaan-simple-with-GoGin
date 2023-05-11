package main

import (
	"log"

	"perpustakaan/handlerbuku"
	"perpustakaan/repositorybuku"
	"perpustakaan/servicebuku"

	"github.com/gin-gonic/gin"
)

func main() {

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

	//run the server
	log.Fatalf("%v", konek.Run())
}
