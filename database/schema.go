package database

//TODO restrict integer sizes

func schema() map[string]string {
	s := make(map[string]string)

	//integer primary key auto increments
	//insert Null into id to auto increment
	s["Customer"] = `(id integer primary key,
	 		 		  name text not null,
	 		 		  phone text not null,
			 		  status boolean not null,
			 		  level integer not null)`

	//TODO need to enforce Uniqueness on name
	s["Employee"] = `(id integer primary key,
	 		 		  name text not null,
			 		  level integer not null)`

	s["Keyfob"] = `(fob_num integer not null primary key,
	 		 		customer_id integer)` //foreign key, but not enforced by sqlite

	return s
}
