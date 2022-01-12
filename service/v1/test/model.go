package test

type GetTestByIdIN struct {
	Id int `json:"id" binding:"required"`
}
