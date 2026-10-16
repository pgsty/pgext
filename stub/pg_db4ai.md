## Usage

Sources:

- [README.md](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/README.md)
- [sql/pg_db4ai--0.1.0.sql](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/sql/pg_db4ai--0.1.0.sql)
- [docs/ARCHITECTURE.md](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/docs/ARCHITECTURE.md)
- [docs/PORTING.md](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/docs/PORTING.md)
- [.github/workflows/ci.yml](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/.github/workflows/ci.yml)
- [LICENSE](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/LICENSE)
- [pg_db4ai.control](https://github.com/wotchin/pg_db4ai/blob/edea3ae743545d558cf27574dfb03e19f55a19aa/pg_db4ai.control)

`pg_db4ai` 0.1.0 runs training and prediction inside a PostgreSQL backend, using native C algorithms and optionally bundled XGBoost. It is an early implementation inspired by openGauss DB4AI, with different algorithms and model formats; existing openGauss model blobs cannot be loaded directly. Upstream targets Linux with PostgreSQL 16–18. Installation requires a superuser, fixes objects in `db4ai`, and needs no preload or restart.

### Core Workflow

```sql
CREATE EXTENSION pg_db4ai;
CREATE TABLE db4ai_example (target float8, feature float8);
INSERT INTO db4ai_example VALUES (2, 1), (4, 2), (6, 3);
SELECT db4ai.train('linear_demo', 'linear_regression',
    'SELECT target, feature FROM db4ai_example ORDER BY feature',
    '{"max_iterations":300,"learning_rate":0.01}'::jsonb);
SELECT db4ai.predict_float8('linear_demo', ARRAY[4.0]::float8[]);
SELECT db4ai.explain_model('linear_demo');
SELECT db4ai.export_model('linear_demo');
```

### Operational Boundaries

Training queries return the label first, followed by scalar numeric or Boolean features. KMeans and PCA take features only. NULL and non-finite values are rejected; prediction features must keep the training order. `db4ai.predict_float8`, `db4ai.predict_bool`, `db4ai.predict_cluster` and `db4ai.transform` serve the corresponding model types. `db4ai.explain_model` reports training metrics, not holdout accuracy. `db4ai.export_model` returns bytes; arbitrary external-model import is absent.

Models live in the WAL-logged `db4ai.pg_db4ai_models` table, are immutable and scoped to the current role. Replace one with `db4ai.drop_model` and training in one transaction. Reads use owner filtering; cross-role sharing is unsupported. Training runs with the caller's table permissions and RLS. It accepts one SELECT but is not a sandbox against side effects in permitted functions. Only protected registry writes switch to the extension owner's identity.

`db4ai.work_mem` defaults to 16 MB before input spills; administrator limits include `db4ai.max_rows` (10000000), `db4ai.max_model_mb` (64 MB), and `db4ai.xgboost_max_input_mb` (128 MB). The last limit covers the dense input copy, not all XGBoost process memory. Training holds a transaction and snapshot; cancellation is cooperative. Model data and its identity sequence participate in logical backups; restore required roles first. All SQL functions are parallel unsafe. Review resource use and privileges before deploying this preview.
