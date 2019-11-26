CREATE DATABASE gocrawler;
use gocrawler;

CREATE TABLE hatebu (
  id int(11) unsigned not null auto_increment,
  title varchar(255) not null,
  link varchar(255) not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  primary key (id)
);


ALTER TABLE hatebu ADD UNIQUE(link)


