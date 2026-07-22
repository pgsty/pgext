## Usage

Sources:

- [Official README](https://github.com/Minoro/pgpyml/blob/70a119164c9af19d36a04d9592fd9e7f5a7ec521/README.md)
- [Official documentation](https://minoro.is-a.dev/pgpyml/get-started/)
- [Official control file](https://github.com/Minoro/pgpyml/blob/70a119164c9af19d36a04d9592fd9e7f5a7ec521/pgpyml.control)

`pgpyml` loads Python joblib/scikit-learn models inside PostgreSQL and exposes prediction functions plus trigger helpers. Upstream labels it work in progress and not production-ready. Model loading uses untrusted Python and server-local files, so a model artifact is executable code from the database server's security perspective.

### Prerequisites and Core Workflow

Install `plpython3u` and the Python packages needed by the saved model in PostgreSQL's Python environment, then create the extension in a dedicated schema:

```sql
CREATE SCHEMA IF NOT EXISTS pgpyml;
CREATE EXTENSION pgpyml SCHEMA pgpyml CASCADE;

SELECT *
FROM pgpyml.predict(
    '/srv/pgpyml/models/iris.joblib',
    ARRAY[[5.2, 3.5, 1.5, 0.2], [7.7, 2.8, 6.7, 2.0]]
);
```

The model path is resolved on the database server and must be readable by its service account. The nested array contains one feature vector per prediction, and the function returns a text array.

### Important Objects

- `predict` applies a saved model to one or more feature arrays.
- `predict_table_row` reads named columns from a table row before predicting.
- `classification_trigger` writes a prediction into a target column on insert or update.
- `trigger_abort_if_prediction_is` and `trigger_abort_unless_prediction_is` reject rows based on a class.
- Combined trigger helpers both save a classification and conditionally reject the operation.

Trigger arguments include model paths and column names. Test generated triggers against schema changes, nulls, type conversions, model feature ordering, and prediction errors.

### Security and Operations

`plpython3u` is untrusted, and joblib/pickle deserialization can execute arbitrary Python code. Only a tightly controlled administrator should install the extension, write model files, or call APIs with a model path. Store models in an immutable, service-owned directory and verify their digest and provenance before deployment.

Inference runs synchronously inside a backend or row trigger, increasing transaction latency and memory use; exceptions can abort writes. Python environment, native dependencies, model files, and exact scikit-learn versions are not captured by a database dump. Pin them separately, load-test concurrency, and prove restore and model compatibility before any non-disposable use.
