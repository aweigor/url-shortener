package main

import (
	"fmt"
	"net/http"
	"url-shortener/configs"
	"url-shortener/internal/auth"
	"url-shortener/internal/heartbeat"
	"url-shortener/internal/link"
	"url-shortener/internal/stat"
	"url-shortener/internal/user"
	"url-shortener/pkg/db"
	"url-shortener/pkg/event"
	"url-shortener/pkg/middleware"
)

func NewApp() http.Handler {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)

	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)
	statRepository := stat.NewStatRepository(database)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})
	// Handlers
	heartbeat.NewHeartbeatHandler(router)
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	go statService.AddClick()

	// Middlewares
	mwStack := middleware.Chain(middleware.CORS, middleware.Logging)

	return mwStack(router)
}

func main() {

	app := NewApp()

	server := http.Server{Addr: ":8081", Handler: app}

	fmt.Println("Server is listening on port 8081")

	server.ListenAndServe()
}
