package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/internal/conn"
	"github.com/spf13/cobra"

	_authHttp "github.com/mokhlesur-rahman/golang-basic-crud-api-server/user/delivery/http"
	_authRepository "github.com/mokhlesur-rahman/golang-basic-crud-api-server/user/repository"
	_authUseCase "github.com/mokhlesur-rahman/golang-basic-crud-api-server/user/usecase"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starting Server...",
	Long:  `Starting Server...`,
	Run:   server,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func server(cmd *cobra.Command, args []string) {
	log.Println("Connecting database")
	if err := conn.ConnectDB(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connected successfully!")

	// boot http server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	srv := buildHTTP(cmd, args)
	go func(sr *http.Server) {
		if err := sr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}(srv)
	<-stop
}

func buildHTTP(_ *cobra.Command, _ []string) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Println(err)
		}

	})

	db := conn.DefaultDB()
	_ = db

	authRepo := _authRepository.New(db)
	authUsecase := _authUseCase.New(authRepo)
	_authHttp.NewHTTPHandler(r, authUsecase)

	return &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8081),
		Handler: r,
	}
}
