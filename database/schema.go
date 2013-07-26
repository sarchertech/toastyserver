package database

//TODO restrict integer sizes

func schema() map[string]string {
	s := make(map[string]string)

	//insert Null into id to auto increment
	s["Customer"] = `(id integer primary key,
	 		 		  name text not null,
	 		 		  phone text not null,
			 		  status boolean not null,
			 		  level integer not null,
			 		  fob_num integer not null unique)`

	s["Employee"] = `(id integer primary key,
	 		 		  name text not null unique,
			 		  level integer not null,
			 		  fob_num integer not null unique)`

	s["Keyfob"] = `(fob_num integer not null primary key,
					admin boolean not null)`

	return s
}
