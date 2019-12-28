CREATE TABLE Accounts (
  id              SERIAL PRIMARY KEY,
  name            VARCHAR(100) NOT NULL,
  balance         int NULL
);

ALTER USER postgres WITH PASSWORD 'testing';
