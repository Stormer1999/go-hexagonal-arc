# docker command

$ docker pull mysql
$ docker images
$ docker ps -a
$ cat /etc/os-release

# run docker app

$ docker run --name dolphin --rm -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql
$ docker run --name dolphin --rm -v mysqldatavolume:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql
$ docker exec -it dolphin bash

# login to database

$ mysql -u root -p
$ show databases;
$ create database drink;
$ use drink
$ create table menu(id int, descr varchar(50), price int);
$ show tables;
$ insert into menu values (1, 'coffee', 40), (2, 'tea', 30);

$ create database banking;
$ create table customer(customer_id int,name VARCHAR(50), date_of_birth VARCHAR(50), city VARCHAR(50), zipcode VARCHAR(50), status int);
$ insert into customer values (1, 'hello', 'dob', 'city', 'zipcode', 0);

# stop docker

$ docker stop dolphin