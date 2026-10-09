## Usage

Sources:

- [packages/object-tree/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-tree/README.md)
- [packages/object-tree/sql/object-tree--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-tree/sql/object-tree--0.47.0.sql)
- [packages/object-tree/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-tree/Makefile)
- [packages/object-tree/object-tree.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-tree/object-tree.control)

`object-tree` layers stores, commits and named references over the PGPM object store. `object_tree_public.init_empty_repo` initializes a store, `object_tree_public.rev_parse` resolves a reference, and commit helpers update a tree and its history together.

### Core Workflow

```sql
CREATE EXTENSION "object-tree";
SELECT scope_id, store_id, message, tree_id, date
FROM object_tree_public.commit
ORDER BY date DESC
LIMIT 10;
```

### Operational Boundaries

`object_tree_public.set_many_and_commit` batches path edits; `object_tree_public.set_props_and_commit` updates node data while preserving children. The installation creates persistent commit/reference/store tables and assumes the application role layout plus a callable `uuidv7()` for UUID defaults. Store and scope IDs must belong to the provisioned application. History consumes database storage and is not automatically garbage-collected by these functions. This is object versioning rather than automatic version control of arbitrary SQL tables. Dependencies are `plpgsql`, `pgcrypto`, `object-store`, `pgpm-verify`. The installation SQL grants privileges to the pre-existing roles `authenticated`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
