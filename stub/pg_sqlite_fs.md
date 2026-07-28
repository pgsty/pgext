## Usage

Sources:

- [Official extension control file (pg_sqlite_fs.control)](https://github.com/silverdaz/pg_sqlite_fs/blob/87cbad18f8d31f47287e92991a21a96feb9c49c1/pg_sqlite_fs.control)
- [Official extension SQL (pg_sqlite_fs.sql)](https://github.com/silverdaz/pg_sqlite_fs/blob/87cbad18f8d31f47287e92991a21a96feb9c49c1/pg_sqlite_fs.sql)

`pg_sqlite_fs` — SQLite File System creation. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_sqlite_fs;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `delete_attribute(filename text, inode bigint, name text)` is an extension function and returns `boolean`.
- `delete_attributes(filename text, inode bigint)` is an extension function and returns `boolean`.
- `delete_entry(text, bigint)` is an extension function and returns `SETOF bigint`.
- `insert_attribute(filename text, inode bigint, name text, value text)` is an extension function and returns `boolean`.
- `insert_entries(text, text)` is an extension function and returns `boolean`.
- `insert_entry(text, bigint, text, bigint, ctime bigint DEFAULT 0, mtime bigint DEFAULT 0, nlink bigint DEFAULT 1, size bigint DEFAULT 0, is_dir boolean DEFAULT TRUE)` is an extension function and returns `void`.
- `insert_file(filename text, inode bigint, mountpoint text, relative_path text, header bytea, payload_size bigint, prepend bytea, append bytea)` is an extension function and returns `boolean`.
- `insert_files(text, text)` is an extension function and returns `boolean`.
- `make(filepath text, umask integer default 0o007)` is an extension function and returns `boolean`.
- `remove(text)` is an extension function and returns `boolean`.
- `truncate_attributes(text)` is an extension function and returns `boolean`.
- `truncate_entries(text)` is an extension function and returns `boolean`.
- `truncate_files(text)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
