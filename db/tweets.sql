CREATE TABLE tweets
(
ID INT NOT NULL,
CONTENT VARCHAR(140) NOT NULL,
POSTEDAT TIMESTAMP NOT NULL,
DEVICE VARCHAR(50),
AUTHOR INT REFERENCES USERs(ID),
PRIMARY KEY (ID, AUTHOR)
)
