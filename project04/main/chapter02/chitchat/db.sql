drop table if exists posts;
drop table if exists threads;
drop table if exists sessions;
drop table if exists users;


create table users (
  id         int(10) not null primary key auto_increment,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null
);

create table sessions (
  id         int(10) not null primary key auto_increment,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    int(10) not null,
  created_at timestamp not null
);

create table threads (
  id         int(10) not null primary key auto_increment,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    int(10) not null,
  created_at timestamp not null
);

create table posts (
  id         int(10) not null primary key auto_increment,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    int(10) not null,
  thread_id  int(10) not null,
  created_at timestamp not null
);
