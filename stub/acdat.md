## Usage

Sources:

- [Official release v0.1.1](https://github.com/pgsty/acdat/releases/tag/v0.1.1)
- [Official README v0.1.1](https://github.com/pgsty/acdat/blob/v0.1.1/README.md)
- [Extension control file](https://github.com/pgsty/acdat/blob/v0.1.1/acdat.control)
- [Versioned installation SQL](https://github.com/pgsty/acdat/blob/v0.1.1/sql/acdat--0.1.1.sql)
- [Official usage guide](https://github.com/pgsty/acdat/blob/v0.1.1/docs/USAGE.md)
- [Runnable SQL demonstration](https://github.com/pgsty/acdat/blob/v0.1.1/examples/demo.sql)

`acdat` 0.1.1 compiles a large dictionary of exact literal patterns into an immutable Aho-Corasick Double-Array machine, then scans each `text` or `bytea` value once for matching or replacement. It is designed for stable, repeatedly used dictionaries such as policy rules, indicators of compromise, entity names, and redaction aliases.

### Core Workflow

Create the extension, compile a dictionary, and reuse the resulting `acdat.machine` value across many inputs:

```sql
CREATE EXTENSION acdat;

WITH machine AS (
    SELECT acdat.compile(
        ARRAY['he', 'she', 'his', 'hers'],
        ARRAY[1, 2, 3, 4]::bigint[]
    ) AS value
)
SELECT acdat.contains('ushers', value) AS matched,
       acdat.info(value)->>'pattern_count' AS patterns
FROM machine;
```

For production dictionaries, the source rules should stay in an application-owned table. The aggregate overload of `acdat.compile()` can build one deterministic machine directly from pattern, ID, replacement, and priority rows; compile once and scan many values.

### Matching and Replacement

`acdat.contains()` stops after the first hit. `acdat.matches()` returns `acdat.hit` rows with the pattern ID, byte and character coordinates, and priority. `acdat.replace()` applies literal, non-recursive replacements:

```sql
WITH machine AS (
    SELECT acdat.compile(
        ARRAY['病毒', '特征码', '病毒特征码'],
        ARRAY[10, 11, 12]::bigint[],
        ARRAY['[VIRUS]', '[SIGNATURE]', '[IOC]'],
        ARRAY[20, 20, 5]::integer[]
    ) AS value
)
SELECT *
FROM acdat.matches('发现病毒特征码', (SELECT value FROM machine), 'all_overlapping');

SELECT acdat.replace(
    'aaa',
    acdat.compile(
        ARRAY['a', 'aa', 'aaa'],
        ARRAY[1, 2, 3]::bigint[],
        ARRAY['[x]', '[yy]', '[zzz]']
    ),
    'leftmost_longest'
);
```

The match policies are `all_overlapping`, `leftmost_longest`, and `leftmost_priority`. Replacement accepts only a non-overlapping policy. Use `acdat.info()` to inspect a compiled machine and the export, validation, import, and fingerprint functions when moving or checking artifacts.

`acdat.matches()` defaults `max_matches` to 10000, and `acdat.replace()` defaults `max_output_bytes` to 268435456. Set tighter limits for untrusted or high-hit inputs so match enumeration and replacement output stay bounded.

### Managed Dictionaries

The optional catalog layer publishes immutable, content-addressed builds and atomically selects one active build. Its control functions use `SECURITY INVOKER` and are not executable by `PUBLIC`:

```sql
WITH machine AS (
    SELECT acdat.compile(pattern, pattern_id)
    FROM app_keyword
    WHERE enabled
), published AS (
    SELECT acdat.publish('moderation', 1, machine) AS build_id
    FROM machine
)
SELECT acdat.activate('moderation', build_id)
FROM published;

SELECT name, version, build_id, machine
FROM acdat.active_machine
WHERE name = 'moderation';
```

Application tables remain the source of truth. Logical dumps include catalog metadata and active machine payloads, but not every historical artifact, so retain the source patterns required to rebuild retired or inactive versions.

### Compatibility and Safety

Version 0.1.1 is tested on PostgreSQL 14 through 18. It needs no preload or server restart, has no external extension dependency, and defines no GUC. The control file fixes the schema to `acdat`, sets `relocatable = false` and `trusted = false`, so `CREATE EXTENSION` requires a superuser.

The 0.1.1 release preserves the 0.1.0 SQL API and format-major-1 machine compatibility and ships the `0.1.0 -> 0.1.1` extension update path. It also adds cancellable compilation, a conservative build-work budget, and a faster materialized scan path without changing the stored machine contract.

ACDAT indexes the pattern dictionary, not the document table: scanning a large existing table still reads its candidate rows. Matching is exact and case-sensitive; the extension does not provide regular expressions, fuzzy matching, tokenization, automatic case folding, Unicode normalization, or a document-side index. The text engine supports UTF-8 and single-byte server encodings, while binary data should use the bytea interface. Materialize `(document_id, pattern_id)` hits into an application table when repeated reverse lookup is required.

The compiled format is self-describing and checksummed, and imported artifacts are validated before use. Inventory dependencies before uninstalling: `DROP EXTENSION acdat` removes managed dictionary state, while adding `CASCADE` can also remove user columns or other objects that depend on `acdat.machine`.
