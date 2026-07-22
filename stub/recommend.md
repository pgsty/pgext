## Usage

Sources:

- [Official README](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/README.md)
- [Official version 1.0 SQL](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/recommend--1.0.sql)
- [Official C implementation](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/recommend.c)
- [Official control file](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/recommend.control)

`recommend` is an experimental PostgreSQL 9.5-era extension with numeric distance, text character-vector similarity, and one-dimensional array similarity functions derived from `smlar`. It is not a complete recommendation system: callers must build candidate selection, ranking, persistence, and evaluation around these low-level scores.

### Core Workflow

After installing a build that has been reviewed and repaired as described below, the intended API is:

```sql
CREATE EXTENSION recommend;

SELECT twoint(10, 4);
SELECT two_float(1.5, 4.0);
SELECT arraysml(ARRAY[1, 2, 3], ARRAY[2, 3, 4]);
```

`twoint` and `two_float` return absolute difference. `arraysml` compares arrays of the same element type; its default bundled formula is cosine-style set similarity.

### API

- `twoint(integer, integer) RETURNS integer`: integer absolute difference.
- `two_float(double precision, double precision) RETURNS double precision`: floating-point absolute difference.
- `similarity(text, text) RETURNS real`: cosine similarity of a custom character-count vector.
- `arraysml(anyarray, anyarray) RETURNS real`: similarity for comparable one-dimensional arrays, with bundled `smlar` formula/configuration code.

### Caveats

The pinned official C source is not reproducibly buildable as published: `generate_vect` is missing its opening function brace. Its `similarity` implementation also computes an array index as each input byte minus lowercase `a` without bounds checks, so uppercase, punctuation, multibyte text, and other bytes can access memory outside the vector. Do not expose that function or claim a verified build without reviewing and patching the source.

The code targets old PostgreSQL internal APIs and has no maintained compatibility statement beyond the README's PostgreSQL 9.5 claim. `twoint` can overflow for extreme integers, empty text can divide by zero in cosine calculation, and array scoring depends on element comparison/canonicalization. Treat every result as an experimental score and validate a patched build with adversarial input before use.
