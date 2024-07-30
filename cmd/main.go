package main

import (
	accounthandlers "BESocialHealth/Internal/account/handlers"
	exersicehandler "BESocialHealth/Internal/exersice_management/handler"
	foodhandler "BESocialHealth/Internal/food_management/handler"
	messagehandler "BESocialHealth/Internal/messaging/handler"
	personalcontenthandler "BESocialHealth/Internal/personal_content_management/handler"
	mealhandler "BESocialHealth/Internal/personal_meal_management/handler"
	schedulehandler "BESocialHealth/Internal/personal_schedule_management/handler"
	reminderhandler "BESocialHealth/Internal/reminder_management/handler"
	userhandler "BESocialHealth/Internal/user_management/handler"
	"BESocialHealth/component/appctx"
	"BESocialHealth/component/ws"
	"BESocialHealth/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
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
	r := gin.Default()
	appctx := appctx.NewAppContext(db)
	v1 := r.Group("/v1")

	//ws

	manager := ws.NewWebSocketManager()
	r.GET("/ws", func(c *gin.Context) {
		manager.WebSocketHandler(c.Writer, c.Request)
	})
	r.POST("/broadcast", func(c *gin.Context) {
		var message struct {
			Message string `json:"message"`
		}
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		manager.BroadcastToAll(message.Message)
		c.JSON(http.StatusOK, gin.H{"status": "Message broadcasted"})
	})
	// Reminder Checker
	reminderChecker := ws.NewReminderChecker(appctx, manager, 1*time.Minute) // Kiểm tra mỗi phút một lần
	reminderChecker.Start()

	// account
	account := v1.Group("/account")
	account.POST("/register", accounthandlers.CreateAccountHandler(appctx))
	account.POST("/login", accounthandlers.LoginHandler(appctx))
	account.POST("/requestpassword", accounthandlers.RequestPasswordResetHandler(appctx))
	account.POST("/confirmpassword", accounthandlers.ConfirmPasswordResetHandler(appctx))
	account.GET("/dashboard", accounthandlers.DashBoardHandler(appctx, manager))

	// food
	food := v1.Group("/food")
	food.Use(middleware.AuthMiddleware(appctx))
	food.POST("", foodhandler.CreateFoodHandler(appctx))
	food.PUT("/:id", foodhandler.UpdateFoodHandler(appctx))
	food.PUT("", foodhandler.UpdateFoodNonePhoto(appctx))

	food.DELETE("/:id", foodhandler.DeleteFoodHandler(appctx))
	food.GET("", foodhandler.GetListFoodHandler(appctx))
	food.GET("/:id", foodhandler.GetFoodHandler(appctx))
	food.DELETE("photo/:id", foodhandler.DeletePhotoHandler(appctx))
	food.POST("photo", foodhandler.CreatePhotoHandler(appctx))
	food.POST("photos", foodhandler.CreatePhotoListHandler(appctx))
	//exersice
	exersice := v1.Group("/exersice")
	exersice.Use(middleware.AuthMiddleware(appctx))
	exersice.POST("", exersicehandler.CreateExersiceHandler(appctx))
	exersice.PUT("/:id", exersicehandler.UpdateExersiceHandeler(appctx))
	exersice.DELETE("/:id", exersicehandler.DeleteExersiceHandler(appctx))
	exersice.GET("", exersicehandler.GetistExersiceHandler(appctx))
	exersice.GET("/:id", exersicehandler.GetExersiceByIdHandler(appctx))
	exersice.PUT("", exersicehandler.UpdateExersiceNonePhotoById(appctx))
	exersice.GET("/type", exersicehandler.GetAllExTypeHandler(appctx))
	//user
	user := v1.Group("/user")
	user.Use(middleware.AuthMiddleware(appctx))

	user.GET("", userhandler.GetAllUserHandler(appctx))
	user.GET("/:id", userhandler.GetUserByIdHandler(appctx))
	user.POST("", userhandler.CreateUserHandler(appctx))
	user.PUT("/:id", userhandler.UpdateUserHandler(appctx))
	user.DELETE("/:id", userhandler.DeleteUserHandler(appctx))
	user.PUT("/password", accounthandlers.ChangePasswordHandler(appctx))
	//meal
	meal := v1.Group("/meal")
	meal.Use(middleware.AuthMiddleware(appctx))
	meal.POST("", mealhandler.CreateMealHandler(appctx))
	meal.GET("/user/:id", mealhandler.GetMealsByUserIdHandler(appctx))
	meal.GET("/:id", mealhandler.GetMealByIdHandler(appctx))
	meal.POST("/detail", mealhandler.CreateMealDetailHandler(appctx))
	meal.PUT("/detail/:id", mealhandler.UpdateMealDetail(appctx))
	meal.DELETE("/:id", mealhandler.DeleteMealById(appctx))
	meal.DELETE("/detail/:id", mealhandler.DeleteDetailMealById(appctx))
	meal.GET("/user/:id/date/:date", mealhandler.GetMealByDateHandler(appctx))
	meal.GET("/calorie/user/:id/date/:date", mealhandler.GetInfomationCaloriesHandler(appctx))

	//content
	content := v1.Group("/content")
	content.Use(middleware.AuthMiddleware(appctx))
	content.POST("", personalcontenthandler.CreatePostHandler(appctx))
	content.POST("/like", personalcontenthandler.LikeHandler(appctx))
	content.POST("/comment", personalcontenthandler.CreateCommentwithimageHandler(appctx))
	content.POST("/commentnonephoto", personalcontenthandler.CreateCommentNoneHandler(appctx))
	content.DELETE("/like", personalcontenthandler.DeleteLikeByUserIdAndPostIdHandler(appctx))
	content.DELETE("/:id", personalcontenthandler.DeletePostHandler(appctx))
	content.PUT("", personalcontenthandler.UpdatePostHandler(appctx))
	content.GET("/:id", personalcontenthandler.GetPostByIdHandler(appctx))
	content.GET("/user/:id", personalcontenthandler.GetAllPostByIdHandler(appctx))
	content.GET("", personalcontenthandler.GetAllPostHandler(appctx))
	content.GET("/coment/:id", personalcontenthandler.GetAllComentByPostIdHandler(appctx))
	content.GET("/islike/user/:id/post/:postid", personalcontenthandler.CheckIsLikeByUserIdAndPosstIdHandler(appctx))
	content.GET("/like/:id", personalcontenthandler.GetAllLikesByUserIddHandler(appctx))

	// schedule
	schedule := v1.Group("/schedule")
	schedule.Use(middleware.AuthMiddleware(appctx))
	schedule.POST("", schedulehandler.CreateScheduleHandler(appctx))
	schedule.POST("/detail", schedulehandler.CreateScheduleDetailHandler(appctx))
	schedule.GET("", schedulehandler.GetAllScheduleHandler(appctx))
	schedule.GET("/:id", schedulehandler.GetScheduleByIdHandler(appctx))
	schedule.GET("/user/:id", schedulehandler.GetScheduleByUserIdHandler(appctx))
	schedule.PUT("", schedulehandler.UpdateScheduleHandler(appctx))
	schedule.PUT("/detail", schedulehandler.UpdateScheduleDetailHandler(appctx))
	schedule.DELETE("/:id", schedulehandler.DeleteScheduleHandler(appctx))
	schedule.DELETE("/detail/:id", schedulehandler.DeleteScheduleDetailHandler(appctx))
	schedule.GET("/user/:id/date/:date", schedulehandler.GetScheduleByDateHandler(appctx))
	schedule.GET("/user/:id/fromdate/:fromdate/date/:date", schedulehandler.GetScheduleDateToDateHandler(appctx))

	// message
	message := v1.Group("/conversation")
	message.Use(middleware.AuthMiddleware(appctx))
	message.POST("", messagehandler.CreateConversationHandler(appctx))
	message.POST("/messages", messagehandler.SendMessageHandler(appctx))
	message.GET("/users/:user_id/conversations", messagehandler.ListUserConversationsHandler(appctx))
	message.GET("/:conversation_id/messages", messagehandler.ListConversationMessagesHandler(appctx))
	// delete

	// reminder
	reminder := v1.Group("/reminder")
	reminder.Use(middleware.AuthMiddleware(appctx))
	reminder.POST("", reminderhandler.CreateReminderHandler(appctx))
	reminder.PUT("", reminderhandler.UpdateReminderHandler(appctx))
	reminder.GET("/:id", reminderhandler.GetReminderByIdHandler(appctx))
	reminder.DELETE("/:id", reminderhandler.DeleteReminderByIdHandler(appctx))
	reminder.GET("/user/:id", reminderhandler.GetReminderByIdUserHandler(appctx))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
