sudo -u postgres createuser -s `whoami`
sudo -u postgres createdb station
psql -d station -a -f pgsetup.sql
