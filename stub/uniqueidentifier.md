## Usage

Sources:

- [Official README](https://github.com/autodatadirect/uniqueidentifier/blob/85976b564e5cd3362c30ba388b95c3df19bc4b34/README.uniqueidentifier)
- [Official extension SQL](https://github.com/autodatadirect/uniqueidentifier/blob/85976b564e5cd3362c30ba388b95c3df19bc4b34/uniqueidentifier--1.0.sql)
- [Official C implementation](https://github.com/autodatadirect/uniqueidentifier/blob/85976b564e5cd3362c30ba388b95c3df19bc4b34/uniqueidentifier.c)

`uniqueidentifier` is a legacy 16-byte UUID-like type with generation, text conversion, comparisons, concatenation, and B-tree/hash operator classes. It was originally written in 2001–2003 and later packaged as a PostgreSQL extension; prefer PostgreSQL's built-in `uuid` type for new schemas unless compatibility with this exact legacy type is required.

### Core Workflow

The native library depends on the e2fsprogs `libuuid` implementation. After installing the matching binary, create the extension and use `newid()` for generated defaults:

```sql
CREATE EXTENSION uniqueidentifier;

CREATE TABLE legacy_entities (
  id uniqueidentifier PRIMARY KEY DEFAULT newid(),
  payload text
);

INSERT INTO legacy_entities(payload) VALUES ('example');

SELECT id, text(id)
FROM legacy_entities;
```

`newid(text)` parses text explicitly. The extension also declares implicit casts in both directions:

```sql
SELECT newid('550e8400-e29b-41d4-a716-446655440000');
SELECT '550e8400-e29b-41d4-a716-446655440000'::uniqueidentifier;
```

### SQL Surface

- `newid()` generates a value using `libuuid` and is volatile.
- `newid(text)` parses text; `text(uniqueidentifier)` formats a value.
- `<`, `<=`, `=`, `<>`, `>=`, and `>` support comparisons.
- Default B-tree and hash operator classes support primary keys, equality lookup, ordering, and indexes.
- `||` overloads concatenate a value with `text` in either argument order.

### Migration and Compatibility

- This type is not PostgreSQL's built-in `uuid`, even though both occupy 16 bytes and use familiar text. Test semantic, ordering, binary, replication, driver, dump/restore, and cast differences before migrating between them.
- Bidirectional implicit casts with `text` can change overload resolution and hide malformed boundaries. Use explicit casts in new SQL and test expressions after upgrades.
- The reviewed install SQL hard-codes `public` for concatenation operators and changes helper-function ownership to the role `postgres`. Review the script against the intended schema and role model before installation.
- Do not infer current PostgreSQL compatibility from the README's PostgreSQL 9.1 packaging note. Build and test the exact source/package against the target major version.
- If compatibility is not required, use core `uuid` with the server's supported UUID generator and migrate with an explicit, validated text conversion.
