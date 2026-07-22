## Usage

Sources:

- [Official README](https://github.com/ferndale-hall/pg_gsl/blob/18630157063cd7a71e5ea6b3a3a3e84bfe9f19f3/README.md)
- [Extension SQL definitions](https://github.com/ferndale-hall/pg_gsl/blob/18630157063cd7a71e5ea6b3a3a3e84bfe9f19f3/sql/pg_gsl.sql)
- [PGXN metadata](https://github.com/ferndale-hall/pg_gsl/blob/18630157063cd7a71e5ea6b3a3a3e84bfe9f19f3/META.json)

`pg_gsl` 0.0.2 exposes a small GNU Scientific Library (GSL) FFT surface to PostgreSQL. Use it to transform a one-dimensional `double precision[]`, inspect or truncate its frequency spectrum, and reconstruct the signal inside SQL. It wraps only two GSL routines; it is not a general GSL binding.

### Core Workflow

Install the GSL shared library where PostgreSQL can load it, install the extension files, and create the extension in the target database:

```sql
CREATE EXTENSION pg_gsl;
SELECT pg_gsl_x_version();
```

Transform a signal, derive its normalized spectrum, zero the final transform entries, and invert the result:

```sql
WITH signal(x) AS (
  VALUES (ARRAY[1.0, 0.0, -1.0, 0.0]::double precision[])
), transformed AS (
  SELECT pg_gsl_fft_real_transform(x) AS fft FROM signal
)
SELECT pg_gsl_x_fftToSpectrum(fft) AS spectrum,
       pg_gsl_fft_halfcomplex_inverse(
         pg_gsl_x_fftTruncate(fft, 2)
       ) AS filtered_signal
FROM transformed;
```

The forward transform returns GSL's half-complex representation in an array. Pass that representation directly to the inverse function; `pg_gsl_x_fftToSpectrum` is for inspection, not for inversion.

### Function Index

- `pg_gsl_fft_real_transform(double precision[])` performs the real FFT.
- `pg_gsl_fft_halfcomplex_inverse(double precision[])` reconstructs a signal from the half-complex array.
- `pg_gsl_x_fftToSpectrum(double precision[])` converts a transform into a normalized spectrum of half the input length.
- `pg_gsl_x_fftTruncate(double precision[], integer)` replaces the final `n` array entries with zero, suppressing higher-frequency components in the representation used by this project.
- `pg_gsl_x_version()` returns the extension version.

### Requirements and Caveats

`pg_gsl` dynamically links to GSL. The upstream build expects `libgsl.so` to be available in PostgreSQL's `pkglibdir`; no `shared_preload_libraries` setting is required. Array shape and length are passed to low-level FFT code, so validate inputs in application code and test representative signals before production use. Version 0.0.2 is an early, narrowly scoped release from 2016.
