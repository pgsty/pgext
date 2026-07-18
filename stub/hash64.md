## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/theory/pg-hash64/blob/2f969d14998f82c98322bb6f53f788b5be5d8b18/README.md)
- [Extension control file](https://github.com/theory/pg-hash64/blob/2f969d14998f82c98322bb6f53f788b5be5d8b18/hash64.control)
- [C implementation](https://github.com/theory/pg-hash64/blob/2f969d14998f82c98322bb6f53f788b5be5d8b18/hash64.c)

`hash64` exposes `hash64(text)`, a stable 64-bit hash returned as `bigint`. Unlike PostgreSQL's internal hash functions, upstream promises that its result will remain unchanged across PostgreSQL major releases, making it useful for deterministic sharding or partition assignment in mixed-version deployments.

```sql
CREATE EXTENSION hash64;
SELECT hash64('customer-42');
SELECT mod(hash64('customer-42') & 9223372036854775807, 32) AS shard_id;
```

A hash is not unique and is not a cryptographic authenticator. Always retain the original key for equality checks, allow for collisions, and do not use the result for password storage, signatures, or adversarial integrity checks.

Version 0.1.0 is old and has no current PostgreSQL compatibility matrix. Before using it as a cross-cluster routing contract, build it on every target architecture and verify a fixed corpus of expected outputs; changing integer normalization or platform behavior would otherwise move data between shards.
