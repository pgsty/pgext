## Usage

Sources:

- [Official README](https://gitlab.com/nfiesta/nfiesta_htc/-/blob/267e25465ee4ea073fdbfbdddf1bf8d7d9c6359c/README.md)
- [Version 2.1 update SQL](https://gitlab.com/nfiesta/nfiesta_htc/-/blob/267e25465ee4ea073fdbfbdddf1bf8d7d9c6359c/htc--2.0--2.1.sql)
- [C estimator interface](https://gitlab.com/nfiesta/nfiesta_htc/-/blob/267e25465ee4ea073fdbfbdddf1bf8d7d9c6359c/functions/htc_compute.sql)
- [Regression workflow](https://gitlab.com/nfiesta/nfiesta_htc/-/blob/267e25465ee4ea073fdbfbdddf1bf8d7d9c6359c/sql/htc.sql)

`htc` implements Horvitz-Thompson estimators for continuous-population sampling. It targets specialized survey and forest-inventory calculations, accepting sampled cluster, density, weight, area, inclusion-probability, and auxiliary-variable arrays rather than ordinary row aggregates.

### Core Workflow

Install the procedural and numerical prerequisites used by the upstream version before creating `htc`:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION plpython3u;
CREATE EXTENSION htc;

SELECT htc_version(), htc_version_ok();
```

For the original one-phase interface, aggregate one stratum into aligned arrays and pass the stratum constants explicitly:

```sql
SELECT (htc_compute(
    ARRAY[1, 2, 3, 4],
    ARRAY[1.1, 2.2, 3.3, 4.4]::double precision[],
    ARRAY[0.75, 0.65, 0.55, 0.45]::double precision[],
    123.456,
    75.0,
    100
)).*;
```

Use the upstream regression workflow as the canonical guide for preparing those arrays and interpreting the returned `total` and `var` values; estimator validity depends on preserving the required ordering, shapes, units, and sampling-design assumptions.

### Object Index

- `htc_compute(...)` returns one-phase `total` and `var` estimates from C.
- `htc_compute_sweight_ha(...)` is the sampling-weight-per-hectare variant.
- `htc_gbeta(...)` builds generalized regression coefficients from JSONB matrices.
- `htc_total2p(...)` returns a two-phase total, total variance, and variance components for each phase.
- `htc_variance_phase1(...)` and `htc_variance_phase2(...)` are two-phase variance helpers.
- `htc_version()` and `htc_version_ok()` check that the loaded shared library matches the SQL release.

### Operational Notes

Version `2.1` is relocatable and has no preload requirement. Its newer two-phase routines are `plpython3u` functions importing NumPy and SciPy, so those Python modules must be available to PostgreSQL's server-side Python. `plpython3u` is untrusted and normally requires superuser-level control to install; restrict who can replace or invoke procedural code.

The JSONB routines validate array counts, dimensions, nulls, and forbidden zero probabilities, but that does not validate the statistical design. Treat the upstream regression inputs as schemas, verify `htc_version_ok()` after binary or SQL upgrades, and independently validate estimator results before operational decisions.
