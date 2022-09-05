package handler

type User struct {
	Name string `json:"name" binding:"NotNullAndAdmin"`
	Age  int    `json:"age"`
}
