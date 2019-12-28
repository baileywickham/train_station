FROM postgres

ENV POSTGRES_USER: postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB station
COPY ./scripts/pgsetup.sql /docker-entrypoint-initdb.d/
