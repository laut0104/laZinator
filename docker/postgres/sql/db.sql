CREATE TABLE IF NOT EXISTS users(
    id serial NOT NULL, 
    lineuserid varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (lineuserid)
);
