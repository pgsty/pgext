## Usage

Sources:

- [Version 0.0.4 SQL objects](https://github.com/adjust/pg-language/blob/36e553fec17c126be3fed68b28c8c13f7f88299a/language--0.0.4.sql)
- [Generated language mapping](https://github.com/adjust/pg-language/blob/36e553fec17c126be3fed68b28c8c13f7f88299a/languages)
- [PGXN README](https://pgxn.org/dist/language/0.0.4/README.html)

`language` defines a compact, one-byte, application-specific `language` type. Its accepted values are generated from the upstream `languages` mapping rather than from PostgreSQL procedural-language catalog entries.

```sql
CREATE EXTENSION language;
SELECT supported_languages();
CREATE TABLE profile (
  profile_id bigint PRIMARY KEY,
  preferred_language language NOT NULL
);
```

The extension supplies text/binary I/O, comparison operators, and default B-tree and hash operator classes, so values can be sorted, indexed, and compared. Use `supported_languages()` to discover the exact mapping shipped by the installed build instead of assuming ISO code coverage.

The internal representation is a generated one-byte code. Changing or regenerating the mapping can reinterpret stored bytes, which makes binary compatibility and logical meaning dependent on the exact build. Pin the source revision across every node, verify dump/restore and replication behavior before upgrades, and do not use the type as a general standards registry without auditing the mapping.
