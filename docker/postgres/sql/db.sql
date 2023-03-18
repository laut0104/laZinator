CREATE TABLE IF NOT EXISTS users(
    id serial NOT NULL, 
    lineuserid varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (lineuserid)
);

CREATE TABLE IF NOT EXISTS clothes(
    id serial NOT NULL, 
    userid integer NOT NULL,
    cloth varchar(255) NOT NULL,
    details TEXT[] NOT NULL,
    weather TEXT[] NOT NULL,
    temperature TEXT[] NOT NULL,
    events TEXT[] NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT menus_userid_fkey
    FOREIGN KEY (userid)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
)
