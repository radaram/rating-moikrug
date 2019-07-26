CREATE TABLE company (
    name varchar(1024) NOT NULL,  
    site varchar(1024),
    about varchar(1024),
    rating float,
    address varchar(1024),
    score int,
    link  varchar(1024),
    employees_left json,
    employees_came json,
    id serial NOT NULL PRIMARY KEY
);
