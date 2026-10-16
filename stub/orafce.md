## Usage

Sources:

- [README.asciidoc](https://github.com/orafce/orafce/blob/542aed2e0297898cc9844fede3c4452f894b37ed/README.asciidoc)
- [orafce.control](https://github.com/orafce/orafce/blob/542aed2e0297898cc9844fede3c4452f894b37ed/orafce.control)
- [orafce--4.17.sql](https://github.com/orafce/orafce/blob/542aed2e0297898cc9844fede3c4452f894b37ed/orafce--4.17.sql)
- [orafce--4.16--4.17.sql](https://github.com/orafce/orafce/blob/542aed2e0297898cc9844fede3c4452f894b37ed/orafce--4.16--4.17.sql)
- [file.c](https://github.com/orafce/orafce/blob/542aed2e0297898cc9844fede3c4452f894b37ed/file.c)

`orafce` provides Oracle-compatible functions, types and utility packages. Distribution 4.17.1 uses control and SQL extension version 4.17; the two numbers describe different layers.

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

Installation requires a superuser and creates fixed schemas; no preload is required. Follow the upstream configuration guidance before altering `search_path`. The 4.16 patch releases did not change the SQL version; upgrading to the 4.17 series does require its SQL upgrade. Reconnect as required to use the updated library and verify behavior against the exact installed distribution.

### 4.16.13 Patch

Distribution 4.16.13 fixes additional possible `dbms_sql` crashes. The control version remains 4.16; update the library distribution without inventing a 4.16.13 SQL upgrade.

### 4.17 Compatibility Objects

Distribution 4.17.1 adds `utl_raw` byte operations, including `utl_raw.cast_to_raw`, `utl_raw.cast_to_varchar2`, `utl_raw.substr` and bitwise functions. `utl_raw.substr` rejects a null value, invalid start, nonpositive length or a range past the end. `oracle.hextoraw`, `oracle.rawtohex`, `oracle.empty_clob`, `oracle.empty_blob` and `oracle.from_tz` are also added. USER_TABLES and USER_TAB_COLUMNS now restrict their results to the current schema. After installing the new distribution, run `ALTER EXTENSION orafce UPDATE TO '4.17'` and reconnect to use the new library.
