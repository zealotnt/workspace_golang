CREATE TABLE weather
(
  id SERIAL,
  city varchar(80), 
  temp_lo int,
  temp_hi int,
  prcp real,
  date date,
  PRIMARY KEY     (id)
);