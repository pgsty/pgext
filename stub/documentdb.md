
## Usage

Add to shared_preload_libraries first:

```bash
-    shared_preload_libraries: documentdb_core, pg_stat_statements, auto_explain
+    shared_preload_libraries: pg_documentdb_core, pg_stat_statements, auto_explain
```

Example, create extension and perform DDL & CRUD

```sql
CREATE EXTENSION documentdb_core ;
```

Currently we only have documentdb_core extension, It can be used with the FerretDB 2.0+