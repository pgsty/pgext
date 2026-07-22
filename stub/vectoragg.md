## Usage

Sources:

- [Official version 1.0 SQL](https://github.com/cybertec-postgresql/vectoragg/blob/master/vectoragg--1.0.sql)
- [Official C implementation](https://github.com/cybertec-postgresql/vectoragg/blob/master/vectoragg.c)
- [Official control file](https://github.com/cybertec-postgresql/vectoragg/blob/master/vectoragg.control)
- [Official archived repository](https://github.com/cybertec-postgresql/vectoragg)

`vectoragg` provides C functions for summing or slicing uniformly spaced `real[]`/`double precision[]` samples over a coordinate interval, plus fixed-factor downsampling helpers. It is an archived, minimally documented version `1.0` project; use only after reviewing the source-level defects below.

### Core Workflow

For a four-sample signal spanning coordinates `0` through `3`, sum or extract the half-open coordinate interval from `1` to `3`:

```sql
CREATE EXTENSION vectoragg;

SELECT array_sum(
  ARRAY[1, 2, 3, 4]::double precision[],
  0, 3, 1, 3
);

SELECT array_clamp(
  ARRAY[1, 2, 3, 4]::double precision[],
  0, 3, 1, 3
);
```

The functions infer sample spacing as `(tend - tstart) / (array_length - 1)`. In this example the sum is `5` and the slice is `{2,3}`.

### API

Both `real[]` and `double precision[]` overloads are installed:

- `array_sum(samples, tstart, tend, astart, aend) RETURNS double precision`: sum selected samples.
- `array_clamp(samples, tstart, tend, astart, aend) RETURNS array`: return the selected slice.
- `array_decimate(samples) RETURNS array`: intended to average blocks of 10.
- `array_hundreth(samples) RETURNS array`: intended to average blocks of 100; the misspelling is part of the SQL API.

### Caveats

Do not call the published `array_decimate` or `array_hundreth` on arbitrary lengths. In the official C source, each residual-element `while` loop never increments its index, so arrays whose length is not an exact multiple of 10 or 100 can loop forever in a PostgreSQL backend. These functions need a source patch and regression tests before use.

`array_sum`/`array_clamp` reject arrays containing nulls but do not validate empty/single-element arrays or coordinate ordering before dividing by `n - 1`. Confirm `n >= 2`, finite ordered bounds, and expected half-open rounding at sample boundaries. The repository is archived and its README is empty, so there is no upstream maintenance or compatibility guarantee.
