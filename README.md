# NOTES

## Create New Postgre User

```sql
CREATE ROLE greenlight WITH LOGIN PASSWORD 'greenlight';
```

## Grant Access To New User

```sql
psql -U postgres    -- login as super user
\c greenlight;      -- navigate to target database
GRANT CREATE ON SCHEMA public TO greenlight;
```
