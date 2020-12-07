FROM mysql:latest
ENV MYSQL_ROOT_PASSWORD=root \
    MYSQL_DATABASE=mydb \
    MYSQL_USER=human \
    MYSQL_PASSWORD=mysql

ADD schema.sql.gz /docker-entrypoint-initdb.d