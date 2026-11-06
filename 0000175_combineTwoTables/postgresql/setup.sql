CREATE TEMP TABLE person (
    personid INTEGER PRIMARY KEY,
    lastname VARCHAR(255) NOT NULL,
    firstname VARCHAR(255) NOT NULL
);

CREATE TEMP TABLE address (
    addressid INTEGER PRIMARY KEY,
    personid INTEGER NOT NULL,
    city VARCHAR(255),
    state VARCHAR(255)
);

INSERT INTO person (personid, lastname, firstname) VALUES
    (1, 'Wang', 'Allen'),
    (2, 'Alice', 'Bob');

INSERT INTO address (addressid, personid, city, state) VALUES
    (1, 2, 'New York City', 'New York'),
    (2, 3, 'Leetcode', 'California');
