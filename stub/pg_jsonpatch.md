## Usage

Sources:

- [Official upstream README](https://github.com/willruggiano/pg_jsonpatch/blob/153c3317bee8707cf076761e352ad6fde5083116/README.md)
- [Official extension control file (pg_jsonpatch.control)](https://github.com/willruggiano/pg_jsonpatch/blob/153c3317bee8707cf076761e352ad6fde5083116/pg_jsonpatch.control)
- [Official extension SQL (pg_jsonpatch--1.0.0.sql)](https://github.com/willruggiano/pg_jsonpatch/blob/153c3317bee8707cf076761e352ad6fde5083116/pg_jsonpatch--1.0.0.sql)

`pg_jsonpatch` — I have not yet written tests for A.11-A.15 from the RFC. See ./test.sql for more examples. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_jsonpatch;

select jsonb_patch(
  '{"foo":{"bar":"baz","waldo":"fred"},"qux":{"corge":"grault"}}',
  '[{"op":"move","from":"/foo/waldo","path":"/qux/thud"}]'
);
                             jsonb_patch
---------------------------------------------------------------------
 {"foo": {"bar": "baz"}, "qux": {"thud": "fred", "corge": "grault"}}
(1 row)
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `jsonb_patch(target jsonb, patches jsonb)` is an extension function and returns `jsonb`.
- `jsonb_patch_add(target jsonb, path text[], value jsonb)` is an extension function and returns `jsonb`.
- `jsonb_patch_apply(target jsonb, patch jsonb)` is an extension function and returns `jsonb`.
- `jsonb_patch_copy(target jsonb, _from text[], path text[])` is an extension function and returns `jsonb`.
- `jsonb_patch_move(target jsonb, _from text[], path text[])` is an extension function and returns `jsonb`.
- `jsonb_patch_remove(target jsonb, path text[])` is an extension function and returns `jsonb`.
- `jsonb_patch_replace(target jsonb, path text[], value jsonb)` is an extension function and returns `jsonb`.
- `jsonb_patch_split_path(path text)` is an extension function and returns `text[]`.
- `jsonb_patch_test(target jsonb, path text[], value jsonb)` is an extension function and returns `jsonb`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
