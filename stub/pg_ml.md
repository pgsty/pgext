## Usage

Sources:

- [extensions/pg_ml/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ml/Cargo.toml)
- [extensions/pg_ml/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ml/src/lib.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_ml/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ml/pgbrew.toml)
- [extensions/pg_ml/pg_ml.control](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ml/pg_ml.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/README.md)
- [extensions/pg_ml/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ml/Cargo.toml)
- [extensions/pg_ml/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ml/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)
- [extensions/pg_ml/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ml/README.md)
- [extensions/pg_ml/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ml/pgbrew.toml)

`pg_ml` 0.3.2 manages PyCaret experiments, model training and prediction through `pgml`, using a managed Python environment and background jobs.

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
