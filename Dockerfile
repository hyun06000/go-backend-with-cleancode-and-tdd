FROM mysql:latest

COPY ./mysql/initdb.sql /docker-entrypoint-initdb.d/

ENV MYSQL_ROOT_PASSWORD 1234

EXPOSE 3306
