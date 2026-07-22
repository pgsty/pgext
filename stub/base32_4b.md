## Usage

Sources:

- [Official SQL definition](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/sql/base32_4b.sql)
- [Official README](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/README.md)
- [Official control file](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/base32_4b.control)

`base32_4b` defines a pass-by-value Base32 type with a four-byte internal representation. Use it only for identifiers accepted by the extension's input function; it is a compact scalar type, not a general encoder for arbitrary byte strings.

### Core Workflow

```sql
CREATE EXTENSION base32_4b;

CREATE TABLE compact_ids (
    id base32_4b PRIMARY KEY,
    payload jsonb NOT NULL
);

INSERT INTO compact_ids (id, payload)
VALUES ($1::text::base32_4b, '{"state":"ready"}');

SELECT id::text, payload
FROM compact_ids
WHERE id = $1::text::base32_4b;
```

Validate application input with an explicit cast before bulk loading. The text casts are assignment casts rather than implicit casts, which helps keep comparisons type-specific.

### Important Objects

- `base32_4b` is the four-byte scalar type with text and binary input/output support.
- Comparison operators `<`, `<=`, `=`, `<>`, `>=`, and `>` provide ordering and equality.
- `base32_4b_ops` is the default B-tree operator class, so primary keys and B-tree indexes are supported.
- `min` and `max` aggregates are defined for the type.
- `base32_4b_to_text` and `base32_4b_from_text` implement assignment casts to and from `text`.

### Caveats

The SQL surface defines no hash operator class, arithmetic, or general binary-to-Base32 conversion. Use a B-tree for indexed equality and ordering, and confirm that the extension's ordering matches the application's desired textual ordering.

The repository has only an initial 2015-era implementation and minimal documentation. Test accepted alphabet, case handling, invalid and boundary input, binary send/receive, index ordering, dump/restore, and the target PostgreSQL major version before adopting it. Four bytes of datum storage do not guarantee four bytes per table row because tuple headers, null bitmaps, and alignment also contribute.
