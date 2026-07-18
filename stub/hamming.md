## Usage

Sources:

- [Upstream README](https://github.com/AmineDiro/pghamming/blob/27b5119efe65f6d7f2173401ea8d20beb2d31817/README.md)
- [Distance function and operator](https://github.com/AmineDiro/pghamming/blob/27b5119efe65f6d7f2173401ea8d20beb2d31817/src/hamming.rs)
- [Index build implementation](https://github.com/AmineDiro/pghamming/blob/27b5119efe65f6d7f2173401ea8d20beb2d31817/src/index/build.rs)

hamming 0.1.0 provides hamming_distance(bytea, bytea) and the <#> operator for byte-level Hamming comparisons.

### Sequential comparison

```sql
CREATE EXTENSION hamming;
SELECT hamming_distance(decode('0f', 'hex'), decode('f0', 'hex'));
SELECT decode('0f', 'hex') <#> decode('f0', 'hex');
```

Both calls return 8 for these one-byte inputs. The stable useful surface is the function and operator evaluated by a normal scan.

### Caveats

- Do not create an ivf index. The access method is a nonfunctional prototype: its build callback does not index tuples, several required callbacks abort when called, and insertion support is absent.
- The README's index example therefore does not provide approximate-nearest-neighbour search or durable index storage.
- The control file marks the extension as untrusted and requiring superuser installation. Limit installation and upgrades to reviewed builds.
- The Rust package targets pgrx 0.10.2 and advertises PostgreSQL 11 through 16 only. There is no compatibility evidence for newer majors and no published license.
