## Usage

Sources:

- [Official source](https://github.com/DBHammer/Metis/blob/e3eed5d90f1a896258ff55c0602a92199bc9dee4/docs/POSTGRESQL_INTEGRATION.md)
- [Official source](https://github.com/DBHammer/Metis/blob/e3eed5d90f1a896258ff55c0602a92199bc9dee4/pg_metis/metis_pg.control)
- [Official source](https://github.com/DBHammer/Metis/blob/e3eed5d90f1a896258ff55c0602a92199bc9dee4/pg_metis/metis_pg--0.1.sql)

`metis_pg` 0.1 is a research cardinality-estimation extension. It supplies workload-trained estimates through PostgreSQL planner hooks, using either prepared JSON histograms inside the server or an external Java estimator. The official integration guide targets PostgreSQL 14.

### Core Workflow

Install as a superuser, prepare statistics for the target relation, and load the library in each session:

```sql
CREATE EXTENSION metis_pg;
LOAD 'metis_pg';
SELECT metis_pg_version();
SET metis.enable = on;
SET metis.estimator = 'c';
SET metis.stats_dir = '/work/metis_stats/dmv/current';
SELECT metis_reset_stats_cache();
SELECT metis_timing_stats();
```

The directory must contain histograms trained for the actual relation and be readable by PostgreSQL. Training uses Java 21, Maven and a separately installed, licensed Gurobi optimizer; the repository does not supply that commercial dependency. C estimation reads the resulting JSON files and does not need the Java service during planning.

### External Estimation and Limits

Choose `metis.estimator = 'http'` for the external service, with `metis.host`, `metis.port` and `metis.timeout_ms` set to the correct endpoint and request budget. Requests send relation and predicate information outside the backend; protect the service and its statistics.

`metis_estimate_selectivity` and `metis_estimate_selectivity_http` expose direct estimates. `metis_reset_timing_stats` clears timing counters. Supported planner predicates are single-table comparisons of columns with constants using equality or inclusive bounds, combined with conjunctions. Other predicates retain PostgreSQL estimates. Session loading requires no server restart; it does not activate hooks in other sessions. Validate estimates and resulting plans against representative data rather than assuming a workload-trained model generalizes.
