package dto

import "fx-golang-server/module/core/model"

type ListMoviesRequest struct {
	Paginate
}

type ListMoviesResponse struct {
	Movies []model.MovieModel `json:"movies"`
}

type FilterMovie struct {
	CommonFilter CommonFilter
	ID           string
}

type AdminCreateMovieRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type AdminCreateMovieResponse struct {
	MovieID string `json:"movie_id"`
}
