package main

import (
	"net/http"
	"os"

	"log/slog"
	"planetbi/config"
	"planetbi/db"
	"planetbi/handlers"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitlab.com/greyxor/slogor"
)

func main() {
	println("Starting...")

	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, slogor.Options{
		ShowSource: true,
	})))

	db.Init()

	router := mux.NewRouter()
	handlers.Handle(
		router,
		handlers.Users,
	)

	headers := gorillaHandlers.AllowedHeaders([]string{"*", "Skip-Recording-Operations", "Content-Type", "Authorization"})

	// TODO: Check if we can/should drop this in production and replace it with "*"
	origins := gorillaHandlers.AllowedOrigins([]string{"http://localhost:5173"})

	methods := gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"})

	http.Handle("/api/", gorillaHandlers.RecoveryHandler()(gorillaHandlers.CORS(headers, methods, origins)(router)))
	http.Handle("/", gorillaHandlers.RecoveryHandler()(gorillaHandlers.CORS(headers, methods, origins)(router)))

	port := os.Getenv("PORT")
	if port == "" {
		port = config.ServerPort
	}

	println("Server started on port " + port + " 🚀")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
