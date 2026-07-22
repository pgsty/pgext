## Usage

Sources:

- [README at the reviewed revision](https://gitlab.com/nfiesta/nfiesta_pg/-/blob/5bf6158866a327db97a1c8513f58a06228573ecb/README.md)
- [nFIESTA project wiki](https://gitlab.com/nfiesta/nfiesta_pg/-/wikis/home)
- [SQL source tree](https://gitlab.com/nfiesta/nfiesta_pg/-/tree/5bf6158866a327db97a1c8513f58a06228573ecb/functions)
- [Extension control file](https://gitlab.com/nfiesta/nfiesta_pg/-/blob/5bf6158866a327db97a1c8513f58a06228573ecb/nfiesta.control)

`nfiesta` is the PostgreSQL estimation engine for nFIESTA, a forest-inventory estimation and analysis system. It combines national forest inventory field samples with auxiliary data, commonly remote-sensing products, and implements single- and two-phase estimation, configuration, ETL, result, and additivity workflows.

The control file declares `postgis`, `plpython3u`, `htc`, and `nfiesta_sdesign` as required extensions. In particular, `plpython3u` is an untrusted procedural language and creates a server-operating-system trust boundary even though `nfiesta` itself is marked `superuser = false`.

```sql
CREATE EXTENSION nfiesta CASCADE;
SELECT e.extname, e.extversion
FROM pg_extension AS e
WHERE e.extname IN (
  'postgis', 'plpython3u', 'htc', 'nfiesta_sdesign', 'nfiesta'
);
```

The extension installs a large domain schema rather than one self-contained function. Follow the upstream installation guide and Minimal Working Example for the exact release, load a documented test data set, configure an estimation period and model, then compare generated totals, ratios, variances, and confidence results with independently calculated reference values.

`nfiesta.fn_make_estimate(integer)` materializes a configured estimate in `nfiesta.t_result`. `nfiesta.fn_system_utilization()` is written in `plpython3u`; it calls Python's `os.getloadavg()` and `os.cpu_count()` and returns server load, CPU count, and the current role's connection limit. Restrict that API and audit every Python function for file, process, environment, and network access.

Statistical correctness depends on sampling design, strata, weights, auxiliary variables, missing-data policy, model configuration, and release-specific SQL. Preserve input lineage and configuration, isolate ETL and GUI roles, review security-definer and dynamic-SQL functions, and lock the extension version for reproducible reports. Test migrations and rollback on a copy because upgrades can alter many functions, views, configuration tables, and result generations.
