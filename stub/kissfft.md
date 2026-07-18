## Usage

Sources:

- [Project README](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/README.md)
- [Extension control file](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/kissfft.control)
- [SQL API](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/sql/kissfft.sql)
- [Native aggregate implementation](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/src/kissfft.c)
- [Regression example](https://github.com/pme/pgfft/blob/f52ac40b568eedfccfb547c8cd4bd7162cebd3f3/test/test.sql)

`kissfft` 0.0.1 is a 2012 C extension that wraps the Kiss FFT algorithm. It defines `fft_agg(real)` to collect an ordered scalar series and return a `real[]` power-like spectrum, plus `fft(real[], real)` to pair array positions with frequencies and `fft_print(real[])` to emit values as messages.

### Experimental query shape

Input order is part of the signal, so make it explicit inside the aggregate:

```sql
CREATE EXTENSION kissfft;

SELECT fft_agg(sample_value ORDER BY sampled_at)
FROM signal_sample;
```

The aggregate buffers the entire input as a PostgreSQL array and performs the transform only in the final function. Memory use grows with the group size, and large groups can exhaust a backend. Split and window data intentionally, and validate scaling, frequency bins, normalization, and real-signal symmetry against a trusted numerical library.

### Unsafe legacy implementation

The native final function deconstructs a PostgreSQL `Datum` array, then writes packed `float` values into that storage through a cast before reconstructing the SQL array. That assumes internal layouts that are unsafe on modern 64-bit PostgreSQL builds and can produce incorrect results or memory faults. The code also relies on assertions for input shape and has no current compatibility matrix.

Do not load this unmodified binary in a production backend. Audit and repair the C memory handling, add deterministic numerical tests and malformed-input tests, and run it under sanitizers on the exact PostgreSQL major. For ordinary analytics, transform exported data with a maintained numerical library instead of putting this abandoned ABI-sensitive implementation inside the database server.
