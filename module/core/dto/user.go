package dto

type FilterUser struct {
	Phone        string
	Email        string
	ID           string
	CommonFilter CommonFilter
}

type SearchUsersRequest struct {
	Search string `json:"search"`
}

type SearchUsersResponse struct {
	
}