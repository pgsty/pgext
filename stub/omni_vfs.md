


## Usage

> [omni_vfs: Virtual File System](https://docs.omnigres.org/omni_vfs/reference/)

The `omni_vfs` extension provides a unified Virtual File System API for PostgreSQL with multiple backend support.

### Core Functions

**`omni_vfs.list(fs, path)`** -- Lists directory contents, returns set of `(name, kind)`.

**`omni_vfs.list_recursively(fs, path, max)`** -- Recursive listing. Use `max` to limit results.

**`omni_vfs.file_info(fs, path)`** -- Returns file metadata: `size`, `created_at`, `accessed_at`, `modified_at`, `kind`.

**`omni_vfs.read(fs, path, file_offset, chunk_size)`** -- Reads file content as `bytea`.

**`omni_vfs.write(fs, path, content, create_file, append)`** -- Writes content, returns bytes written.

### Backends

**Local filesystem:**

```sql
SELECT omni_vfs.local_fs('/mount/path');
-- No access outside this directory is permitted
```

**Database-backed (table_fs):**

```sql
SELECT omni_vfs.table_fs('filesystem_name');
```

**Remote filesystem:**

```sql
SELECT omni_vfs.remote_fs('dbname=otherdb host=127.0.0.1', $$omni_vfs.local_fs('/path')$$);
```

### Example

```sql
WITH fs AS (SELECT omni_vfs.local_fs('/data') AS fs),
     entries AS (SELECT fs, (omni_vfs.list(fs, '')).* FROM fs)
SELECT entries.name, entries.kind,
       (omni_vfs.file_info(fs, '/' || entries.name)).*
FROM entries;
```
