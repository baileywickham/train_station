sudo -u postgres createdb station
psql -h localhost -d station -U postgres -W postgres -a -f pgsetup.sql
