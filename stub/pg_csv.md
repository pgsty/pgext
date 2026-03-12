

## Usage

> [pg_csv: flexible CSV processing as aggregate functions](https://github.com/PostgREST/pg_csv)

Provides a CSV aggregate that composes with SQL expressions, unlike COPY which requires a special protocol. RFC 4180 compliant with proper quoting.

```sql
CREATE EXTENSION pg_csv;
```

### Functions

| Function | Description |
|---|---|
| `csv_agg(record)` | Aggregate rows into CSV text with headers |
| `csv_agg(record, csv_options(...))` | Aggregate with custom options |
| `csv_options(delimiter, bom, header, nullstr)` | Configure CSV output options |

### Examples

```sql
CREATE TABLE projects AS SELECT * FROM (VALUES
  (1, 'Death Star OS', 1),
  (2, 'Windows 95 Rebooted', 1),
  (3, 'Project "Comma,Please"', 2)
) AS _(id, name, client_id);

-- Basic CSV output
SELECT csv_agg(x) FROM projects x;
-- id,name,client_id
-- 1,Death Star OS,1
-- 2,Windows 95 Rebooted,1
-- 3,"Project ""Comma,Please""",2

-- Pipe-separated values
SELECT csv_agg(x, csv_options(delimiter := '|')) FROM projects x;

-- Tab-separated values
SELECT csv_agg(x, csv_options(delimiter := E'\t')) FROM projects x;

-- With BOM for Excel compatibility
SELECT csv_agg(x, csv_options(bom := true)) FROM projects x;
```
