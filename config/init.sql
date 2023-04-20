create table Users (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username VARCHAR ( 50 ) UNIQUE NOT NULL, 
    password VARCHAR ( 50 ) NOT NULL
);

create table Tokens (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    token VARCHAR ( 50 ) UNIQUE NOT NULL,
    user_id INT,
    FOREIGN KEY (user_id)
      REFERENCES Users (id)
);

create table Tags (
    id serial PRIMARY KEY,
    tag VARCHAR ( 50 )
);

create table Bookmarks (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id int,
    FOREIGN KEY (user_id)
        REFERENCES Users (id),
    bookmark TEXT,
    tags int[]
);