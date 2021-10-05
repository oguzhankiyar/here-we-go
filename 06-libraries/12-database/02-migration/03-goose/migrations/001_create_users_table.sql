-- +goose Up
CREATE TABLE users (
   Id int NOT NULL PRIMARY KEY,
   Username text,
   FirstName text,
   LastName text
);

INSERT INTO users VALUES
  (1, 'root', 'root', 'root'),
  (2, 'gopher', 'gopher', 'gopher');

-- +goose Down
DROP TABLE users;