CREATE TABLE users
(
    id PRIMARY KEY,
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
)

CREATE TABLE todo_lists
(
    id PRIMARY KEY,
    user_id FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE,


)