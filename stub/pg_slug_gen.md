## Usage

Sources: [official PGXN release page](https://pgxn.org/dist/pg_slug_gen/), [official release README](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/README.md), [official release SQL](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/sql/pg_slug_gen--1.0.sql), [official release metadata](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/META.json)

`pg_slug_gen` generates timestamp-based slugs with cryptographic randomness. The official 1.0.0 release describes it as a secure, URL-friendly short ID generator where the requested length selects the timestamp precision.

```sql
CREATE EXTENSION pg_slug_gen;

SELECT gen_random_slug();
SELECT gen_random_slug(13);
```

### Function

- `gen_random_slug(slug_length int DEFAULT 16) returns text`

The release SQL comment and README document these supported values:

- `10`: seconds
- `13`: milliseconds
- `16`: microseconds, also the default
- `19`: nanoseconds

### Precision And Format

Each precision maps to a timestamp width and a fixed slug shape:

- `10` digits: `5-5` format, 11 characters total
- `13` digits: `6-7` format, 14 characters total
- `16` digits: `8-8` format, 17 characters total
- `19` digits: `9-10` format, 20 characters total

The README states the collision-free window is bounded by timestamp precision: at most 1 insert per second, millisecond, microsecond, or nanosecond respectively.

### Examples

```sql
SELECT gen_random_slug();
SELECT gen_random_slug(10);
SELECT gen_random_slug(16);

CREATE TABLE products (
  id serial PRIMARY KEY,
  name text NOT NULL,
  slug text DEFAULT gen_random_slug() UNIQUE
);
```

### How It Works

The official README describes this algorithm:

- take the current timestamp at the chosen precision
- map each digit to a QWERTY-based character bucket
- choose one random character from that bucket with `pg_strong_random()`
- insert a hyphen at the midpoint

### Caveats

- This is a secure short-ID generator, not a text transliteration or title-to-URL slugifier.
- Same-timestamp collisions are still possible; upstream only claims uniqueness when inserts do not exceed one per chosen time unit.
- The official release metadata still points to `https://github.com/fernandoolle/pg_slug_gen`, but that repo URL currently returns 404.
