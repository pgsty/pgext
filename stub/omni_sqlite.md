


## Usage

> [omni_sqlite: Embedded SQLite](https://docs.omnigres.org/omni_sqlite/sqlite/)

The `omni_sqlite` extension provides SQLite databases as a first-class PostgreSQL data type. It is a templated extension.

### Create an SQLite Database

```sql
SELECT 'CREATE TABLE user_config (key text, value text)'::omni_sqlite.sqlite;
```

### Store SQLite in a Column

```sql
CREATE TABLE customer (
    id   bigserial PRIMARY KEY,
    name text NOT NULL,
    data omni_sqlite.sqlite DEFAULT 'CREATE TABLE user_config (key text, value text);'
);
```

### Execute Statements

```sql
UPDATE customer
SET data = omni_sqlite.sqlite_exec(data,
    $$INSERT INTO user_config VALUES ('color', 'blue')$$)
RETURNING *;
```

With bound parameters:

```sql
SELECT omni_sqlite.sqlite_exec('CREATE TABLE tab (val text)',
    'INSERT INTO tab VALUES ($1)', row('hello'));
```

### Query SQLite Data

```sql
SELECT * FROM omni_sqlite.sqlite_query(
    (SELECT data FROM customer),
    'SELECT rowid, key, value FROM user_config')
AS (id bigint, key text, value text);
```

### Serialization

```sql
-- Serialize to bytea:
SELECT omni_sqlite.sqlite_serialize('CREATE TABLE foo (x)');
-- Deserialize from bytea:
SELECT omni_sqlite.sqlite_deserialize(bytea_data);
```

Useful for multitenancy (isolated SQLite databases per row) and peer synchronization.
