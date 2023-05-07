#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table todo (
    id int(10) AUTO_INCREMENT NOT NULL primary key,
    text varchar(1000) NOT NULL,
    due_date datetime NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    );"
$CMD_MYSQL -e  "insert into todo (text, due_date) values ('テストのメモ1', cast('2023-05-09 24:00:00' as datetime));"