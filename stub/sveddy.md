## Usage

Sources:

- [Official upstream README](https://github.com/robbiegm/sveddy/blob/aee5495b2e4cfc0207a415a1fcd479c63808ed8e/README.md)
- [Official extension control file (sveddy.control)](https://github.com/robbiegm/sveddy/blob/aee5495b2e4cfc0207a415a1fcd479c63808ed8e/sveddy.control)
- [Official extension SQL (sveddy--0.1.0.sql)](https://github.com/robbiegm/sveddy/blob/aee5495b2e4cfc0207a415a1fcd479c63808ed8e/sveddy--0.1.0.sql)

`sveddy` — Sveddy is an in-database collaborative filtering system for PostgreSQL. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION sveddy;

SELECT predict_uv(
    (SELECT weights FROM ratings_sveddy_model_u WHERE id = 3),
    (SELECT weights FROM ratings_sveddy_model_v WHERE id = 5)
);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_initial_weights_uv(integer)` is an extension function and returns `real[]`.
- `predict_uv(real[], real[])` is an extension function and returns `real`.
- `update_model_uv()` is an extension function and returns `TRIGGER`.
- `garbage_collect_uv` is an extension procedure.
- `initialize_model_uv` is an extension procedure.
- `train_uv` is an extension procedure.
- `sveddy_models_uv` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
