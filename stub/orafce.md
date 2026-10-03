## Usage

Sources:

- [README.asciidoc](https://github.com/orafce/orafce/blob/d905cb474fb8e2e31589f3c75940a3b9e7feb014/README.asciidoc)
- [orafce.control](https://github.com/orafce/orafce/blob/d905cb474fb8e2e31589f3c75940a3b9e7feb014/orafce.control)
- [orafce--4.16.sql](https://github.com/orafce/orafce/blob/d905cb474fb8e2e31589f3c75940a3b9e7feb014/orafce--4.16.sql)
- [4.16.12 release notes](https://github.com/orafce/orafce/releases/tag/VERSION_4_16_12)
- [File-access implementation](https://github.com/orafce/orafce/blob/d905cb474fb8e2e31589f3c75940a3b9e7feb014/file.c)

`orafce` provides Oracle-compatible functions, types and utility packages. Distribution 4.16.12 still uses control and SQL extension version 4.16; the two numbers describe different layers.

### Core Workflow

```sql
CREATE EXTENSION orafce;
SELECT oracle.add_months(date '2026-01-31', 1);
SELECT oracle.nvl(NULL::text, 'fallback');
SELECT oracle.decode(1, 1, 'one', 2, 'two', 'other');
SELECT dbms_output.enable();
SELECT dbms_output.put_line('Hello');
SELECT * FROM dbms_output.get_line();
```

### Types and Packages

Use `oracle.date` when an Oracle-style date must retain the time of day. Date functions include `oracle.add_months`, `oracle.last_day`, `oracle.next_day`, `oracle.months_between`, rounding and truncation. Qualify `oracle.decode`, `oracle.greatest` and `oracle.least` explicitly because PostgreSQL parser handling can otherwise select built-in semantics.

`dbms_output` manages buffered output; `dbms_pipe` and `dbms_alert` support session communication. `dbms_sql` exposes dynamic cursors and typed column retrieval. `dbms_utility`, `dbms_assert`, `plvstr`, `plvchr` and `plvsubst` supply diagnostic, validation and string helpers. These compatibility functions do not turn PostgreSQL into Oracle or provide an Oracle procedural-language runtime.

### File Access and 4.16.12 Changes

`utl_file` accesses server-side files within administrator-configured allowed directories; restrict grants and file-system permissions. The implementation rejects parent-directory references that remain after path canonicalization. Avoid parent references in file paths. The 4.16.12 release specifically fixes possible crashes in `dbms_sql`.

Installation requires a superuser and creates fixed schemas; no preload is required. Follow the upstream configuration guidance before altering `search_path`. Since the SQL version remains 4.16, an installed 4.16 extension need not acquire a new SQL version merely because its binary distribution was patched. Reconnect as required to use the updated library and verify behavior against the exact installed distribution.
