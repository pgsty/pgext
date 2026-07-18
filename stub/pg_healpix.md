## Usage

Sources:

- [pg_healpix README at the reviewed commit](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/README.md)
- [pg_healpix.control at the reviewed commit](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/pg_healpix.control)
- [Version 1.0 SQL interface](https://github.com/segasai/pg_healpix/blob/93220a4588974d9942d5bca7cf1e58048a441a04/pg_healpix--1.0.sql)

`pg_healpix` exposes HEALPix sky-pixel conversions as PostgreSQL functions. `healpix_ang2ipix_nest` and `healpix_ang2ipix_ring` convert an `nside`, right ascension, and declination into a nested or ring `ipix`. `healpix_ipix2ang_nest` and `healpix_ipix2ang_ring` perform the reverse conversion and return a two-element `double precision[]`.

### Basic Conversions

```sql
CREATE EXTENSION pg_healpix;

SELECT healpix_ang2ipix_nest(1024, 12.0, 25.0);
SELECT healpix_ang2ipix_ring(1024, 12.0, 25.0);

SELECT healpix_ipix2ang_nest(1024, 4323);
SELECT healpix_ipix2ang_ring(1024, 4323);
```

The functions are declared `IMMUTABLE`, `PARALLEL SAFE`, and `STRICT`, so they can be used in generated columns, expression indexes, and parallel queries when their input conventions fit the workload.

### Caveats

- Keep nested and ring identifiers separate; the inverse function must use the same ordering convention as the forward conversion.
- Confirm the upstream right-ascension and declination convention and valid `nside` range before importing external astronomy data.
- Control and catalog both identify version `1.0`. The repository provides no current PostgreSQL compatibility matrix, so test on the exact server major version you intend to run.
