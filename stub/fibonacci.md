## Usage

Sources:

- [Upstream README](https://github.com/japinli/fibonacci/blob/aace7f2ec1deb3e32659fa79d9615549e7598790/README.md)
- [Version 0.0.2 install SQL](https://github.com/japinli/fibonacci/blob/aace7f2ec1deb3e32659fa79d9615549e7598790/fibonacci--0.0.2.sql)
- [C implementation](https://github.com/japinli/fibonacci/blob/aace7f2ec1deb3e32659fa79d9615549e7598790/fibonacci.c)

fibonacci 0.0.2 provides fibonacci(integer), an iterative C implementation that returns a 32-bit integer.

### Example

```sql
CREATE EXTENSION fibonacci;
SELECT n, fibonacci(n) FROM generate_series(0, 10) AS n;
```

For exact nonnegative 32-bit results, restrict input to the inclusive range from 0 through 46.

### Caveats

- Negative inputs silently return zero.
- Fibonacci number 47 exceeds the signed 32-bit return type. Larger inputs trigger signed overflow and can return invalid results.
- Runtime grows linearly with the input. A very large positive integer can keep a backend CPU busy for a long time, so do not expose the function to unbounded input.
- Upstream provides no license, current compatibility matrix, or meaningful test suite. Treat it as an educational example rather than a production numeric implementation.
