package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SilverLuhtoja/blog_aggregator/internal/auth"
	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/SilverLuhtoja/blog_aggregator/internal/models"
	"github.com/google/uuid"
)

type User struct {
	Name string `json:"name"`
}

func (cfg *ApiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req, err := GetParamsFromRequestBody(User{}, r)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("createUserHandler - ", err))
		return
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      req.Name,
	}

	user, err := cfg.DB.CreateUser(r.Context(), userParams)
	if err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "createUserHandler - couldn't create user to database")
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DatabaseUserToUser(user))
}

func (cfg *ApiConfig) GetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	key, err := auth.GetApiKey(r.Header)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := cfg.DB.GetUserByApiKey(r.Context(), key)
	if err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "getUserByApiKey - couldn't find user")
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DatabaseUserToUser(user))
}
