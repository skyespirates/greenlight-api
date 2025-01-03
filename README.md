# NOTES FOR FUTURE ME

I make this notes as a guide for my self in the future, in case I want to run this service I will know what it needs to run the app properly. Because in the prior projects I struggle a lot just to run that project because there is no guide how to run it. Hopefully with this note will help me or other to run the service.

## DEPENDENCIES

1. go version 1.23 [golang](https://go.dev/dl/)
2. postgres version 17.2 (run in docker btw) [postgres](https://hub.docker.com/_/postgres)
3. migrate cli version 4.18 [migrate](https://github.com/golang-migrate/migrate/releases)
4. air version 1.61 (watch file changes) -> optional [air](https://github.com/air-verse/air)
5. make or GNU make version 3.82 (automate script) [make](https://sourceforge.net/projects/gnuwin32/files/make/3.81/make-3.81-bin.zip/download?use_mirror=onboardcloud&download=) -> optional

## POSTGRES

### Create New Postgre User

```sql
CREATE ROLE greenlight WITH LOGIN PASSWORD 'greenlight';
```

### Grant Access To New User

```sql
psql -U postgres    -- login as super user
\c greenlight;      -- navigate to target database
GRANT CREATE ON SCHEMA public TO greenlight;
```

### Migration

if you have `make` installed in your system, you can just run the make script, like so:

```sh
make -v # check whether make installed or not
make migrate-run # run the migrate script, but make sure you are able to connect to postgre sql first
```

if you dont have make installed, just run the migration script manually, LOL awokawok, by nagivate to `migrations` directory and copy paste the sql script to your postgresql client sequentially, and you'll be good

NOTE: to connect postgresql using `database/sql` requires you to have a user AND password, you could'nt connect to it only using username, so make sure you create a new user and set the password as well, you will need that username and password as your dsn to connect to posgresql

## RUN SERVICE

if you resolved the dependencies and migration issue, you can run the script using this command

```
go run ./cmd/api
```
