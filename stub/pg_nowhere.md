## Usage

Sources:

- [Official upstream README](https://github.com/viveknathani/pg_nowhere/blob/0d3ae6a044a4954bd2e65cee01d47fc03394e5b6/README.md)
- [Official extension control file (pg_nowhere.control)](https://github.com/viveknathani/pg_nowhere/blob/0d3ae6a044a4954bd2e65cee01d47fc03394e5b6/pg_nowhere.control)
- [Official extension SQL (pg_nowhere--0.1.sql)](https://github.com/viveknathani/pg_nowhere/blob/0d3ae6a044a4954bd2e65cee01d47fc03394e5b6/pg_nowhere--0.1.sql)

`pg_nowhere` — A PostgreSQL extension to disallow UPDATE and DELETE queries without a WHERE clause. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_nowhere;

-- Create test table
CREATE TABLE users (id SERIAL, name TEXT);
INSERT INTO users (name) VALUES ('Alice'), ('Bob');

-- This should work
UPDATE users SET name = 'Charlie' WHERE id = 1;

-- This should fail
UPDATE users SET name = 'Dave';
DELETE FROM users;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
