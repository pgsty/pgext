## Usage

Sources:

- [Official extension control file](https://github.com/adunstan/file_fixed_length_record_fdw/blob/53f7875d71cd457ae62daf142156af927630c999/file_fixed_length_fdw.control)

`file_fixed_length_fdw` — Foreign data wrapper for files with fixed-length fields.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "file_fixed_length_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'file_fixed_length_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
