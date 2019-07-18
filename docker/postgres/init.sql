CREATE TABLE company (
    name varchar(512) NOT NULL,  
    site varchar(512),
    about varchar(1024),
    rating float,
    address varchar(512),
    score int,
    link  varchar(512),
    employees_left json,
    employees_came json,
    id serial NOT NULL PRIMARY KEY
);
