package main

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrUserExists = errors.New("user with this email already exists")

// UserService provides database operations for users
type UserService struct{}

// FindUserByEmail finds a user by email
func (us *UserService) FindUserByEmail(email string) (*User, error) {
	email = strings.ToLower(email)
	
	if useInMemory {
		usersMutex.RLock()
		defer usersMutex.RUnlock()
		
		for _, user := range users {
			if strings.ToLower(user.Email) == email {
				return &user, nil
			}
		}
		return nil, mongo.ErrNoDocuments
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByID finds a user by ID
func (us *UserService) FindUserByID(userID primitive.ObjectID) (*User, error) {
	if useInMemory {
		usersMutex.RLock()
		defer usersMutex.RUnlock()
		
		userIDStr := userID.Hex()
		if user, exists := users[userIDStr]; exists {
			return &user, nil
		}
		return nil, mongo.ErrNoDocuments
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	err := userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (us *UserService) CreateUser(user *User) error {
	if useInMemory {
		usersMutex.Lock()
		defer usersMutex.Unlock()
		
		// Check if email already exists
		for _, existingUser := range users {
			if strings.ToLower(existingUser.Email) == strings.ToLower(user.Email) {
				return ErrUserExists
			}
		}
		
		// Generate new ObjectID if not set
		if user.ID.IsZero() {
			user.ID = primitive.NewObjectID()
		}
		
		users[user.ID.Hex()] = *user
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, user)
	return err
}

// Global user service instance
var userService = &UserService{}