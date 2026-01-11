package schemas

var id int = 0

type Users struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
	Age  int    `json:"age"`
}

func (u *Users) AddId() {
	id++
	u.Id = id
}