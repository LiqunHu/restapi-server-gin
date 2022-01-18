package test

type GetTestByIdIN struct {
	Id int `json:"id" binding:"required"`
}

type UpdateTestIN struct {
	Id int `json:"id" binding:"required"`
	A  string
	B  string
	C  string
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
