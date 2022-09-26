CREATE TABLE users(
    id int,
    balance decimal
);

CREATE TABLE transactions(
    id  PRIMARY KEY serial,
    from_id int,
    to_id int,
    amount decimal
);