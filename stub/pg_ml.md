## Usage

Sources:

- [Official source](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/extensions/pg_ml/Cargo.toml)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/extensions/pg_ml/src/lib.rs)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/crates/pg_bgworker/src/supervision.rs)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/CHANGELOG.md)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/extensions/pg_ml/pgbrew.toml)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/extensions/pg_ml/pg_ml.control)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/README.md)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/LICENSE)
- [Official source](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/extensions/pg_ml/README.md)

`pg_ml` 0.3.4 manages PyCaret experiments, model training and prediction through `pgml`, using a managed Python environment and background jobs.

### Core Workflow

```ini
shared_preload_libraries = 'pg_ml'
```

```sql
CREATE EXTENSION pg_ml;
SELECT pgml.setup_venv();
SELECT * FROM pgml.load_dataset('iris');
SELECT * FROM pgml.setup('iris', 'species', project_name => 'iris_classifier');
SELECT * FROM pgml.create_model('iris_classifier', 'rf');
SELECT pgml.predict('iris_classifier', ARRAY[5.1,3.5,1.4,0.2]);
```

### Operational Boundaries

The control requires superuser installation. Preload the library and restart. Python 3.10+ and PyCaret are installed/managed through the upstream environment setup, which can download packages. `setup`, `create_model` and `predict` drive training; `status` and `python_info` inspect the runtime, and 0.3.0 adds by-name probability prediction. Restrict model/environment operations and external data access. The fixed schema overlaps the naming used by other ML projects; do not assume co-installation compatibility. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

### Current Release and Upgrade

Extension version 0.3.2 uses `pg_ml.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. `pg_ml.training_database` remains a deprecated alias. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above.

### Worker Supervision

Version 0.3.2 adds `pgml.worker_status()` and the superuser-only `pgml.reset_workers()`. Status reports worker name, state, failures, restarts, PID, time of entry into the current state and last failure. `pg_ml.max_worker_failures` defaults to 10 consecutive failures; 0 retries indefinitely. Restart delays back off from 5 to 60 seconds, and reaching the limit leaves a failed worker idle until reset.

Install matching libraries, restart PostgreSQL to initialize the new shared memory, and update the extension SQL in the configured worker database. Inspect the failure before resetting; a reset does not repair its cause.

```sql
ALTER EXTENSION pg_ml UPDATE;
SELECT * FROM pgml.worker_status();
```

### Idle Workers in 0.3.3

Repository release 0.7.2 contains extension 0.3.3. Read-only or empty-claim polls no longer allocate a transaction ID merely because they use the SPI bridge; an actual row write still does. This removes empty commit WAL from idle polling. Install matching libraries, restart the preloaded workers, and update the SQL extension. The existing privilege, external-service and pre-0.3.0 migration boundaries remain.

### macOS Workers in 0.3.4

Repository release 0.7.3 contains extension 0.3.4 and fixes the macOS worker's Python interpreter environment. The macOS virtual environment excludes CatBoost where a suitable ARM64 wheel is unavailable; XGBoost and LightGBM also require libomp. Do not assume identical model availability across platforms; inspect the actual environment after setup. PostgreSQL features cover 14–18, with the existing preload, restart and worker-database requirements unchanged.
