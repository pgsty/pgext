## Usage

Sources:

- [Pinned upstream README](https://github.com/PixNyanNyan/postgres-phash-hamming/blob/4b9b69c41acd75be1fe4e6d91f1117622d2c77f2/README.md)
- [Version 1.0 installation SQL](https://github.com/PixNyanNyan/postgres-phash-hamming/blob/4b9b69c41acd75be1fe4e6d91f1117622d2c77f2/pg_phash--1.0.sql)
- [Pinned C implementation](https://github.com/PixNyanNyan/postgres-phash-hamming/blob/4b9b69c41acd75be1fe4e6d91f1117622d2c77f2/pg_phash.c)
- [Pinned extension control file](https://github.com/PixNyanNyan/postgres-phash-hamming/blob/4b9b69c41acd75be1fe4e6d91f1117622d2c77f2/pg_phash.control)

pg_phash version 1.0 compares two 64-bit perceptual-image hashes represented as base-10 strings. It XORs the parsed integers and returns the number of differing bits.

### Example

```sql
CREATE EXTENSION pg_phash;

SELECT phash_hamming(
    '13121266429874464083',
    '10869537466045227287'
);
```

The upstream example returns 16. Identical valid hashes return zero. The SQL function is strict, immutable, and parallel safe, so a null input produces null and constant calls may be folded.

### Parsing and range limits

Despite the README's stated 0 through 255 range, the implementation compares one 64-bit value and can return only 0 through 64. Inputs are converted with the C strtoull function, but the code does not check the conversion end pointer or errno. Empty, negative, malformed, trailing-character, and overflow inputs can therefore be silently interpreted as zero, a prefix value, or the maximum unsigned integer rather than rejected.

Input length is also stored in an unsigned one-byte variable before allocation, so strings longer than 255 bytes can wrap and be truncated. Accept only canonical decimal strings in the unsigned 64-bit range, validate them before calling, and add regression tests for zero, the maximum value, malformed text, and overflow.

The project documents a PostgreSQL 9.6 CI example and has not changed since 2021. It has no release matrix for current PostgreSQL majors. The repository is not archived, but this small C function should still be rebuilt and tested against the exact server version.
