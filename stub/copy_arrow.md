## Usage

Sources:

- [README.md](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/README.md)
- [copy_arrow.control](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/copy_arrow.control)
- [copy_arrow--0.0.1.sql](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/copy_arrow--0.0.1.sql)
- [meson.build](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/meson.build)
- [sql/copy_to_arrow/int32.sql](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/sql/copy_to_arrow/int32.sql)

`copy_arrow` 0.0.1 is an experimental Apache Arrow streaming export and COPY-format extension. It targets the upstream copy-format-extendable-2024-11-v26 PostgreSQL branch; it is not a compatibility claim for ordinary PostgreSQL releases.

### Core Workflow

```sql
CREATE EXTENSION copy_arrow;
CREATE TABLE arrow_example (value integer);
INSERT INTO arrow_example SELECT generate_series(1, 10);
SELECT format_arrow(copy_to_arrow('arrow_example'::regclass));
```

### Objects and Requirements

`copy_to_arrow(regclass)` and `scan_to_arrow(regclass)` return binary `bytea` data; `format_arrow(bytea)` renders it as text. The reviewed regression case exports an integer column. The build requires Apache Arrow C++ and the patched server headers. Installation uses native C functions and therefore needs a superuser; the control allows relocation. No preload or separate background worker is specified. Treat the implementation as a proof of concept and verify the exact patched server and data types before use.
