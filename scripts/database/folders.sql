CREATE TABLE folders (
  id SERIAL,
  parent_id INT,
  name VARCHAR(60) NOT NULL,
  created_at TIMESTAMP DEFAULT current_timestamp,
  modifed_at TIMESTAMP NOT NULL,
  deleted BOOL NOT NULL DEFAULT false,
  PRIMARY KEY(id),
  CONSTRAINT fk_parent
    FOREIGN KEY(parent_id)
    REFERENCES folders(id)
)