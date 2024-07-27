package dto

type CommonFilter struct {
	Limit    int
	Offset   *int
	Select   []string
	Sort     string
	Preloads []string
}
