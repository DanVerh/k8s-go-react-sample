FROM postgres:alpine
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=admin
EXPOSE 5432
COPY ./dbinit.sql /docker-entrypoint-initdb.d/create_tables.sql
VOLUME /var/lib/postgresql/data