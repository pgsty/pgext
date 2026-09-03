## Usage

Sources:

- [Official documentation](https://github.com/thehyve/pg_bitcount/blob/77d4fa8e26dd46166a9d906920e8b341fb8fb643/README.md)
- [Extension control file](https://github.com/thehyve/pg_bitcount/blob/77d4fa8e26dd46166a9d906920e8b341fb8fb643/pg_bitcount.control)
- [Official repository](https://github.com/thehyve/pg_bitcount)

`pg_bitcount` Bit-count scalar and aggregate functions for integers and bit strings.

### Enablement

Install the files for the intended server, then create `pg_bitcount` in the target database:

```sql
CREATE EXTENSION pg_bitcount;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Register the extension in PostgreSQL
create extension pg_bitcount version '0.0.3';

-- Use the pg_bitcount function
select public.pg_bitcount(127::bit(8)); -- 7
select public.pg_bitcount(B'101010101'); -- 5
select public.pg_bitcount((17^15)::bigint::bit(128) << 64 | (17^14)::bigint::bit(128)); -- 58

-- Use the pg_int_to_bit_agg aggregate using a bit string of size 24
select public.pg_bitcount(public.pg_int_to_bit_agg(i::int, 24)) from (select generate_series(2, 8) as i) data; -- 7
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `public.pg_bitcount` | FUNCTION | Callable function from the reviewed install surface. |
| `public.pg_int_to_bit_agg` | AGGREGATE | Aggregate exposed by the extension. |
| `public.pg_int_to_bit_agg_transfn` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 11, 12, 13; do not infer unlisted majors.
- Catalog lifecycle is archived; test upgrades, dump/restore, and server compatibility before production use.
