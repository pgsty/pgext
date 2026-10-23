## Usage

Sources:

- [Official source](https://github.com/soliyaanurag/Bloom-Filter-based-semi-join/blob/009b0dca08d240b773efeb92c2b57e3d4d8f88e7/README.md)
- [Official source](https://github.com/soliyaanurag/Bloom-Filter-based-semi-join/blob/009b0dca08d240b773efeb92c2b57e3d4d8f88e7/sjbloom/sjbloom.control)
- [Official source](https://github.com/soliyaanurag/Bloom-Filter-based-semi-join/blob/009b0dca08d240b773efeb92c2b57e3d4d8f88e7/sjbloom/sjbloom--1.0.sql)

`sjbloom` 1.0 builds Bloom, exact-key and pass-through filters for a research implementation of semi-join filtering. The supplied automatic planner integration patches `postgres_fdw` from PostgreSQL 18.6; installing this extension alone does not enable that planner behavior.

### Filter Functions

```sql
CREATE EXTENSION sjbloom;
SELECT sjb_contains(sjb_build(ARRAY[1,2,3], 1024, 3), 2);
SELECT sjb_info(sjb_build(ARRAY[1,2,3], 1024, 3));
```

`sjb_build` accepts integer keys, a bit count and a hash count. `sjb_contains` tests whether an integer passes the byte-array filter. Bloom matches may be false positives; the final exact join must remove them. NULL inputs remain NULL.

### Automatic Filtering and Limits

`sjb_build_auto` chooses a filter from cardinalities, row widths, network speed and calibrated costs. `sjb_choice_report` exposes the model's alternatives; `sjb_best_bits` estimates filter size. Default cost values came from upstream's test machine and are not universal performance guarantees.

For distributed use, install matching functions at both participating sites and follow the patched FDW setup. The standalone functions need no preload. The extension is relocatable; installation requires a superuser. This course prototype has no verified current-major portability or production-readiness claim, and the repository does not provide a verified license file.
