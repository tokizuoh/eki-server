FROM mysql:8.0
WORKDIR ./
COPY init /docker-entrypoint-initdb.d/
ADD ./my.cnf /etc/mysql/my.cnf
