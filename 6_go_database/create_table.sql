CREATE TABLE customer
(
  id    VARCHAR(100) NOT NULL,
  name  VARCHAR(100) NOT NULL,
  PRIMARY KEY(id)
) ENGINE = InnoDB;

SELECT * FROM customer;

DELETE FROM customer;

ALTER TABLE customer
  ADD COLUMN email      VARCHAR(100),
  ADD COLUMN balance    INTEGER DEFAULT 0,
  ADD COLUMN rating     DOUBLE DEFAULT 0.0,
  ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  ADD COLUMN birth_date DATE,
  ADD COLUMN married    BOOLEAN DEFAULT false;

DESC customer;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
  VALUES  ('1', 'Zulfikar', 'isnaen@gmail.com', 10000, 95, '1999-09-09', 1),
          ('2', 'Budi', 'budi@gmail.com', 90000, 88, '1999-03-09', 1),
          ('3', 'Jack', 'jack@gmail.com', 1000000, 78, '1999-02-09', 1);

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
  VALUES  ('1', 'Zulfikar', NULL, 10000, 95, NULL, 1);

CREATE TABLE user
(
  username    VARCHAR(100) NOT NULL,
  password    VARCHAR(100) NOT NULL,
  PRIMARY KEY(username)
) ENGINE = InnoDB;

INSERT INTO user(username, password) VALUES('admin', 'admin');

CREATE TABLE comments
(
  id      INT NOT NULL AUTO_INCREMENT,
  email   VARCHAR(100) NOT NULL,
  comment TEXT,
  PRIMARY KEY(id)
) ENGINE = InnoDB;