


## Usage

> [pg_readme: Auto-generate README documentation from PostgreSQL COMMENT objects](https://github.com/bigsmoke/pg_readme)

Generates `README.md` documents for extensions or schemas based on `COMMENT` objects in the `pg_description` system catalog.

### Generate Extension README

```sql
SELECT pg_extension_readme('my_extension'::name);
```

### Generate Schema README

```sql
SELECT pg_schema_readme('my_schema'::regnamespace);
```

### Processing Instructions

Include these XML processing instructions in your `COMMENT ON EXTENSION` or `COMMENT ON SCHEMA`:

- `<?pg-readme-reference?>` -- replaced with a full object reference
- `<?pg-readme-colophon?>` -- adds a colophon with pg_readme info

### Typical Workflow

Create a readme-generating function in your extension:

```sql
CREATE FUNCTION my_ext_readme() RETURNS text
    VOLATILE SET search_path FROM current
    SET pg_readme.include_view_definitions TO 'true'
    LANGUAGE plpgsql AS $$
DECLARE _readme text;
BEGIN
    CREATE EXTENSION IF NOT EXISTS pg_readme WITH VERSION '0.1.0';
    _readme := pg_extension_readme('my_extension'::name);
    RAISE transaction_rollback;
EXCEPTION WHEN transaction_rollback THEN RETURN _readme;
END; $$;
```

Then generate with: `make README.md`

### Settings

| Setting | Default |
|---------|---------|
| `pg_readme.include_view_definitions` | `true` |
| `pg_readme.include_routine_definitions_like` | `'{test__%}'` |
| `pg_readme.include_this_routine_definition` | `null` |
