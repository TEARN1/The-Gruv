package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// MongoDB connection
	mongoClient *mongo.Client
	userCollection *mongo.Collection
	
	// In-memory fallback storage
	users      = make(map[string]User)
	usersMutex = &sync.RWMutex{}
	useInMemory = false
	
	// JWT secret key
	jwtSecret = []byte("your-secret-key") // In production, use environment variable
)

// InitDatabase initializes the MongoDB connection or falls back to in-memory storage
func InitDatabase() error {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Println("MONGODB_URI not set, using in-memory storage for development")
		useInMemory = true
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Printf("Failed to connect to MongoDB, falling back to in-memory storage: %v", err)
		useInMemory = true
		return nil
	}

	// Test the connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Printf("Failed to ping MongoDB, falling back to in-memory storage: %v", err)
		useInMemory = true
		return nil
	}

	mongoClient = client
	userCollection = client.Database("thegruv").Collection("users")
	
	log.Println("Connected to MongoDB successfully")
	return nil
}

// CloseDatabase closes the MongoDB connection
func CloseDatabase() {
	if mongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}
}