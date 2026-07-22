## Usage

Sources:

- [pg_healpix README at the reviewed commit](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/README.md)
- [Version 1.0 SQL interface](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/pg_healpix--1.0.sql)
- [Coordinate conversion implementation](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/pg_healpix.c)

`pg_healpix` exposes four C functions for converting astronomical right ascension and declination to HEALPix pixel identifiers and back. It supports both nested and ring pixel ordering. Inputs and returned coordinates are degrees; the reverse functions return a two-element array ordered as right ascension and declination.

### Core Workflow

```sql
CREATE EXTENSION pg_healpix;

SELECT healpix_ang2ipix_nest(1024, 12.0, 25.0);
SELECT healpix_ang2ipix_ring(1024, 12.0, 25.0);

SELECT healpix_ipix2ang_nest(1024, 4323);
SELECT healpix_ipix2ang_ring(1024, 4323);
```

Use the inverse function with the same ordering scheme as the forward conversion. The returned coordinates represent the center of the selected HEALPix pixel, so a coordinate-to-pixel-to-coordinate round trip is a quantization operation rather than an exact restoration of the original point.

### Function Index

- `healpix_ang2ipix_nest(bigint, double precision, double precision)` returns a nested `ipix`.
- `healpix_ang2ipix_ring(bigint, double precision, double precision)` returns a ring `ipix`.
- `healpix_ipix2ang_nest(bigint, bigint)` returns a right-ascension/declination array for a nested `ipix`.
- `healpix_ipix2ang_ring(bigint, bigint)` returns a right-ascension/declination array for a ring `ipix`.

All four functions are declared `IMMUTABLE`, `PARALLEL SAFE`, and `STRICT`, making them eligible for expression indexes, generated expressions, and parallel plans when PostgreSQL's normal rules permit.

### Input Constraints

- `nside` must be a positive power of two no greater than `8192`.
- `ipix` must be at least zero and less than `12 * nside * nside`.
- Declination must map to a polar angle between zero and pi, which corresponds to the conventional range from `-90` through `90` degrees. Right ascension is converted from degrees and may wrap beyond one revolution.
- Non-finite angular inputs return `NULL`; SQL null inputs also return `NULL` because the functions are strict.

Catalog and control identify version `1.0`. Upstream publishes no current PostgreSQL compatibility matrix, and the catalog only records older server majors; test compilation, numerical agreement, index behavior, and representative boundary values on the exact server major and HEALPix tooling used by the application.
