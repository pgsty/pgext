## Usage

Sources:

- [Official README](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/README.md)
- [Extension SQL for version 1.0.0](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/shbf--1.0.0.sql)
- [C implementation](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/shbf.c)
- [Original Shifting Bloom Filter paper](https://www.vldb.org/pvldb/vol9/p408-yang.pdf)

`shbf` implements Count-Min Sketch, ordinary Bloom filter, and three Shifting Bloom Filter variants as compact probabilistic PostgreSQL values. These structures trade exact answers for bounded memory: positive membership and association results can be false positives, and frequency estimates can be overestimates.

### Core Workflow

Create a filter with a planned bit count `m` and expected element count `n`. Insert functions return a new value, so always assign the result; they do not persistently mutate the stored datum in place.

```sql
CREATE EXTENSION shbf;

CREATE TABLE seen_key (
  id integer PRIMARY KEY,
  filter bf NOT NULL
);

INSERT INTO seen_key VALUES (1, new_bf(1000000, 10000));

UPDATE seen_key
SET filter = insert_bf(filter, 'alpha')
WHERE id = 1;

SELECT query_bf(filter, 'alpha')
FROM seen_key
WHERE id = 1;
```

Use a Count-Min Sketch for approximate frequencies.

```sql
CREATE TEMP TABLE frequency(sketch cms);
INSERT INTO frequency VALUES (new_cms(0.001, 0.99));

UPDATE frequency SET sketch = insert_cms(sketch, 'alpha');
UPDATE frequency SET sketch = insert_cms(sketch, 'alpha');

SELECT query_cms(sketch, 'alpha') FROM frequency;
```

### Important Objects

- `cms`: Count-Min Sketch; `new_cms(error_bound, confidence)` defaults to `0.001` and `0.99`, with `insert_cms` and `query_cms` for approximate counts.
- `bf`: ordinary Bloom filter; `new_bf(m, n)`, `insert_bf`, and `query_bf` provide membership operations.
- `new_shbf_m(m, n)`, `insert_shbf_m`, `query_shbf_m`: shifting membership filter.
- `new_shbf_a(m, n)`, `insert_shbf_a`, `query_shbf_a`: association filter for two sets.
- `new_shbf_x(m, n, max_x)`, `insert_shbf_x`, `query_shbf_x`: multiplicity filter up to a planned maximum.

`query_shbf_a` returns encoded states: `-1` absent; `0` second set; `1` first set; `2` both; `3` first-or-second; `4` both-or-second; `5` both-or-first; `6` any of the three membership states. Ambiguous codes are a consequence of collisions.

### Sizing and Safety

Choose positive `m`, `n`, and `max_x` values from a measured workload, and keep inserted multiplicities within the planned bound. The reviewed constructors do little validation: zero or invalid sizes can reach division, modulo, allocation, or offset arithmetic in C. Test sizing and saturation outside production before storing a filter.

These types provide no deletion, merge aggregate, index operator class, automatic resize, or exact confidence report. Keep authoritative data elsewhere and treat the filter as an acceleration or telemetry summary only. The upstream project is compact research-oriented code; validate serialization across architectures and PostgreSQL upgrades before using stored values as durable data.
