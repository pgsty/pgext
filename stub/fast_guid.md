## Usage

Sources:

- [Official README](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/README.rst)
- [Extension control file](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/fast_guid.control)
- [Version 0.1 extension SQL](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/fast_guid--0.1.sql)
- [C implementation](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/fast_guid.c)

`fast_guid` 0.1 returns a time-, shard-, and sequence-derived `bigint`. It does not return PostgreSQL `uuid` values and does not allocate sequence numbers itself; the caller is responsible for supplying collision-free `smallint` inputs.

### Core Workflow

Use a sequence to supply the low-order value:

```sql
CREATE EXTENSION fast_guid;
CREATE SEQUENCE global_id_sequence;

SELECT fast_guid(
  7::smallint,
  (nextval('global_id_sequence') % 1024)::smallint
);
```

The function reads wall-clock time in milliseconds since `2011-01-01 00:00:00 UTC`, reduces `shard_id` modulo 8192, and reduces `seq_id` modulo 1024 before combining them into a signed `bigint`.

### Interface

`fast_guid(shard_id smallint, seq_id smallint) RETURNS bigint` is `STRICT` and otherwise uses PostgreSQL's default `VOLATILE` classification. NULL input therefore yields NULL without calling the C function.

The upstream comments describe an intended 41-bit time, 13-bit shard, and 10-bit sequence layout. The actual source shifts the shard value by 13 bits, not 10, so consumers must not decode IDs using the documented contiguous 41/13/10 layout without first validating the implementation they installed.

### Collision and Ordering Caveats

Two calls in the same millisecond with the same reduced shard and sequence inputs can return the same value. More than 1024 IDs per shard per millisecond, sequence resets, clock rollback, duplicated shard assignments, or negative `smallint` inputs can all undermine uniqueness or ordering. The function has no coordination across nodes and no persistent clock state.

Use a primary-key uniqueness constraint even if the function is retained, and do not treat its output as a cryptographic, random, globally coordinated, or opaque identifier. For new systems, prefer a maintained identifier generator whose bit layout, overflow behavior, and distributed-node allocation are explicitly specified and tested.
