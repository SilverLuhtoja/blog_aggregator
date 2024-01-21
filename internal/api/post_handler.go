package api

import (
	"fmt"
	"net/http"

	"github.com/SilverLuhtoja/blog_aggregator/internal/database"
	"github.com/SilverLuhtoja/blog_aggregator/internal/models"
)

type Limit struct {
	Limit int32 `json:"limit"`
}

func (cfg *ApiConfig) GetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	params, err := GetParamsFromRequestBody(Limit{}, r)
	if err != nil {
		fmt.Println("NO parameter")
		params.Limit = 10
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), params.Limit)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	convertedPosts := []models.Post{}
	for _, post := range posts {
		convertedPosts = append(convertedPosts, models.DatabasePostToPost(post))
	}
	RespondWithJSON(w, 200, convertedPosts)
}
