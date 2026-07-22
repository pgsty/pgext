## Usage

Sources:

- [Official extension control file](https://github.com/zilder/pg_lz4/blob/f3259686848a61c544dbe58df5627ed3d77a2144/pg_lz4.control)
- [Official upstream README](https://github.com/zilder/pg_lz4/blob/f3259686848a61c544dbe58df5627ed3d77a2144/README.md)
- [Required PostgreSQL patch discussion](https://www.postgresql.org/message-id/20190315125203.5da43edb%40ildus-work.localdomain)

`pg_lz4` 1.0 registers LZ4 as a column compression method. This historical implementation requires a development PostgreSQL tree with the linked custom-compression patch and the system `liblz4` library; it is not an extension for an unmodified supported server.

### Core Workflow

After building the patched server and installing the matching extension binary, create the extension and choose its compression method per column.

```sql
CREATE EXTENSION pg_lz4;

CREATE TABLE messages (
  id bigint GENERATED ALWAYS AS IDENTITY,
  body text COMPRESSION pg_lz4
);

ALTER TABLE messages
  ALTER COLUMN body SET COMPRESSION pg_lz4 WITH (acceleration '8');
```

The `acceleration` option adjusts the LZ4 speed/compression tradeoff. Changing the column setting affects newly written values; it does not by itself rewrite existing rows.

### Compatibility Boundary

The upstream README targets a 2019 development patch rather than the later core compression-method API. Do not assume compatibility with stock PostgreSQL or with built-in LZ4 TOAST support. Use the exact patched server ABI, test backup and restore portability, and plan a data rewrite before moving compressed values to a server that lacks this method.
