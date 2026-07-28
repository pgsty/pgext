## Usage

Sources:

- [Official upstream README](https://github.com/z-xiao-m/pg_ocr/blob/7069c5ecfcec4d5eea16d3a843d577ff61bdcb5e/README.md)
- [Official extension control file (pg_ocr.control)](https://github.com/z-xiao-m/pg_ocr/blob/7069c5ecfcec4d5eea16d3a843d577ff61bdcb5e/pg_ocr.control)
- [Official extension SQL (pg_ocr--0.0.1.sql)](https://github.com/z-xiao-m/pg_ocr/blob/7069c5ecfcec4d5eea16d3a843d577ff61bdcb5e/pg_ocr--0.0.1.sql)

`pg_ocr` — A PostgreSQL extension for OCR (Optical Character Recognition) — run text recognition on images directly from SQL. No external services needed, just a single SQL query. Use it for the corresponding vector, model, or retrieval workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_ocr;

-- Recognize text from a file
SELECT pg_ocr_text_from_file('/path/to/data/plate.png');

-- Recognize text from bytea
SELECT pg_ocr_text(pg_read_binary_file('/path/to/data/plate.png'));

-- Get structured result with confidence and bounding box
SELECT r.* FROM pg_ocr(pg_read_binary_file('/path/to/data/plate.png')) AS r;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_ocr(image bytea)` is an extension function and returns `jsonb`.
- `pg_ocr_from_file(image_path text)` is an extension function and returns `jsonb`.
- `pg_ocr_get_engine()` is an extension function and returns `void`.
- `pg_ocr_set_model_path(path text)` is an extension function and returns `void`.
- `pg_ocr_setup()` is an extension function and returns `void`.
- `pg_ocr_text(image bytea)` is an extension function and returns `text`.
- `pg_ocr_text_from_file(image_path text)` is an extension function and returns `text`.
- `pg_ocr_version()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `plpython3u`.
- The control file marks the extension as relocatable.
- The control file marks the extension as trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
