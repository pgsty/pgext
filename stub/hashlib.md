


## Usage

> [hashlib: Stable hash functions library for PostgreSQL](https://github.com/markokr/pghashlib)

Provides stable hash functions whose implementations do not change across PostgreSQL versions.

### String Hashing (32-bit)

```sql
SELECT hash_string('hello', 'crc32');
SELECT hash_string('hello', 'murmur3');
```

With optional initial value:

```sql
SELECT hash_string('hello', 'crc32', 42);
```

### String Hashing (64-bit)

```sql
SELECT hash64_string('hello', 'city64');
SELECT hash64_string('hello', 'siphash24');
SELECT hash64_string('hello', 'lookup3');
```

### String Hashing (128-bit)

```sql
SELECT hash128_string('hello', 'md5');
SELECT hash128_string('hello', 'city128');
SELECT hash128_string('hello', 'spooky');
```

### Integer Hashing

```sql
SELECT hash_int4(42);        -- 32-bit hash of 32-bit integer
SELECT hash_int8(42::bigint); -- 64-bit hash of 64-bit integer
```

### Available Algorithms

| Algorithm | CPU-indep | Bits | Description |
|-----------|-----------|------|-------------|
| `crc32` | yes | 32 | CRC32 |
| `murmur3` | no | 32 | MurmurHash v3 |
| `md5` | yes | 128 | MD5 |
| `city64` | no | 64 | CityHash64 |
| `city128` | no | 128 | CityHash128 |
| `siphash24` | yes | 64 | SipHash-2-4 |
| `spooky` | no | 128 | SpookyHash |
| `lookup2` | no | 64 | Jenkins lookup2 |
| `lookup3` | no | 64 | Jenkins lookup3 CPU-native |
| `lookup3be` | yes | 64 | Jenkins lookup3 big-endian |
| `lookup3le` | yes | 64 | Jenkins lookup3 little-endian |
| `pgsql84` | no | 64 | Hacked lookup3 in Postgres 8.4+ |

Integer algorithms: `wang32`, `wang32mult`, `jenkins` (32-bit); `wang64`, `wang64to32` (64-bit). All are reversible (1:1 mapping), useful for creating random sort orders over unique IDs.
