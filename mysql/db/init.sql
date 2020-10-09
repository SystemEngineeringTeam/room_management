use production_db;

create table emails(
    id int auto_increment primary key not null,
    email varchar(45) not null,
    password varchar(256) not null
);

create table users(
    id int auto_increment primary key not null,
    student_number varchar(10) not null,
    name varchar(45) not null,
    isEntry tinyint(1),
    email_id int,
    foreign key (email_id) references emails(id) on delete cascade
);

create table cards(
    id int auto_increment primary key not null,
    uid varchar(45) not null,
    user_id int,
    foreign key (user_id) references users(id) on delete cascade
);

create table logs(
    id int auto_increment primary key not null,
    cards_id int not null,
    card_read_datetime datetime,
    isEntry tinyint(1),
    foreign key (cards_id) references cards(id) on delete cascade
);

insert into
    emails(id,email,password)
values
    (1,"hoge@hoge.jp","4c716d4cf211c7b7d2f3233c941771ad0507ea5bacf93b492766aa41ae9f720d");

insert into
    users(id,student_number,name,isEntry,email_id)
values
    (1,"K12345","Hoge",1,1);

insert into
    cards(id,uid,user_id)
values
    (1,"1 2 3 4",1);

insert into
    logs(id,cards_id,card_read_datetime,isEntry)
values
    (1,1,Now(),1);