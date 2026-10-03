## Usage

Sources:

- [Official schema_analyzer.control](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/schema_analyzer/schema_analyzer.control)
- [Official schema_analyzer--1.0.0.sql](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/schema_analyzer/schema_analyzer--1.0.0.sql)
- [Official schema_analyzer.c](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/schema_analyzer/schema_analyzer.c)
- [Official README.md](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/README.md)

`schema_analyzer` 1.0.0 supplies Sinew trigger functions that count document-key occurrences and classify frequently used keys. It is a companion to `document_type`, not a generic trigger for arbitrary JSONB tables. This is historical research source from 2014.

### Enable the Pair

```sql
CREATE EXTENSION document_type;
CREATE EXTENSION schema_analyzer;
```

### Trigger Contract

`analyze_document()` handles row changes by incrementing or decrementing key counts in Sinew metadata tables. Its implementation reads the document from column 2 and expects a per-relation table in `document_schema` with key ID, count, dirty and upgraded fields. `analyze_schema()` uses row-count statistics and a 0.5 frequency threshold to mark keys for promotion or demotion. These functions return trigger values and are invoked by the trigger manager, not ordinary SELECT calls.

### Operational Limits

Install the document extension first, because the installation script attempts to create it if absent. The extension alone does not provision application tables, per-table metadata or triggers: those must follow the Sinew layout. The C code interpolates relation names into metadata SQL and assumes a particular binary document layout; use only trusted, controlled research schemas. Administrator installation is required, no preload is documented, and current PostgreSQL compatibility is unverified.
