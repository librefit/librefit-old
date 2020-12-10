package api

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/librefitness/librefitness/internal/config"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "PATCH", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", ""},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

type HTTPServer struct {
	*gin.Engine
	config *config.Config
}

func NewHTTPServer(cfg *config.Config) (*HTTPServer, error) {
	server := &HTTPServer{
		config: cfg,
	}
	err := server.init()
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (s *HTTPServer) Run() {
	s.Engine.Run(":4000")
}

func (s *HTTPServer) init() error {
	r := gin.New()

	r.Use(CORS())

	r.Use(static.Serve("/", static.LocalFile("./static", true)))
	r.Use(static.Serve("/img", static.LocalFile("./data/img", true)))

	// Middleware for JWT
	authMiddleware, err := jwt.New(newGinJWTMiddleware())
	if err != nil {
		log.Fatal("JWT Error: ", err.Error())
		return err
	}

	err = authMiddleware.MiddlewareInit()
	if err != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:", err.Error())
		return err
	}

	r.Use(gin.Logger(), gin.Recovery())

	// When we don't match any route we go here.
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found."})
	})

	// Because of fileUpload:
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// API Routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user/login", authMiddleware.LoginHandler)
		v1.POST("/user/logout", authMiddleware.LogoutHandler)
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
		v1.GET("/healthcheck", healthCheck)

		// We don't need middleware for OpenFoodFacts search
		v1.GET("/off/search", offSearch)
		v1.GET("/off/product/:code", offProduct)

		// Uploads
		v1.POST("/upload", fileUpload)
	}

	v1.Use(authMiddleware.MiddlewareFunc())
	{
		// User
		v1.GET("/user/info", userInfo)
		v1.GET("/user/preferences", userPreferences)
		v1.PUT("/user/preferences", userPreferencesUpdate)

		// Measures (a.k.a weight)
		v1.GET("/measures", measures)
		v1.POST("/measures", measureCreate)
		v1.DELETE("/measures/:id", measureDelete)
		v1.PUT("/measures/:id", measureUpdate)

		// Fluids (a.k.a water)
		v1.GET("/fluids", fluids)
		v1.POST("/fluids", fluidsWithRange)
		v1.POST("/fluid", fluidCreate)
		v1.DELETE("/fluids/:id", fluidDelete)
		v1.PUT("/fluids/:id", fluidUpdate)

		// Stats
		v1.GET("/stats/fluids", statsFluids)
		v1.GET("/stats/food_diary", statsFoodDiary)

		// Food Inventory
		v1.GET("/food/inventory", foodInventory)
		v1.POST("/food/inventory", foodInventoryCreate)
		v1.DELETE("/food/inventory/:id", foodInventoryDelete)
		v1.PUT("/food/inventory/:id", foodInventoryUpdate)

		// Food Diary
		v1.GET("/food/diary", foodDiary)
		v1.GET("/food/diary/:id", foodDiaryID)
		v1.POST("/food/diary", foodDiaryCreate)
		v1.DELETE("/food/diary/:id", foodDiaryDelete)
		v1.PUT("/food/diary/:id", foodDiaryUpdate)
	}

	if s.config.Opts.Debug {
		r.Use(gin.Recovery())
	}

	s.Engine = r

	return nil
}
