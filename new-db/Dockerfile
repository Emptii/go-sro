FROM mysql:8.0

COPY new-db/my.cnf /etc/mysql/conf.d/

COPY new-db/create_databases.sql /docker-entrypoint-initdb.d/
COPY new-db/fill_databases.sql /docker-entrypoint-initdb.d/
COPY new-db/MasterServer /docker-entrypoint-initdb.d/MasterServer
COPY new-db/GameServer /docker-entrypoint-initdb.d/GameServer

COPY new-db/Media /data/Media
