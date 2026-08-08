## Usage

Sources:

- [pg_readme 0.7.1 README](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/README.md)
- [pg_readme 0.7.1 control file](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/pg_readme.control)
- [pg_readme 0.7.1 upgrade SQL](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/sql/pg_readme--0.7.0--0.7.1.sql)
- [Pigsty package matrix](https://pgext.cloud/ext/pg_readme)

`pg_readme` generates Markdown documentation for a PostgreSQL extension or schema from `COMMENT` objects and live catalog metadata. Use it to keep an extension's README close to its SQL definitions and verify the generated output in source control.

### Install and Generate Markdown

```sql
CREATE EXTENSION pg_readme CASCADE;

SELECT pg_extension_readme('my_extension'::name);
SELECT pg_schema_readme('my_schema'::regnamespace);
```

The control file requires `hstore`, is relocatable, and permits non-superuser installation when the caller can install its dependencies and create the objects.

### Add Processing Instructions

Put Markdown and processing instructions in the extension or schema comment:

```sql
COMMENT ON EXTENSION my_extension IS $markdown$
### `my_extension`

What the extension does.

### Reference

<?pg-readme-reference?>

### Colophon

<?pg-readme-colophon?>
$markdown$;
```

`<?pg-readme-reference?>` expands to a catalog-derived object reference. `<?pg-readme-colophon?>` adds generation metadata. Optional instruction attributes can adjust the heading depth when embedding generated sections.

### Settings

- `pg_readme.include_view_definitions`: include view definitions; default `true`.
- `pg_readme.include_routine_definitions_like`: array of routine-name patterns whose definitions are included; default `'{test__%}'`.
- `pg_readme.include_this_routine_definition`: routine-local override for including the current definition.
- `pg_readme.readme_url`: upstream README link used by generated material.

Use `SET` options on a wrapper function or transaction when a project needs reproducible generation settings.

### Version 0.7.1 and Caveats

- Version 0.7.1 fixes PostgreSQL 18 reference generation that could duplicate array/composite table types and `NOT NULL` markers.
- Upstream and the current Pigsty DEB package are 0.7.1, while the current Pigsty RPM package remains 0.7.0. Check `pg_available_extension_versions` before relying on the PostgreSQL 18 fix.
- Generated output reflects the current database catalog, installed extension versions, comments, and generation time. Review diffs instead of assuming two environments produce identical text.
- Catalog introspection does not replace hand-written operational guidance. Keep prerequisites, preload/restart behavior, upgrade notes, and unsafe operations in curated prose.
- The singular setting `pg_readme.include_routine_definition_like` appears in an old README wrapper example, but the documented current GUC is the plural `pg_readme.include_routine_definitions_like`.
