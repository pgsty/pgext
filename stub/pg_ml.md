## Usage

Sources:

- [extensions/pg_ml/pg_ml.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_ml/pg_ml.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_ml/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_ml/Cargo.toml)
- [extensions/pg_ml/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_ml/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_ml/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_ml/README.md)
- [extensions/pg_ml/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_ml/pgbrew.toml)

`pg_ml` 0.3.0 manages PyCaret experiments, model training and prediction through `pgml`, using a managed Python environment and background jobs.

### Core Workflow

```conf
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
