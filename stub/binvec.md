## Usage

Sources:

- [Upstream README](https://github.com/sany1231/binvec/blob/1bca9cab92fb199717c206e7d6601b82c727cf04/README.md)
- [Version 0.1 install SQL](https://github.com/sany1231/binvec/blob/1bca9cab92fb199717c206e7d6601b82c727cf04/binvec--0.1.sql)
- [C implementation](https://github.com/sany1231/binvec/blob/1bca9cab92fb199717c206e7d6601b82c727cf04/binvec.c)

binvec 0.1 supplies vec_sum_bin(anyarray, integer), a transition function that adds the set bits of an integer to a 32-element vector. The aggregate shown in the README is not installed by the extension and must be created separately if needed.

### Basic call

Use only an integer array in an isolated test database:

```sql
CREATE EXTENSION binvec;
SELECT vec_sum_bin(ARRAY[1, 0, 1]::integer[], 5);
```

The result has 32 elements. Bit zero maps to the first element, so the first and third elements are incremented for the value 5.

### Caveats

- The C code accepts several numeric array element types but reads and writes every element as a 32-bit integer. Arrays other than integer arrays can be corrupted or produce invalid results.
- Input array null flags are ignored, and multidimensional arrays are rejected. An empty array returns null, while a zero right-hand value returns the original array without expanding it to 32 elements. The SQL function is not strict, but the C code reads its second argument without checking for null; never pass a null right-hand value.
- The signed high bit and negative input values are not handled safely. Restrict the right-hand value to nonnegative values below 2^31.
- This is an old prototype with no tests, license, release history, or current PostgreSQL compatibility statement. Validate the exact build before any non-disposable use.
