## Usage

Sources:

- [Official extension control file](https://github.com/katrinaibast/constr_name_unif/blob/e69b3a79fcaa4d013e648b812807a607bd14b7ff/constr_name_unif.control)
- [Official upstream documentation](https://github.com/katrinaibast/constr_name_unif/blob/e69b3a79fcaa4d013e648b812807a607bd14b7ff/doc/constr_name_unif.md)

`constr_name_unif` inventories and renames PostgreSQL constraints according to a configurable convention. Use it to standardize names on existing base tables, but treat every rename operation as schema-changing DDL that can affect migrations and tools which refer to constraint names.

### Core Workflow

Install the pure-SQL extension, inspect the constraints and available patterns, and only then run the desired rename function:

```sql
CREATE EXTENSION constr_name_unif;

SELECT * FROM get_all_constraints_in_schema('public');
SELECT * FROM get_all_patterns();

BEGIN;
SELECT rename_all_constraints_in_schema(
  'public',
  'snake_case_with_short_prefix'
);
COMMIT;
```

The bundled `postgresql_default` pattern uses suffixes such as `pkey` and `fkey`. The bundled `snake_case_with_short_prefix` pattern uses prefixes such as `pk` and `fk`. Custom patterns can choose a delimiter, prefix or suffix placement, and abbreviations for each supported constraint type.

### Important Objects

- `get_all_constraints()` and `get_all_constraints_in_schema(text)` return the current constraint inventory; overloads can restrict results by constraint type.
- `get_all_patterns()` lists configured naming patterns.
- `add_pattern(...)` and `delete_pattern(...)` manage naming patterns.
- `add_to_abbreviated_tables(...)`, `empty_abbreviated_tables()`, and `get_abbreviated_tables()` manage table-name abbreviations.
- `rename_all_constraints(text)` and its schema/type-specific variants issue `ALTER TABLE ... RENAME CONSTRAINT` statements for primary-key, foreign-key, check, unique, and exclusion constraints.

### Operational Notes

Generated identifiers are kept within PostgreSQL's 63-byte identifier limit and collisions are disambiguated. Foreign-key names also incorporate the referenced table. The extension has no dry-run API that reports every proposed target name, so first capture the inventory, test the selected pattern on a copy, and run the rename in a controlled transaction with suitable DDL privileges. Review application migrations, monitoring, and administrative scripts for hard-coded constraint names before committing.
