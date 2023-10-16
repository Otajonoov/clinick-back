package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "gitlab.com/QuvonchbekOtajonov/clinic-back/api/docs"
	v1 "gitlab.com/QuvonchbekOtajonov/clinic-back/api/handlers/v1"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/config"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/pkg/logger"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/storage"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RoutetOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
	Log     logger.Logger
}

// @Description Created by Otajonov Quvonchbek
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option RoutetOptions) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowHeaders = append(corConfig.AllowHeaders, "*")
	router.Use(cors.New(corConfig))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "App is running...",
		})
	})

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:     option.Log,
		Cfg:     option.Cfg,
		Storage: &option.Storage,
	})

	router.Static("/media", "./media")
	api := router.Group("/v1")

	// File-Upload
	api.POST("file-upload", handlerV1.UploadFile)

	// Doctor...
	api.POST("/create-doctor",handlerV1.AuthMiddleWare ,handlerV1.HandleCreateDoctor)
	api.GET("/get-doctors",handlerV1.HandleGetDoctors)
	api.DELETE("/delete-doctor/:id",handlerV1.AuthMiddleWare ,handlerV1.HandleDeleteDoctor)
	api.POST("/update-doctor/:id",handlerV1.AuthMiddleWare ,handlerV1.HandleUpdateDoctor)
	api.GET("doctorById-get/:id", handlerV1.DoctorGet)

	// Service
	api.POST("/create-service",handlerV1.AuthMiddleWare ,handlerV1.CreateService)
	api.GET("get-services",handlerV1.GetAllServices)
	api.DELETE("delete-service/:id",handlerV1.AuthMiddleWare ,handlerV1.DeleteService)
	api.POST("/update-service/:id",handlerV1.AuthMiddleWare ,handlerV1.UpdateService)
	api.GET("serviceById-get/:id", handlerV1.ServiceGet)


	// Customer
	api.POST("/create-customer",handlerV1.AuthMiddleWare ,handlerV1.CreateCustomer)
	api.GET("/get-customers",handlerV1.GetAllCustomers)
	api.DELETE("/delete-customer/:id",handlerV1.AuthMiddleWare ,handlerV1.DeleteCustomer)
	api.POST("/update-customer/:id",handlerV1.AuthMiddleWare ,handlerV1.UpdateCustomer)
	api.GET("customerById-get/:id", handlerV1.CustomerGet)


	// User 
	api.POST("/create-user", handlerV1.CreateNewUser)
	api.POST("/login", handlerV1.Login)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
