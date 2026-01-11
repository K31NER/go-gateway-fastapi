package schemas

var id int = 0

type Users struct {
	Id   int
	Name string `json:"name"`
	Mail string `json:"mail"`
	Age  int    `json:"age"`
}

// Metodo para auto incrementar el id
func (u *Users) AddId() {
	id++
	u.Id = id
}