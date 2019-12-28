sudo -u postgres createdb station
PGPASSWORD=postgres psql -h localhost -d station -U postgres -a -f ./scripts/pgsetup.sql
