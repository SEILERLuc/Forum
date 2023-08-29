CREATE TABLE forum_data_test (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		User_id INTEGER,
		Username TEXT,
		User_post TEXT,
		Type  []TEXT,
		FOREIGN KEY (User_id) REFERENCES forum_login (ID),
		FOREIGN KEY (Username) REFERENCES forum_login (Username)
);