## Usage

Sources:

- [Official README](https://github.com/asotolongo/jx_io/blob/4098d0db4d71596deed1e32a71c587bb1285c01a/README.md)
- [Extension SQL for 0.1.1](https://github.com/asotolongo/jx_io/blob/4098d0db4d71596deed1e32a71c587bb1285c01a/jx_io--0.1.1.sql)

`jx_io` imports and exports JSON or XML between PostgreSQL and files visible to the database server. It creates its objects in the `jx_io` schema and is intended for controlled server-side data exchange, not for handling arbitrary client-local paths.

### Core Workflow

```sql
CREATE EXTENSION jx_io;

SELECT jx_io.export_json_query(
  '/srv/exports/active_users.json',
  'SELECT id, name FROM app_user WHERE active'
);

SELECT jx_io.import_json_doc('/srv/imports/document.json');
```

Export functions return true after writing the file. Import functions either return a JSON/XML document or copy line-oriented JSON documents into a target column.

### Main Functions

- `export_json_query(filename text, query text)` and `export_json_table(filename text, tbl text)` write JSON rows.
- `export_xml_query(filename text, query text)` and `export_xml_table(filename text, tbl text)` write XML.
- `import_json_doc(filename text)` returns one JSON document.
- `import_json_docs_into_table(filename text, tbl text, col text)` copies documents into a table column.
- `import_xml(filename text)` returns one XML document.

### Operational and Security Boundaries

Paths are interpreted on the PostgreSQL server. The export functions build and execute server-side `COPY` commands; import uses PostgreSQL large-object file access. The caller therefore needs the relevant server-file privileges, and the PostgreSQL operating-system account must be able to access the path.

Version 0.1.1 constructs dynamic SQL by concatenating the supplied query, table, column, filename, and path text. Pass only administrator-controlled values; do not expose these functions directly to untrusted users or application input. The upstream README also warns that this version contains known bugs, so validate data and error handling before production use.
