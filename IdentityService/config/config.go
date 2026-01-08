package config

import (
	"fmt"
	"os"
	"time"
)

// Config holds all application configuration
type Config struct {
	Database     DatabaseConfig
	Server       ServerConfig
	Timeouts     TimeoutConfig
	CORS         CORSConfig
	Validation   ValidationConfig
}

// DatabaseConfig holds database-specific configuration
type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Address         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

// TimeoutConfig holds various timeout configurations
type TimeoutConfig struct {
	DatabaseConnect time.Duration
	Handler         time.Duration
	SchemaInit      time.Duration
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	ExposedHeaders   []string
	AllowCredentials bool
	MaxAge           int
}

// ValidationConfig holds validation rules
type ValidationConfig struct {
	MaxNameLength  int
	MaxEmailLength int
	EmailRegex     string
}

// LoadConfig loads all application configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		Database: loadDatabaseConfig(),
		Server:   loadServerConfig(),
		Timeouts: loadTimeoutConfig(),
		CORS:     loadCORSConfig(),
		Validation: loadValidationConfig(),
	}
}

func loadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:            getEnv("DB_HOST", "postgres"),
		Port:            getEnv("DB_PORT", "5432"),
		User:            getEnv("DB_USER", "identity_user"),
		Password:        getEnv("DB_PASSWORD", "postgres"),
		DBName:          getEnv("DB_NAME", "identity_db"),
		SSLMode:         getEnv("DB_SSLMODE", "disable"),
		MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 5),
		ConnMaxLifetime: getEnvDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		ConnMaxIdleTime: getEnvDuration("DB_CONN_MAX_IDLETIME", 1*time.Minute),
	}
}

func (c *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

func loadServerConfig() ServerConfig {
	return ServerConfig{
		Address:         getEnv("SERVER_ADDRESS", "0.0.0.0:8080"),
		ReadTimeout:     getEnvDuration("SERVER_READ_TIMEOUT", 15*time.Second),
		WriteTimeout:    getEnvDuration("SERVER_WRITE_TIMEOUT", 15*time.Second),
		ShutdownTimeout: getEnvDuration("SERVER_SHUTDOWN_TIMEOUT", 10*time.Second),
	}
}

func loadTimeoutConfig() TimeoutConfig {
	return TimeoutConfig{
		DatabaseConnect: getEnvDuration("TIMEOUT_DB_CONNECT", 10*time.Second),
		Handler:         getEnvDuration("TIMEOUT_HANDLER", 5*time.Second),
		SchemaInit:      getEnvDuration("TIMEOUT_SCHEMA_INIT", 30*time.Second),
	}
}

func loadCORSConfig() CORSConfig {
	return CORSConfig{
		AllowedOrigins:   getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"*"}),
		AllowedMethods:   getEnvSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		AllowedHeaders:   getEnvSlice("CORS_ALLOWED_HEADERS", []string{"Content-Type", "Authorization"}),
		ExposedHeaders:   getEnvSlice("CORS_EXPOSED_HEADERS", []string{}),
		AllowCredentials: getEnvBool("CORS_ALLOW_CREDENTIALS", false),
		MaxAge:           getEnvInt("CORS_MAX_AGE", 86400), // 24 hours
	}
}

func loadValidationConfig() ValidationConfig {
	return ValidationConfig{
		MaxNameLength:  getEnvInt("VALIDATION_MAX_NAME_LENGTH", 100),
		MaxEmailLength: getEnvInt("VALIDATION_MAX_EMAIL_LENGTH", 100),
		EmailRegex:     getEnv("VALIDATION_EMAIL_REGEX", `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var intVal int
		if _, err := fmt.Sscanf(value, "%d", &intVal); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true" || value == "1"
	}
	return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return []string{value}
	}
	return defaultValue
}
