package main

import (
	"context"
	"identity-service/config"
	"identity-service/database"
	"identity-service/handlers"
	"identity-service/middleware"
	"identity-service/repository"
	"identity-service/service"
	"identity-service/validation"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Load configuration
	cfg := config.LoadConfig()

	logger.Info("starting application",
		slog.String("database_host", cfg.Database.Host),
		slog.String("database_name", cfg.Database.DBName),
		slog.String("server_address", cfg.Server.Address),
	)

	// Connect to database
	db, err := database.NewDatabase(&cfg.Database)
	if err != nil {
		logger.Error("failed to connect to database",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("successfully connected to PostgreSQL")

	// Initialize database schema
	schemaCtx, schemaCancel := context.WithTimeout(context.Background(), cfg.Timeouts.SchemaInit)
	defer schemaCancel()

	if err := db.InitSchema(schemaCtx); err != nil {
		logger.Error("failed to initialize schema",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db.DB)
	validator := validation.NewValidator(&cfg.Validation)
	userService := service.NewUserService(userRepo, validator)
	userHandler := handlers.NewUserHandler(userService, cfg, logger)

	// Setup middleware
	corsMiddleware := middleware.CORS(&middleware.CORSConfig{
		AllowedOrigins:   cfg.CORS.AllowedOrigins,
		AllowedMethods:   cfg.CORS.AllowedMethods,
		AllowedHeaders:   cfg.CORS.AllowedHeaders,
		ExposedHeaders:   cfg.CORS.ExposedHeaders,
		AllowCredentials: cfg.CORS.AllowCredentials,
		MaxAge:           cfg.CORS.MaxAge,
	})

	// Setup router with middleware
	mux := http.NewServeMux()

	// Apply CORS middleware to all routes
	mux.Handle("/api/users", corsMiddleware(http.HandlerFunc(userHandler.GetAllUsers)))
	mux.Handle("/api/users/create", corsMiddleware(http.HandlerFunc(userHandler.CreateUser)))

	// Create HTTP server with proper configuration
	server := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      mux,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Graceful shutdown handling
	go func() {
		logger.Info("server starting",
			slog.String("address", cfg.Server.Address),
		)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server failed to start",
				slog.String("error", err.Error()),
			)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("server shutting down")

	// Graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("server forced to shutdown",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	logger.Info("server exited properly")
}
