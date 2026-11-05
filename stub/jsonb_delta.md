## Usage

Sources:

- [Official README](https://github.com/evoludigit/jsonb_delta/blob/e3ed58ca75b22cdef187674ff0447ab47d77091b/README.md)
- [Control file](https://github.com/evoludigit/jsonb_delta/blob/e3ed58ca75b22cdef187674ff0447ab47d77091b/jsonb_delta.control)
- [Versioned SQL](https://github.com/evoludigit/jsonb_delta/blob/e3ed58ca75b22cdef187674ff0447ab47d77091b/sql/jsonb_delta--0.3.1.sql)
- [Changelog](https://github.com/evoludigit/jsonb_delta/blob/e3ed58ca75b22cdef187674ff0447ab47d77091b/CHANGELOG.md)

`jsonb_delta` 0.3.1 provides JSONB changesets, keyed array edits and merge helpers. Functions return new JSONB values; persisting a result still uses an ordinary PostgreSQL row update and MVCC.

### Core Workflow

Enable the installed extension and update an array element by its key:

```sql
CREATE EXTENSION jsonb_delta;

SELECT jsonb_array_update_where(
  '{"orders":[{"id":1,"status":"new"}]}'::jsonb,
  'orders', 'id', '1'::jsonb,
  '{"status":"shipped"}'::jsonb
);
```

The result contains the updated order. Use the returned value in an UPDATE expression when the change should persist; a standalone SELECT does not modify a stored document.

### Main Functions

- `jsonb_apply_changeset(doc, ops)` applies an ordered JSONB array of edits. Supported operations include set, remove, merge, increment and array edits; paths accept strings or arrays of segments.
- `jsonb_array_update_where` and `jsonb_array_delete_where` match an element by key and JSONB value. `jsonb_array_contains_id` tests membership; `jsonb_array_insert_where` supports sorted insertion.
- `jsonb_array_update_where_batch` groups edits for one document; `jsonb_array_update_multi_row` accepts an array of documents and returns a set of results.
- `jsonb_smart_patch_scalar`, `jsonb_smart_patch_nested` and `jsonb_smart_patch_array` merge at different scopes. `jsonb_merge_shallow` and `jsonb_merge_at_path` provide explicit merge helpers.
- `jsonb_extract_id` preserves numeric precision when extracting an identifier.

### Compatibility and Boundaries

Version 0.3.x requires PostgreSQL 14–18: the changelog explicitly drops PostgreSQL 13 despite a stale README badge and Cargo feature. No shared preload or restart is required for ordinary use. The control file sets superuser to false but does not mark the extension trusted; its SQL creates C-language functions, so installation still needs the corresponding privileges.

The extension is not relocatable. Most transformation functions are immutable, strict and parallel safe; check the versioned SQL for each signature, especially optional arguments. Changesets are limited to 10,000 operations and paths to 1,000 segments; integer increment overflow raises an error.

Version 0.3.0 corrected numeric-value matching, text/UUID batch keys and identifier precision. The 0.3.1 upgrade retains the same compiled behavior and only advances the SQL version label. Earlier installations have an upgrade chain through 0.2.0 and 0.3.0. Install matching binaries before applying the SQL upgrade, and replace sessions that have loaded an older library.
