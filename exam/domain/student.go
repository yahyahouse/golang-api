package domain

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Grade  int    `json:"grade"`
	Status string `json:"status"`
}

func (Student) TableName() string {
	return "student"
}
