package test

type GetTestByIdIN struct {
	Id int `json:"id" binding:"required"`
}

type DeleteTestIN struct {
	Id int `json:"id" binding:"required"`
}

// sql struct
type TestResult struct {
	A string
	B string
	C string
}
