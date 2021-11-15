package api

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}
