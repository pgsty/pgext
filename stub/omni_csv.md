


## Usage

> [omni_csv: CSV toolkit](https://docs.omnigres.org/omni_csv/)

The `omni_csv` extension provides CSV parsing, inspection, and generation functions.

### Inspect CSV Structure

```sql
SELECT * FROM omni_csv.csv_info(E'name,age,city\nJohn,25,NYC\nJane,30,LA');
-- Returns: name, age, city (one row per column)
```

### Parse CSV into Records

```sql
SELECT * FROM omni_csv.parse(E'name,age,city\nJohn,25,NYC\nJane,30,LA')
    AS t(name text, age text, city text);
```

Column structure must be specified using the `AS t(...)` clause and should match the CSV structure.

### Generate CSV from Query Results

```sql
SELECT omni_csv.csv_agg(t)
FROM (SELECT name, age, city FROM employees ORDER BY name) t;
```

Returns CSV text with headers and properly escaped values. Returns empty string on empty result sets.

### Limitations

- Large CSV strings are processed entirely in memory (no streaming)
- Column types must be declared in the `AS t(...)` clause
- Consider materializing parsed data for repeated access
