CREATE TABLE users (
  id SERIAL,
  name VARCHAR(80) NOT NULL,
  login VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(200) NOT NULL,
  created_at TIMESTAMP DEFAULT current_timestamp,
  modifed_at TIMESTAMP NOT NULL,
  deleted BOOL NOT NULL DEFAULT false,
  last_login TIMESTAMP DEFAULT current_timestamp,
  PRIMARY KEY(id)
)