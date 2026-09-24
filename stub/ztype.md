## Usage

Sources:

- [README](https://github.com/xvaara/pg_ztype/blob/ff6fcccda0c5f6bfea59cd0acd2a14ef8cf88d85/README.md)
- [Control](https://github.com/xvaara/pg_ztype/blob/ff6fcccda0c5f6bfea59cd0acd2a14ef8cf88d85/ztype.control)
- [SQL](https://github.com/xvaara/pg_ztype/blob/ff6fcccda0c5f6bfea59cd0acd2a14ef8cf88d85/ztype--0.9.sql)
- [Source](https://github.com/xvaara/pg_ztype/blob/ff6fcccda0c5f6bfea59cd0acd2a14ef8cf88d85/ztype.c)

`ztype`, distributed as pg_ztype, provides zstd-compressed `ztext`, `zjsonb`, and `zbytea` column types. Version 0.9 is a pre-release source snapshot: its storage and type-modifier formats are not frozen, and it supplies no extension upgrade scripts. Retain a working logical export path before storing durable data.

### Core Workflow

The upstream requirements are PostgreSQL 18 or newer and libzstd 1.5 or newer; upstream tests PostgreSQL 18 and 19. Installation requires a superuser, with PL/pgSQL available for the dictionary helpers. The extension creates its types in `public` and administration objects in `ztype`; it is not relocatable and does not require preloading.

```sql
CREATE EXTENSION ztype;
CREATE TABLE messages (
  id bigint PRIMARY KEY,
  body ztext(6),
  metadata zjsonb(6)
);
INSERT INTO messages VALUES (1, 'hello', '{"source":"email"}');
SELECT body::text, metadata ->> 'source' FROM messages;
CREATE INDEX messages_metadata ON messages USING gin ((metadata::jsonb));
```

The modifier selects compression level 1–22 and optionally a registered dictionary name or slot. A value shorter than 64 bytes, or one that does not shrink, is stored raw. Cast to `text`, `jsonb`, or `bytea` when invoking base-type functions. Use typed input parameters to avoid compressing an untyped value twice when the destination has a non-default policy.

### Dictionaries and Inspection

An administrator can train a dictionary from representative data and then name it in a new column:

```sql
SELECT ztype.train_and_add('message-json',
  'SELECT metadata::jsonb FROM messages LIMIT 20000');
CREATE TABLE archive (metadata zjsonb(6, 'message-json'));
SELECT ztype.inspect(metadata), ztype.validate(metadata) FROM messages;
SELECT * FROM ztype.dictionary_inventory;
SELECT * FROM ztype.column_policies;
```

Training needs at least eight nonempty samples; the query must return one `text`, `bytea`, or `jsonb` column. Train JSON dictionaries on the binary `jsonb` value. `ztype.validate` returns NULL for a valid value or a diagnostic message; `ztype.inspect` reports the stored codec, size, policy and dictionary identity without full decoding.

| Interface | Purpose |
| --- | --- |
| `ztype.train_dictionary`, `ztype.add_dictionary`, `ztype.import_dictionary` | Train, register, or import dictionary bytes at an explicit slot |
| `ztype.recompress`, `ztype.matches_policy` | Produce a value under a policy or inspect whether it matches |
| `ztype.set_column_policy`, `ztype.finish_column_policy` | Change the policy without an immediate full rewrite, then validate completion |
| `raw_length`, `prefix` | Inspect logical byte length or read a character-safe text prefix |
| `ztype.zstd` | Return a zstd frame for clients that can decode it |
| `ztype.build_info`, `ztype.dictionary_cache_stats`, `ztype.decode_cache_stats` | Inspect library/format versions and per-backend caches |

### Indexing, Privileges and Maintenance

- Equality, grouping and hash indexes work on compressed types. There is no ordering operator or B-tree operator class: sort the cast value. JSON read operators work on `zjsonb`, and GIN indexes use the cast expression shown above.
- Dictionary registration requires explicitly delegated function execution; ordinary column users do not need registry access. Dictionary bytes may reveal training data. The database-wide registry is append-only, and slots must never be reused.
- Reads of dictionary-compressed frames fail if the dictionary is missing. Physical replication carries the registry; logical replication and restore require the documented dictionary/slot coordination. The supplied `ztype-sync` helper transports and checks the registry. Never replace one dictionary with unrelated bytes at the same slot.
- Changing a column's type modifier normally rewrites its table under an exclusive lock. The staged policy helpers leave existing rows pending until recompressed; completion still scans under an exclusive lock. Keep backups and budget for WAL and table bloat during a batched rewrite.
- `ztype.dictionary_cache_size` defaults to 64 MB per backend. It limits cached dictionaries, not all zstd memory: `work_mem` does not cap codec contexts or training buffers. Compression costs CPU on writes and decompression on reads; choose a level using your own workload.
