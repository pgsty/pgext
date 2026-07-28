## Usage

Sources:

- [Official upstream README](https://github.com/eto-ai/rikai-pg/blob/9a245266bb6190fda7e5cc9edc82c80a267f6a84/README.md)
- [Official extension control file (rikai.control)](https://github.com/eto-ai/rikai-pg/blob/9a245266bb6190fda7e5cc9edc82c80a267f6a84/rikai.control)
- [Official extension SQL (rikai--0.1.sql)](https://github.com/eto-ai/rikai-pg/blob/9a245266bb6190fda7e5cc9edc82c80a267f6a84/sql/rikai--0.1.sql)

`rikai` — rikai ML extensions. Use it for the corresponding vector, model, or retrieval workflow. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION rikai;

Create a model via `INSERT INTO`
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `iou(box1 box, box2 box)` is an extension function and returns `real`.
- `ml.create_model_trigger()` is an extension function and returns `TRIGGER`.
- `ml.cuda_info()` is an extension function and returns `JSON`.
- `ml.delete_model_trigger()` is an extension function and returns `TRIGGER`.
- `ml.is_cuda_available()` is an extension function and returns `BOOL`.
- `ml.version()` is an extension function and returns `table`.
- `detection` is an extension-defined type.
- `image` is an extension-defined type.
- `mask` is an extension-defined type.
- `mask_type` is an extension-defined type.
- `ml.models` is a table installed or managed by the extension.
- `ml` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Upstream explicitly says the project is not production-ready.
- Upstream labels part or all of the project experimental.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
