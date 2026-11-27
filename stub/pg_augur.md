## Usage

Sources:

- [Official source](https://github.com/matroidbe/pg_extensions-releases/releases/tag/pg_augur-v0.4.2)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_augur-v0.4.2/pg_augur-0.4.2-pg18-darwin-arm64.tar.gz)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_augur-v0.4.2/SHA256SUMS)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_augur-v0.4.2/LICENSE-pg_augur)

`pg_augur` 0.4.2 brings Augur training, model storage and inference into PostgreSQL 18. This release publishes binary bundles with control and SQL files; implementation source is not published there. It uses the Matroid Source Available License v1.0 with internal-use restrictions, not an unrestricted open-source license.

### Enable and Inspect

Install the official bundle matching the server and platform. The published platforms are Linux amd64 and macOS arm64. Linux requires glibc 2.36 or later and libgomp1; the macOS build requires macOS 14 or later and trains XGBoost with one thread. Retain other entries when adding the library, then restart:

```ini
shared_preload_libraries = 'pg_augur'
```

```sql
CREATE EXTENSION pg_augur;
SELECT pgaugur.extension_docs();
```

Creation requires a superuser. The extension is not trusted or relocatable and uses the fixed `pgaugur` schema. Its background training worker requires preload.

### Train and Use a Model

For an existing training table with a target column, enqueue a job and retain its returned identifier:

```sql
SELECT pgaugur.start_training(
    project_name => 'demo',
    source_table => 'public.training_rows',
    target_column => 'target',
    automl => true);
```

Commit the submitting transaction before expecting the worker to see the job. Use `pgaugur.training_status` with the returned ID to inspect state, progress, model ID and errors; `pgaugur.cancel_training` requests cancellation. Inspect metrics and input compatibility before selecting a model with `pgaugur.deploy`. `pgaugur.predict_row` accepts named JSONB features; `pgaugur.predict` accepts a positional double-precision array. `pgaugur.inference_schema` exposes the model's expected input.

### Model State and Access

Projects, models, training jobs, experiments and run artifacts are stored in `pgaugur` tables. Keep their data together in verified backups. Restrict schema and function access to intended roles and review permissions on training relations; a superuser-only installer does not itself make every SQL function private. Training can consume substantial CPU and memory. The release offers versioned upgrade scripts, but exposes no public implementation source for a broader safety audit; test models, upgrades and recovery on a disposable database before adoption.

### Upgrading from 0.4.0

Install the matching 0.4.2 binaries, restart the preloaded server, then run `ALTER EXTENSION pg_augur UPDATE`. The intermediate migration adds project drift tolerances and a trailing optional `tolerances` argument to `pgaugur.import_model` and `pgaugur.check_representative`, plus `pgaugur.set_drift_tolerances`. Existing calls retain their argument behavior. The 0.4.1 to 0.4.2 script changes no SQL objects; model-selection and SHAP behavior changes are in the binary.
