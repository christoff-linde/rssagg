package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/christoff-linde/rssagg/internal/auth"
	"github.com/christoff-linde/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	newUser, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(newUser))
}

func (apiConfig *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
	}
	user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Could not get user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
