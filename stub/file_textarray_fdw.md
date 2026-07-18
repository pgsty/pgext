## Usage

Sources:

- [Official extension control file](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/file_textarray_fdw.control)
- [Official upstream documentation](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/file_textarray_fdw/)

`file_textarray_fdw` — Foreign data wrapper returning each text-file row as a text array.

The reviewed catalog snapshot records version `1.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "file_textarray_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'file_textarray_fdw';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
