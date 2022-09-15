package routes

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/myrachanto/microservice/tag/src/cmd/tag"
	"github.com/myrachanto/microservice/tag/src/middle"
)

type Key struct {
	PORT          string `mapstructure:"PORT"`
	EncryptionKey string `mapstructure:"EncryptionKey"`
}

// func LoadKeyConfig() (k Key, err error) {
// 	viper.AddConfigPath(".")
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}
// 	err = viper.Unmarshal(&k)
// 	return
// }

// ApiMicroservice ...
func ApiSingle() {
	// key, err := LoadKeyConfig()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println("Key", key)
	// set a DDD
	// business.Bizs.Initdb()
	t := tag.NewtagController(tag.NewtagService(tag.NewtagRepo()))
	e := echo.New()

	// e.File("/", "public/index.html")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// e.File("/*", "public/index.html")
	e.Static("/", "src/public")

	JWTgroup := e.Group("/api")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in routes")
	}
	////////////////////////////////////////////////////////
	/////////////tags//////////////////////////////////
	JWTgroup.POST("/tags", t.Create, middle.PasetoAuthMiddleware)
	JWTgroup.POST("/tags1", t.Create1, middle.PasetoAuthMiddleware)
	JWTgroup.GET("/tags", t.GetAll, middle.PasetoAuthMiddleware)
	JWTgroup.PUT("/featured/:code", t.Featured, middle.PasetoAuthMiddleware)
	JWTgroup.GET("/tags/:code", t.GetOne, middle.PasetoAuthMiddleware)
	JWTgroup.PUT("/tags/:code", t.Update, middle.PasetoAuthMiddleware)
	JWTgroup.DELETE("/tags/:code", t.Delete, middle.PasetoAuthMiddleware)

	// e.Logger.Fatal(e.Start(":1200"))

	PORT := os.Getenv("PORT")
	// log.Println("fired up .... on port :1200")
	e.Logger.Fatal(e.Start(PORT))
}
