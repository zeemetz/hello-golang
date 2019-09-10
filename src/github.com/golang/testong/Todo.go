package main

// Todo is a entity of what todo
type Todo struct {
	ID        int
	Name      string
	Completed bool
	Due       string
}

//Todos is
type Todos []Todo

func init() {
	db().AutoMigrate(&Todo{})
}
