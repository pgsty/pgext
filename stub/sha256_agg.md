## Usage

Sources:

- [README.md](https://github.com/danolivo/sha256_agg/blob/bc0ead44962c4afa53c2b0446e5d73492f422405/README.md)
- [sha256_agg--1.0.sql](https://github.com/danolivo/sha256_agg/blob/bc0ead44962c4afa53c2b0446e5d73492f422405/sha256_agg--1.0.sql)
- [Makefile](https://github.com/danolivo/sha256_agg/blob/bc0ead44962c4afa53c2b0446e5d73492f422405/Makefile)
- [sha256_agg.control](https://github.com/danolivo/sha256_agg/blob/bc0ead44962c4afa53c2b0446e5d73492f422405/sha256_agg.control)

`sha256_agg` computes a single digest over ordered rows or selected values, for comparing logical table contents across copies.

### Core Workflow

```sql
CREATE EXTENSION sha256_agg;
CREATE TABLE digest_demo (id integer PRIMARY KEY, value text);
INSERT INTO digest_demo VALUES (1, 'one'), (2, 'two');
SELECT encode(sha256_agg(t ORDER BY id), 'hex') FROM digest_demo AS t;
SELECT sha256_agg_kernel();
```

### Operational Boundaries

Use a unique ordering key inside the aggregate. Empty input returns NULL; a row of NULL values still changes the digest. Comparisons require the same endianness. `sha256_agg_encode` exposes the hashed row bytes, and `sha256_agg_kernel` identifies the active implementation. `sha256_agg.hardware_acceleration` can disable acceleration without changing results. There is no partial/parallel aggregation. Upstream targets PostgreSQL 14+, with tests reported on 18 and master; the control is trusted and relocatable.
