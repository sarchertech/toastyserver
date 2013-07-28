package database

type Customer struct {
	Id      int
	Name    string
	Phone   string
	Status  bool
	Level   int
	Fob_num int
}

type Employee struct {
	Id      int
	Name    string
	Level   int
	Fob_num int
}

type Keyfob struct {
	Fob_num int
	Admin   bool
}
