## Usage

Sources:

- [Official README](https://github.com/madebykite/pgunicoll/blob/af4e0c71fa3fe109ecca9acb2b254af1d545c2c8/README.md)
- [Extension SQL](https://github.com/madebykite/pgunicoll/blob/af4e0c71fa3fe109ecca9acb2b254af1d545c2c8/sql/pgunicoll.sql)
- [ICU-backed implementation](https://github.com/madebykite/pgunicoll/blob/af4e0c71fa3fe109ecca9acb2b254af1d545c2c8/pgunicoll.c)

`pgunicoll` generates ICU Unicode Collation Algorithm sort keys as `bytea`. It is useful when an application needs an explicit ICU locale ordering on a PostgreSQL text expression, but it does not create a PostgreSQL collation object or automatically affect ordinary `ORDER BY` expressions.

### Core Workflow

Create the extension and order rows by the sort key returned from `pgunicoll(text, text)`.

```sql
CREATE EXTENSION pgunicoll;

SELECT name
FROM person
ORDER BY pgunicoll(name, 'de_DE');
```

The locale argument defaults to an empty string. Empty string and `root` both request ICU's root collator. The implementation enables ICU numeric collation, so digit sequences compare by numeric value rather than only byte-by-byte text order.

For repeated searches, use an expression index with exactly the same locale expression used by queries.

```sql
CREATE INDEX person_name_de_idx
ON person (pgunicoll(name, 'de_DE'));

SELECT name
FROM person
ORDER BY pgunicoll(name, 'de_DE');
```

### Important Object

- `pgunicoll(text, text DEFAULT '') RETURNS bytea`: returns an ICU sort key; it is declared STRICT and IMMUTABLE.

### Stability and Compatibility

The function always interprets input bytes as UTF-8, regardless of the database encoding. It is therefore unsuitable for non-UTF-8 databases. Locale availability and ordering depend on the linked ICU library.

ICU upgrades can change sort keys. Rebuild every expression index and refresh any stored or materialized key after upgrading ICU; keep compatible ICU versions across primary, replicas, and restore targets. Although the function is declared IMMUTABLE for index use, its result is only stable while the ICU version and locale data remain fixed. Store the original text as the source of truth, not the generated key.
