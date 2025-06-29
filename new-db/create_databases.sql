drop database if exists `DB_Manager`;
drop database if exists `DB_Shard`;

#
#	INITIALIZE MANAGER DATABASE
#

create database `DB_Manager`;
use `DB_Manager`;

source /docker-entrypoint-initdb.d/MasterServer/create_tables.sql

#
#	INITIALIZE SHARD DATABASE
#

create database `DB_Shard`;
use `DB_Shard`;

source /docker-entrypoint-initdb.d/GameServer/create_tables.sql
