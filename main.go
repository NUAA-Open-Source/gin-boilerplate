package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"a2os/gin-boilerplate/common"
	"a2os/gin-boilerplate/controller/misc"
	_ "a2os/gin-boilerplate/docs"
	"a2os/gin-boilerplate/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	csrf "github.com/utrack/gin-csrf"
)

// @title A2OS example
// @version 0.0.1
// @description A2OS example API Documentation.

// @contact.name A2OS Dev Team
// @contact.url https://groups.google.com/group/a2os-general
// @contact.email a2os-general@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.example.a2os.club

func migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin auto_increment=1").AutoMigrate(&model.Example{})
}

func init() {
	// init config
	common.DefaultConfig()
	common.SetConfig()
	common.WatchConfig()

	// init sentry error tracking service
	common.InitSentry()

	// init logger
	common.InitLogger()

	// init Database
	db := common.InitMySQL()
	migrate(db)
}

func main() {

	// Before init router
	if viper.GetBool("basic.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		// Redirect log to file
		gin.DisableConsoleColor()
		logFile := common.GetLogFile()
		defer logFile.Close()
		gin.DefaultWriter = io.MultiWriter(logFile)
	}

	r := gin.Default()
	// Error handling
	r.Use(common.ErrorHandling())
	r.Use(common.MaintenanceHandling())
	// After init router
	// CORS

	if viper.GetBool("basic.debug") {
		r.Use(cors.New(cors.Config{
			// The value of the 'Access-Control-Allow-Origin' header in the
			// response must not be the wildcard '*' when the request's
			// credentials mode is 'include'.
			AllowOrigins:     common.CORS_ALLOW_DEBUG_ORIGINS,
			AllowMethods:     common.CORS_ALLOW_METHODS,
			AllowHeaders:     common.CORS_ALLOW_HEADERS,
			ExposeHeaders:    common.CORS_EXPOSE_HEADERS,
			AllowCredentials: true,
			AllowWildcard:    true,
			MaxAge:           12 * time.Hour,
		}))
		//r.Use(CORS())
	} else {
		// RELEASE Mode
		r.Use(cors.New(cors.Config{
			AllowOrigins:     common.CORS_ALLOW_ORIGINS,
			AllowMethods:     common.CORS_ALLOW_METHODS,
			AllowHeaders:     common.CORS_ALLOW_HEADERS,
			ExposeHeaders:    common.CORS_EXPOSE_HEADERS,
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	// CSRF
	store := cookie.NewStore(common.CSRF_COOKIE_SECRET)
	r.Use(sessions.Sessions(common.CSRF_SESSION_NAME, store))
	CSRF := csrf.Middleware(csrf.Options{
		Secret: common.CSRF_SECRET,
		ErrorFunc: func(c *gin.Context) {
			//c.String(http.StatusBadRequest, "CSRF token mismatch")
			c.JSON(http.StatusBadRequest, gin.H{
				"err_code": 10007,
				"message":  common.Errors[10007],
			})
			log.Println(c.ClientIP(), "CSRF token mismatch")
			c.Abort()
		},
	})

	// swagger router
	if viper.GetBool("basic.debug") {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// misc operations
	r.GET("/ping", misc.Ping)
	r.GET("/csrf", CSRF, misc.Csrf)

	// the API with CSRF middleware
	v1Csrf := r.Group("/v1", CSRF)
	{
		// Write your service operations here...
		v1Csrf.GET("example")
	}

	r.Run(":" + viper.GetString("basic.port")) // listen and serve on 0.0.0.0:PORT
}
