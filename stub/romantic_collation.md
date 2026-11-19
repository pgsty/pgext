## Usage

Sources:

- [Official source](https://github.com/pmarreck/romantic_collation/blob/0cf958291205b2cbdd8ae48a95ea372deaad7c8f/docs/DATABASES.md)
- [Official source](https://github.com/pmarreck/romantic_collation/blob/0cf958291205b2cbdd8ae48a95ea372deaad7c8f/db/romantic_collation.control)
- [Official source](https://github.com/pmarreck/romantic_collation/blob/0cf958291205b2cbdd8ae48a95ea372deaad7c8f/LICENSE)
- [Official source](https://github.com/pmarreck/romantic_collation/blob/0cf958291205b2cbdd8ae48a95ea372deaad7c8f/db/romantic_collation--1.0.sql)

romantic_collation exposes reproducible natural ordering as bytea sort keys. It does not register a PostgreSQL COLLATE provider. The native adapter requires UTF8 database encoding and a library built against the server's headers; no fixed supported-major list is declared.

### Generate and index keys

```sql
CREATE EXTENSION romantic_collation;
SELECT label FROM (VALUES ('item10'), ('item2')) AS x(label)
ORDER BY romantic_key_v1(label, 0);
CREATE TABLE labels(label text);
CREATE INDEX labels_natural ON labels(romantic_key_v1(label, 0));
```

romantic_key_v1(text, integer DEFAULT 0) is immutable, strict and parallel safe. Use the same option bits in queries and indexed expressions: codepoint mode is 1, decimal parsing 8, scientific notation 32 and Roman-numeral parsing 64. The default provides the library's natural-order semantics.

### Persistence

Keep the key version and options fixed for stored/indexed keys. Recompute affected keys and rebuild expression indexes when changing key semantics; earlier scientific-notation keys are not interchangeable. Equal keys do not prove byte-identical input. Long generated keys can exceed PostgreSQL B-tree key limits. No preload is required.
