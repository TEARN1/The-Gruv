package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtSecret = []byte("dev-secret-change-me")

func initDB(path string) error {
	var err error
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	create := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(create)
	return err
}

func hashPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(b), err
}

func checkPassword(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

func createUser(id, username, password string) error {
	h, err := hashPassword(password)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", id, username, h)
	return err
}

func findUserByUsername(username string) (id, hashed string, err error) {
	row := db.QueryRow("SELECT id, password FROM users WHERE username = ?", username)
	err = row.Scan(&id, &hashed)
	return
}

func issueJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func main() {
	dbPath := os.Getenv("USER_DB_PATH")
	if dbPath == "" {
		dbPath = "data/dev.db"
	}
	if err := initDB(dbPath); err != nil {
		log.Fatalf("failed init db: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "User Service", "status": "UP"})
	})

	r.POST("/api/users/register", func(c *gin.Context) {
		var payload struct {
			UserID   string `json:"userId" binding:"required"`
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := createUser(payload.UserID, payload.Username, payload.Password); err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "username taken or db error: " + err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "userId": payload.UserID})
	})

	r.POST("/api/users/login", func(c *gin.Context) {
		var payload struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, hashed, err := findUserByUsername(payload.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		if err := checkPassword(hashed, payload.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		token, err := issueJWT(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to issue token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "login successful", "userId": id, "token": token})
	})

	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Printf("User Service listening on port %s (db: %s)\n", port, dbPath)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}