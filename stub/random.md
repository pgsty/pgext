

## Usage

> [random: reproducible random data generators for PostgreSQL](https://github.com/tvondra/random)

Provides functions to generate random values for various data types with reproducible output controlled by a seed.

```sql
CREATE EXTENSION random;
```

### Functions

All functions accept `seed` (for reproducibility) and `nvalues` (number of distinct values).

| Function | Description |
|---|---|
| `random_string(seed, nvalues, min_length, max_length)` | Random ASCII string |
| `random_bytea(seed, nvalues, min_length, max_length)` | Random bytea |
| `random_int(seed, nvalues, min_value, max_value)` | Random 32-bit integer |
| `random_bigint(seed, nvalues, min_value, max_value)` | Random 64-bit integer |
| `random_real(seed, nvalues, min_value, max_value)` | Random 32-bit float |
| `random_double_precision(seed, nvalues, min_value, max_value)` | Random 64-bit float |
| `random_inet(seed, nvalues)` | Random INET address (/32 mask) |
| `random_cnet(seed, nvalues)` | Random CIDR with masks 8/16/24/32 |
| `random_cnet2(seed, nvalues)` | Random CIDR with equal fraction per mask length |
| `random_macaddr(seed, nvalues)` | Random 6-byte MAC address |
| `random_macaddr8(seed, nvalues)` | Random 8-byte MAC address |

### Examples

```sql
-- Generate reproducible random integers
SELECT random_int(42, 100, 1, 1000) FROM generate_series(1, 10);

-- Random strings of length 5-10
SELECT random_string(42, 1000, 5, 10) FROM generate_series(1, 5);

-- Random IP addresses
SELECT random_inet(42, 256) FROM generate_series(1, 5);
```
