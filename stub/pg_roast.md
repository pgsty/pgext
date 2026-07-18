## Usage

Sources:

- [Upstream README](https://github.com/samirketema/pg_roast/blob/ccbf012d01ebbb8edcb13b02add981705dab2308/README.md)
- [Version 1.0 install SQL](https://github.com/samirketema/pg_roast/blob/ccbf012d01ebbb8edcb13b02add981705dab2308/pg_roast--1.0.sql)
- [Background-worker implementation](https://github.com/samirketema/pg_roast/blob/ccbf012d01ebbb8edcb13b02add981705dab2308/pg_roast_bgw.c)

pg_roast runs opinionated PostgreSQL health checks and stores findings in its fixed roast schema. It checks configuration, schema design, indexes, vacuum and bloat indicators, security posture, replication, connections, and workload signals. Version 1.0 targets PostgreSQL 14 and later.

### Manual audits

Manual mode does not require preloading:

```sql
CREATE EXTENSION pg_roast;

SELECT * FROM roast.run();
SELECT severity, check_name, object_name, roast
FROM roast.latest
ORDER BY severity, check_name;

SELECT * FROM roast.summary;
```

Each run persists audit history and findings. Use the latest view for the newest run and the summary view for grouped results.

### Scheduled audits

The optional background worker requires preload configuration and a restart:

```ini
shared_preload_libraries = 'pg_roast'
pg_roast.database = 'mydb'
pg_roast.interval = 3600
```

The database setting is fixed at server start. Review the upstream settings before enabling automatic audits across a production workload.

### Caveats

- Findings are heuristic advice, not automatic proof of a defect. Review workload context, maintenance windows, and PostgreSQL documentation before applying any recommendation.
- Audits execute catalog and statistics queries and write history tables. Measure overhead on large catalogs and protect the roast schema from untrusted users.
- Results can expose object names, configuration, security findings, and query-related operational details. Apply least privilege and an explicit retention policy.
- Preloading is unnecessary for manual runs but mandatory for the background worker; changing startup-only settings requires a restart.
