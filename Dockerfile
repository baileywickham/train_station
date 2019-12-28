FROM postgres

ENV POSTGRES_PASSWORD testing
ENV POSTGRES_DB station
COPY ./scripts/pgsetup.sql /docker-entrypoint-initdb.d/
