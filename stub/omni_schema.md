


## Usage

> [omni_schema: Advanced schema management tooling](https://docs.omnigres.org/omni_schema/reference/)

The `omni_schema` extension manages application schemas through migrations and object reloading from filesystems.

### Run Migrations

```sql
SELECT * FROM omni_schema.migrate_from_fs(omni_vfs.local_fs('/path/to/migrations'));
```

Recursively finds `.sql` files, applies them in path name order, skips previously applied migrations. Tracks execution in `omni_schema.migrations` table.

### Reload Schema Objects

```sql
SELECT * FROM omni_schema.load_from_fs(omni_vfs.local_fs('/path/to/project'));
```

Reloads functions, policies, and views from the filesystem. To ignore a file, place `omni_schema[[ignore]]` in it (typically in a comment).

### Migrations Table

| Column       | Type      | Description         |
|:-------------|:----------|:--------------------|
| `id`         | int       | Unique identifier   |
| `name`       | text      | Migration filename  |
| `migration`  | text      | Source code         |
| `applied_at` | timestamp | Application time    |

### Multi-Language Support

Supports SQL, PL/pgSQL (`.sql`), Python (`.py`), Perl (`.pl`), Tcl (`.tcl`), and Rust (`.rs`). Use the `SQL[[...]]` directive to specify function signatures:

```python
# times_ten.py
# SQL[[create function times_ten(a integer) returns integer]]
return a * 10
```
