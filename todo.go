package todo

type TodoList struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Desription string `json:"desription"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Desription string `json:"desription"`
	Done       bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ItemId int
	ListId int
}
