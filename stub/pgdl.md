## Usage

Sources:

- [Official README](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/README.md)
- [Base extension SQL](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/sql/pgdl--1.0.0.sql)
- [Vector-type upgrade SQL](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/sql/pgdl--1.0.0--1.1.0.sql)
- [Model API upgrade SQL](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/sql/pgdl--1.1.0--1.2.0.sql)
- [Current API upgrade SQL](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/sql/pgdl--1.2.0--1.3.0.sql)
- [Model manager implementation](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/src/pgdl/model_manager.cpp)
- [Extension control file](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/pgdl.control)

`pgdl` 1.3.0, also called MorphingDB, embeds deep-learning inference and shaped vector storage in PostgreSQL. It loads TorchScript models into the database backend through libtorch and exposes the custom `mvec` type. The reviewed project targets PostgreSQL 12 and requires a source build with heavyweight native dependencies and model-specific C++ handlers; treat it as a research prototype rather than a drop-in inference service.

### Vector Storage Workflow

After validating a locally patched installation, the self-contained part of the API stores an array together with shape information:

```sql
CREATE EXTENSION pgdl;

CREATE TABLE feature_vectors (
    id bigint PRIMARY KEY,
    value mvec NOT NULL
);

INSERT INTO feature_vectors VALUES
    (1, '[1.0,2.2,3.123,4.2]{2,2}'),
    (2, ARRAY[1.0,2.0,3.0,1.2345]::float4[]::mvec);

SELECT id, get_mvec_shape(value), get_mvec_data(value)
FROM feature_vectors;

UPDATE feature_vectors SET value = value + value WHERE id = 1;
```

`mvec` accepts text or numeric arrays, can preserve a multidimensional shape, supports addition, subtraction, and equality, and can be converted back with `get_mvec_data(mvec)` and inspected with `get_mvec_shape(mvec)`.

### Model Inference Workflow

Model-centric use starts with `create_model`, which records a server-local model path and optional base model in `model_info`. `register_process()` binds pre/post-processing callbacks compiled into the extension, and `predict_float` or `predict_text` runs the registered model on CPU or GPU. Moving-window aggregates `predict_batch_float8` and `predict_batch_text` provide batch inference. Version 1.3.0 adds the task-oriented `ai_operator`, `api_load_model`, and `api_predict` interface.

This is not a data-only registration mechanism: upstream instructs users to implement the model's input and output handlers under `src/external_process`, then rebuild and reinstall the shared library. `image_to_vector`, `text_to_vector`, and `image_classification` add image and SentencePiece preprocessing for the bundled examples.

### Important Objects

- `model_info`, `model_layer_info`, and `base_model_info`: model paths, metadata, layers, and base-model relationships.
- `create_model`, `modify_model`, `drop_model`, `load_base_model`, and `register_process`: model lifecycle and compiled callback registration.
- `predict_float`, `predict_text`, `predict_batch_float8`, and `predict_batch_text`: scalar and windowed inference entry points.
- `ai_operator`, `api_load_model`, and `api_predict`: task-oriented API introduced in 1.3.0.
- `mvec`, `get_mvec_data`, `get_mvec_shape`, and `text_to_mvec`: shaped vector storage and conversion.

### Security, Compatibility, and Packaging Caveats

Model and tokenizer paths are opened by PostgreSQL server processes, and image preprocessing accepts paths that OpenCV opens from the server environment. Grant these functions only to trusted roles: an allowed caller can consume large CPU/GPU memory, access server-reachable files, and run model code inside a database backend. Validate model provenance, resource limits, concurrency, crash isolation, and GPU behavior outside a valuable cluster.

The reviewed source has no direct `pgdl--1.3.0.sql` install script and instead depends on an install/update chain from 1.0.0. Several vector functions in the 1.1.0 upgrade reference `MODULE_PATHNAME/pgdl.so`, although `MODULE_PATHNAME` already resolves to the library path; this is likely to produce an invalid library name on an unpatched installation. The README examples also differ from current SQL signatures. Build and run `CREATE EXTENSION`, the full upgrade chain, vector round trips, and a representative model on the exact PostgreSQL 12 artifact before retaining any data.
