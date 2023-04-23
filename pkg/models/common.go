package models

// Response - response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Student - student struct
type StudentClass1 struct {
	StudentID   int    `json:"id"`
	StudentName string `json:"name"`
	StudentAge  int    `json:"age"`
}
