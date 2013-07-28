package database

//TODO restrict integer sizes

func schema() map[string]string {
	s := make(map[string]string)

	//insert Null into id to auto increment
	s["Customer"] = `(Id integer primary key,
	 		 		  Name text not null,
	 		 		  Phone text not null,
			 		  Status boolean not null,
			 		  Level integer not null,
			 		  Fob_num integer not null unique)`

	s["Employee"] = `(Id integer primary key,
	 		 		  Name text not null unique,
			 		  Level integer not null,
			 		  Fob_num integer not null unique)`

	s["Keyfob"] = `(Fob_num integer not null primary key,
					Admin boolean not null)`

	return s
}
