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
		// AllowOrigins:     []string{"http://localhost:3000"},
		// AllowOrigins: []string{"http://localhost:3000", "https://uk.openfoodfacts.org"},
		AllowOrigins:     []string{"*"},
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

	// API Routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user/login", authMiddleware.LoginHandler)
		v1.POST("/user/logout", authMiddleware.LogoutHandler)
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
		v1.GET("/healthcheck", healthCheck)

		// We don't need middleware for OpenFoodFacts search
		v1.GET("/off/search", offSearch)
	}

	v1.Use(authMiddleware.MiddlewareFunc())
	{
		// User
		v1.GET("/user/info", userInfo)
		v1.GET("/user/preferences", userPreferences)
		v1.PUT("/user/preferences", userPreferencesUpdate)

		// Measures (a.k.a weight)
		v1.GET("/measures", measures)
		v1.POST("/measure", measureCreate)
		v1.DELETE("/measure/:id", measureDelete)
		v1.PUT("/measure/:id", measureUpdate)

		// Fluids (a.k.a water)
		v1.GET("/fluids", fluids)
		v1.POST("/fluids", fluidsWithRange)
		v1.POST("/fluid", fluidCreate)
		v1.DELETE("/fluids/:id", fluidDelete)
		v1.PUT("/fluids/:id", fluidUpdate)

		// Food Inventory
		v1.GET("/food/inventories", foodInventory)
		v1.POST("/food/inventories", fluidsWithRange)
		v1.DELETE("/food/inventories/:id", fluidDelete)
		v1.PUT("/food/inventories/:id", fluidUpdate)
	}

	if s.config.Opts.Debug {
		r.Use(gin.Recovery())
	}

	s.Engine = r

	return nil
}
