## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/dsef/dsef-2024.4.9/README.md)
- [Official extension control file (dsef.control)](https://api.pgxn.org/src/dsef/dsef-2024.4.9/dsef.control)
- [Official extension SQL (dsef--unpackaged--2024.4.9.sql)](https://api.pgxn.org/src/dsef/dsef-2024.4.9/sql/dsef--unpackaged--2024.4.9.sql)

`dsef` — *Detailed SQL reports for third party help & support*. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION dsef;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ds_capture()` is an extension function and returns `int`.
- `ds_insert(p_run int)` is an extension function and returns `int`.
- `ds_report(IN global_scope boolean DEFAULT TRUE,IN all_rows boolean DEFAULT FALSE)` is an extension function and returns `TABLE`.
- `ds_report_diff(IN global_scope boolean DEFAULT TRUE,IN all_rows boolean DEFAULT FALSE,IN cnt_diff_pct_threshold numeric DEFAULT 0)` is an extension function and returns `TABLE`.
- `ds_set(p_setting text)` is an extension function and returns `int`.
- `ds_start()` is an extension function and returns `int`.
- `ds_version()` is an extension function and returns `text`.
- `explain_analyze_full(p_sql text,p_format text DEFAULT 'TEXT',p_verbose boolean DEFAULT false)` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `2024.4.9`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
