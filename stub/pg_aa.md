## Usage

Sources:

- [Official extension SQL](https://github.com/akorotkov/pg_aa/blob/4ae0a182dfd66e65e18ef8faf460973c8706ac44/pg_aa--1.0.sql)
- [Official C implementation](https://github.com/akorotkov/pg_aa/blob/4ae0a182dfd66e65e18ef8faf460973c8706ac44/pg_aa.c)
- [Official extension control file](https://github.com/akorotkov/pg_aa/blob/4ae0a182dfd66e65e18ef8faf460973c8706ac44/pg_aa.control)

`pg_aa` converts a PNG image stored as `bytea` into text art inside PostgreSQL. It exposes one monochrome ASCII renderer backed by aalib and one UTF-8 renderer backed by libcaca; it is a demonstration extension, not an image-storage or document-processing system.

### Core Workflow

Store the raw PNG bytes in a table or bind them as a query parameter, then choose an output width:

```sql
CREATE EXTENSION pg_aa;

CREATE TABLE images (
  id bigint PRIMARY KEY,
  png bytea NOT NULL
);

SELECT aa_out(png, 80)
FROM images
WHERE id = 1;

SELECT caca_out(png, 80)
FROM images
WHERE id = 1;
```

The output height is derived from the source aspect ratio and the requested width, with an adjustment for terminal character proportions.

### Function Index

- `aa_out(bytea, int4)` returns newline-separated, locale-independent ASCII using aalib.
- `caca_out(bytea, int4)` returns libcaca's UTF-8 text rendering.

Both functions are immutable and strict. Input must be a decodable PNG; other formats raise an error.

### Operational Boundaries

The server binary must link compatible versions of libgd, aalib, and libcaca. Decode and raster processing happen synchronously in the database backend, so restrict image size and output width to prevent excessive CPU or memory consumption. The reviewed source dates from 2015 and has no README, declared compatibility matrix, or size guards; test malformed PNGs, zero/negative widths, extreme aspect ratios, Unicode client encodings, library upgrades, and parallel use before exposing either function to untrusted callers.
