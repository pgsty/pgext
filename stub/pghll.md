## Usage

Sources:

- [Extension control file](https://github.com/baverman/pghll/blob/692b1daa98c3a6b721acf850a90f6029e42e84ca/pghll.control)
- [Version 1.0 SQL API](https://github.com/baverman/pghll/blob/692b1daa98c3a6b721acf850a90f6029e42e84ca/pghll--1.0.sql)
- [HLL decoding and aggregation implementation](https://github.com/baverman/pghll/blob/692b1daa98c3a6b721acf850a90f6029e42e84ca/pghll.c)

`pghll` 1.0 reads a specific serialized HyperLogLog counter format from `bytea`. It can decompress counters, estimate one counter's cardinality, merge decoded counters, or merge compressed counters and return their combined estimate.

### Aggregate compatible counters

Assuming `hll_samples.counter` contains counters produced in the format expected by this extension:

```sql
CREATE EXTENSION pghll;

SELECT hll_count(hll_merge(hll_decode(counter)))
FROM hll_samples;

SELECT hll_sum(counter)
FROM hll_samples;
```

`hll_decode` expands a zlib-compressed, network-byte-order payload. The `hll_merge` aggregate accepts decoded counters and returns a decoded counter; `hll_count` estimates its cardinality. The `hll_sum` aggregate combines compressed counters directly and returns the estimate.

### Input trust and compatibility

The extension does not provide a constructor, hashing function, format identifier, or payload validator. Its C code trusts header values inside decoded `bytea` when choosing register counts and memory offsets. Pass only counters from a known compatible producer; arbitrary or truncated input can raise an error and may reach unsafe native-code paths.

All counters in one aggregate must use the same precision and serialization layout. The format is not the same thing as the data types supplied by other PostgreSQL HLL extensions merely because they share the HLL name. This repository has no README, release artifacts, upgrade path, or current PostgreSQL compatibility matrix and is cataloged as abandoned; isolate it from untrusted input and validate it on the exact server build before reading historical data.
