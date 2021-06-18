create table goods(
                      id int auto_increment,
                      type smallint,
                      `name` varchar(10),
                      description text,
                      price float,
                      picture_url text,
                      discount smallint default 10 comment '折扣，代表几折',
                      inventory int comment '库存',
    check (discount<10 and discount>0),
    check (price>0),
    check (inventory>0),
    primary key (id)
);
create table client(
                       id int  auto_increment,
                       `name` varchar(10) ,
                       tel char(11) unique ,
                       passwd varchar (16),
    check (LENGTH(tel)=11),
    check (LENGTH(passwd)>8),
    primary key (id)
);
create table `order`(
                        id int auto_increment,
                        goods_id int,
                        client_id int,
                        `num` int,
                        `status` smallint,
    constraint FK_GOODS_ID foreign key (goods_id) references goods(id),
    constraint FK_CLIENT_ID foreign key (client_id) references client(id),
    primary key (id)
);