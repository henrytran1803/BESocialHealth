package main

import (
	accounthandlers "BESocialHealth/Internal/account/handlers"
	exersicehandler "BESocialHealth/Internal/exersice_management/handler"
	foodhandler "BESocialHealth/Internal/food_management/handler"
	personalcontenthandler "BESocialHealth/Internal/personal_content_management/handler"
	mealhandler "BESocialHealth/Internal/personal_meal_management/handler"
	schedulehandler "BESocialHealth/Internal/personal_schedule_management/handler"
	userhandler "BESocialHealth/Internal/user_management/handler"
	"BESocialHealth/component/appctx"
	"BESocialHealth/middleware"
	"github.com/gin-gonic/gin"
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

	//meal
	meal := v1.Group("/meal")
	meal.POST("", mealhandler.CreateMealHandler(appctx))
	meal.GET("/user/:id", mealhandler.GetMealsByUserIdHandler(appctx))
	meal.GET("/:id", mealhandler.GetMealByIdHandler(appctx))
	meal.POST("/detail", mealhandler.CreateMealDetailHandler(appctx))
	meal.PUT("/detail", mealhandler.UpdateMealDetail(appctx))
	meal.DELETE("/:id", mealhandler.DeleteMealById(appctx))
	meal.DELETE("/detail/:id", mealhandler.DeleteDetailMealById(appctx))

	//content
	content := v1.Group("/content")
	content.POST("", personalcontenthandler.CreatePostHandler(appctx))
	content.POST("/like", personalcontenthandler.LikeHandler(appctx))
	content.DELETE("/like", personalcontenthandler.DeleteLikeByUserIdAndPostIdHandler(appctx))
	content.DELETE("/:id", personalcontenthandler.DeletePostHandler(appctx))
	content.POST("/coment", personalcontenthandler.CreatePostHandler(appctx))
	content.PUT("/:id", personalcontenthandler.UpdatePostHandler(appctx))
	content.GET("/:id", personalcontenthandler.GetPostByIdHandler(appctx))
	content.GET("", personalcontenthandler.GetAllPostHandler(appctx))
	content.GET("/coment/:id", personalcontenthandler.GetAllComentByPostIdHandler(appctx))
	// schedule
	schedule := v1.Group("/schedule")
	schedule.POST("", schedulehandler.CreateScheduleHandler(appctx))
	schedule.POST("/detail", schedulehandler.CreateScheduleDetailHandler(appctx))
	schedule.GET("", schedulehandler.GetAllScheduleHandler(appctx))
	schedule.GET("/:id", schedulehandler.GetScheduleByIdHandler(appctx))
	schedule.PUT("", schedulehandler.UpdateScheduleHandler(appctx))
	schedule.PUT("/detail", schedulehandler.UpdateScheduleDetailHandler(appctx))
	schedule.DELETE("/:id", schedulehandler.DeleteScheduleHandler(appctx))
	schedule.DELETE("/detail/:id", schedulehandler.DeleteScheduleDetailHandler(appctx))
	//lay toan bo danh sach theo id nguoi dung
	// them schedule moi
	// them detail moi
	// update 2 cai
	// lay chi tiet theo id
	// xoa tung cai xoa het

	// xoa coment ca nhatcoment
	// WebSocket routes
	r.GET("/ws/admin", func(c *gin.Context) {
		handleAdminConnections(c.Writer, c.Request)
	})

	r.GET("/ws/user", func(c *gin.Context) {
		handleUserConnections(c.Writer, c.Request)
	})

	go handleAdminMessages()
	go handleUserMessages()
	//ctx := context.Background()
	//sa := option.WithCredentialsFile("./cmd/beorderfood-de62b3f3f8d0.json")
	//app, err := firebase.NewApp(ctx, nil, sa)
	//if err != nil {
	//	log.Fatalf("error initializing app: %v\n", err)
	//}
	//
	//client, err := app.Firestore(ctx)
	//if err != nil {
	//	log.Fatalf("error getting Firestore client: %v\n", err)
	//}
	//defer client.Close()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
