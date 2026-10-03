## Usage

Sources:

- [Official README](https://github.com/pyhalov/mufdw/blob/5dabf4a2dde02e5d9a3fdcc1b59ea6f1790c4d96/README.md)
- [Extension control file](https://github.com/pyhalov/mufdw/blob/5dabf4a2dde02e5d9a3fdcc1b59ea6f1790c4d96/mufdw.control)
- [Installation SQL](https://github.com/pyhalov/mufdw/blob/5dabf4a2dde02e5d9a3fdcc1b59ea6f1790c4d96/mufdw--0.1.sql)

`mufdw` is a standalone educational foreign data wrapper that reads an ordinary local table through SPI. Use it to study FDW planning and execution; it is not a connector to a separate database.

### Core Workflow

The upstream example creates a loopback server, user mapping, local table and matching foreign table.

```sql
CREATE EXTENSION mufdw;
CREATE SERVER loopback FOREIGN DATA WRAPPER mufdw;
CREATE USER MAPPING FOR PUBLIC SERVER loopback;
CREATE TABLE players(id int PRIMARY KEY, nick text, score int);
CREATE FOREIGN TABLE f_players(id int, nick text, score int)
  SERVER loopback OPTIONS (table_name 'players', schema_name 'public');
SELECT * FROM f_players;
```

### Boundaries

`schema_name` and `table_name` select the underlying local relation; maintain compatible column definitions. The source defines `mufdw_handler()` and `mufdw_validator(text[], oid)` and control version `0.1`. This is a demonstration implementation without a declared supported-major matrix or production guarantees. Do not use foreign-table grants to bypass the intended privileges on underlying data.
