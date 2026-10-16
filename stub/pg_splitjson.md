## Usage

Sources:

- [README.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/README.md)
- [docs/api-reference.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/docs/api-reference.md)
- [docs/support-matrix.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/docs/support-matrix.md)
- [docs/production-runbook.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/docs/production-runbook.md)
- [sql/pg_splitjson--0.3.0.sql](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/sql/pg_splitjson--0.3.0.sql)
- [sql/pg_splitjson--0.2.0--0.3.0.sql](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/sql/pg_splitjson--0.2.0--0.3.0.sql)
- [pg_splitjson.control](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/pg_splitjson.control)
- [CHANGELOG.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/CHANGELOG.md)

`pg_splitjson` 0.3.0 supports frequent JSON field updates. Upstream validation covers PostgreSQL 18.6 on Linux x86_64; other minors and architectures are not promised. Managed views expose a complete document while dedicated operations can update existing hot fields without rewriting the cold template.

### Core Workflow

```sql
CREATE EXTENSION pg_splitjson;
SELECT splitjson.create_table('public.events', '[["counter"],["state"]]'::jsonb);
INSERT INTO public.events(id, doc)
VALUES (1, '{"counter":1,"state":"new","payload":{"body":"cold"}}');
SELECT splitjson.set_field('public.events', 1, ARRAY['state'], '"ready"'::jsonb);
SELECT splitjson.increment_field('public.events', 1, ARRAY['counter'], 1);
SELECT * FROM public.events;
```

### Operational Boundaries

Run creation, migration and index-management functions as the extension installer. Grant application access to the business view and API schema, keeping private storage inaccessible. `splitjson.set_fields` batches updates; `splitjson.increment_field` increments under a row lock; `splitjson.get_field` and `splitjson.find_ids` read fields; `splitjson.create_path_index` adds a B-tree; `splitjson.drop_table` removes managed objects.

Declare 1–64 nonoverlapping hot paths. Missing fields, structural changes and ordinary whole-document updates can repack the document. Index maintenance, WAL, row locks and vacuum costs remain. Ordinary concurrent view updates can raise `40001` and require an application retry.

Optional SELECT rewriting requires loading the library before planning, using `LOAD 'pg_splitjson'` or session preload; control it with `splitjson.enable_query_rewrite`. The core API does not require a server-wide preload or restart. Formal releases support the 0.1.0 → 0.2.0 → 0.3.0 upgrade chain; unrecognized early prototypes still need logical migration. `splitjson.migrate_table` copies a snapshot into a new view; it does not synchronize subsequent writes or copy all constraints and permissions.

### Conditional Writes and Access Control

`splitjson.row_revision` returns an `epoch:id:revision` token; `splitjson.set_fields_if_revision` rejects stale writes with `40001`. `splitjson.apply_patch` supports RFC 6902 except root-document removal, `splitjson.merge_patch` supports RFC 7396, and `splitjson.upsert_document` writes whole documents. Fixed nonnegative array indexes can appear in hot paths. Ordinary views do not support ON CONFLICT.

`splitjson.enable_rls` and `splitjson.set_policy` apply caller-aware policies on private storage while applications use business views. Policy expressions cannot reference cold data, whole documents, subqueries or user-defined functions. `splitjson.create_partitioned_table` supports only HASH/RANGE partitioning by the identifier column.

### Upgrade, Backup and Recovery

Back up first and install matching libraries on every node, then run `ALTER EXTENSION pg_splitjson UPDATE TO '0.3.0'`. Recreate sessions that loaded the old library. Back up the whole database; exporting only the business view omits private storage.

After backup restore, PITR or takeover, call `splitjson.reinitialize_table` before writes, choosing `restore`, `recovery` or `takeover` for the actual event, then run `splitjson.check_all(true)`. This rotates the revision-token epoch. Online migration uses a separate WAL workflow; native publication/subscription supports identical layouts only. Fence the previous writer before takeover; dual-primary writes are unsupported.
