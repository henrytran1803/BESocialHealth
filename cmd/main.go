package main

import (
	accounthandlers "BESocialHealth/Internal/account/handlers"
	exersicehandler "BESocialHealth/Internal/exersicemanager/handler"
	foodhandler "BESocialHealth/Internal/food_management/handler"
	userhandler "BESocialHealth/Internal/user_management/handler"
	"BESocialHealth/component/appctx"
	"BESocialHealth/middleware"
	"context"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/api/option"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	if dsn == "" {
		log.Fatal("MYSQL_CONN_STRING environment variable not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	//db.AutoMigrate(&usermodel.User{})
	r := gin.Default()
	appctx := appctx.NewAppContext(db)
	r.Use(middleware.Recover(appctx))
	log.Println("Role found:")
	v1 := r.Group("/v1")
	// account
	account := v1.Group("/account")
	account.POST("/register", accounthandlers.CreateAccountHandler(appctx))
	account.POST("/login", accounthandlers.LoginHandler(appctx))
	account.POST("/requestpassword", accounthandlers.RequestPasswordResetHandler(appctx))
	account.POST("/confirmpassword", accounthandlers.ConfirmPasswordResetHandler(appctx))

	// food
	food := v1.Group("/food")
	food.POST("", foodhandler.CreateFoodHandler(appctx))
	food.PUT("/:id", foodhandler.UpdateFoodHandler(appctx))
	food.DELETE("/:id", foodhandler.DeleteFoodHandler(appctx))
	food.GET("", foodhandler.GetListFoodHandler(appctx))
	food.GET("/:id", foodhandler.GetFoodHandler(appctx))

	//exersice
	exersice := v1.Group("/exersice")
	exersice.POST("", exersicehandler.CreateExersiceHandler(appctx))
	exersice.PUT("/:id", exersicehandler.UpdateExersiceHandeler(appctx))
	exersice.DELETE("/:id", exersicehandler.DeleteExersiceHandler(appctx))
	exersice.GET("", exersicehandler.GetistExersiceHandler(appctx))
	exersice.GET("/:id", exersicehandler.GetExersiceByIdHandler(appctx))

	//user
	user := v1.Group("/user")
	user.GET("", userhandler.GetAllUserHandler(appctx))
	user.GET("/:id", userhandler.GetUserByIdHandler(appctx))
	user.POST("", userhandler.CreateUserHandler(appctx))
	user.PUT("/:id", userhandler.UpdateUserHandler(appctx))
	user.DELETE("/:id", userhandler.DeleteUserHandler(appctx))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ctx := context.Background()
	sa := option.WithCredentialsFile("./cmd/beorderfood-de62b3f3f8d0.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}
	defer client.Close()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
