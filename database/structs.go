package database

//check documentation for optional StructTag metadata
//TODO add documentation on StructTag metadata

type Customer struct {
	Id      int `db:"autoInc"`
	Name    string
	Phone   string
	Status  bool
	Level   int
	Fob_num int
}

type Employee struct {
	Id      int `db:"autoInc"`
	Name    string
	Level   int
	Fob_num int
}

type Keyfob struct {
	Fob_num int
	Admin   bool
}

type Bed struct {
	Bed_num  int
	Level    int
	Max_time int
	Name     string
	Status   bool `db:"false"` //not DB backed
}

type Session struct {
	Id          int `db:"autoInc"`
	Bed_num     int
	Customer_id int
	Time_stamp  int64
}
