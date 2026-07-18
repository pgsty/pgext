## Usage

Sources:

- [README at the reviewed revision](https://gitlab.com/nfiesta/nfiesta_pg/-/blob/5bf6158866a327db97a1c8513f58a06228573ecb/README.md)
- [nFIESTA project wiki](https://gitlab.com/nfiesta/nfiesta_pg/-/wikis/home)
- [SQL source tree](https://gitlab.com/nfiesta/nfiesta_pg/-/tree/5bf6158866a327db97a1c8513f58a06228573ecb/functions)

`nfiesta` is the PostgreSQL estimation engine for nFIESTA, a forest-inventory estimation and analysis system. It combines national forest inventory field samples with auxiliary data, commonly remote-sensing products, and implements single- and two-phase estimation, configuration, ETL, result, and additivity workflows.

```sql
CREATE EXTENSION nfiesta;
SELECT e.extname, e.extversion
FROM pg_extension AS e
WHERE e.extname = 'nfiesta';
```

The extension installs a large domain schema rather than one self-contained function. Follow the upstream installation guide and Minimal Working Example for the exact release, load a documented test data set, configure an estimation period and model, then compare generated totals, ratios, variances, and confidence results with independently calculated reference values.

Statistical correctness depends on sampling design, strata, weights, auxiliary variables, missing-data policy, model configuration, and release-specific SQL. Preserve input lineage and configuration, isolate ETL and GUI roles, review security-definer and dynamic-SQL functions, and lock the extension version for reproducible reports. Test migrations and rollback on a copy because upgrades can alter many functions, views, configuration tables, and result generations.
