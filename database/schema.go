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

	s["Bed"] = `(Bed_num integer not null primary key,
				 Level integer not null,
				 Max_time integer not null,
				 Name text not null)`

	s["Session"] = `(Id integer primary key,
					 Bed_num integer not null,
					 Customer_id integer not null,
					 Session_time integer not null,
					 Time_stamp integer not null)`

	s["DoorAccess"] = `(Id integer primary key,
						Customer_id integer not null,
						Time_stamp integer not null)`

	return s
}
