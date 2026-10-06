## Usage

Sources:

- [README.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/README.md)
- [docs/api-reference.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/docs/api-reference.md)
- [sql/pg_splitjson--0.1.0.sql](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/sql/pg_splitjson--0.1.0.sql)
- [CHANGELOG.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/CHANGELOG.md)
- [pg_splitjson.control](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/pg_splitjson.control)

`pg_splitjson` 0.1.0 is an initial PostgreSQL 18 implementation for frequent JSON field updates. Managed views expose a complete document while dedicated operations can update existing hot fields without rewriting the cold template.

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

Optional SELECT rewriting requires loading the library before planning, using `LOAD 'pg_splitjson'` or session preload; control it with `splitjson.enable_query_rewrite`. The core API does not require a server-wide preload or restart. Earlier same-version prototypes require fresh installation and logical migration. `splitjson.migrate_table` copies a snapshot into a new view; it does not synchronize subsequent writes or copy all constraints and permissions.
