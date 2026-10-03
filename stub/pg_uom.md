## Usage

Sources:

- [extensions/pg_uom/pg_uom.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/pg_uom.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_uom/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/Cargo.toml)
- [extensions/pg_uom/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_uom/src/seed.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/src/seed.rs)
- [extensions/pg_uom/src/convert.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/src/convert.rs)
- [extensions/pg_uom/src/unit.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/src/unit.rs)

`pg_uom` 0.3.0 stores unit categories and definitions in `pguom`, supporting scalar and compound-unit conversions with dimensional checks.

### Core Workflow

```sql
CREATE EXTENSION pg_uom;
SELECT pguom.seed_si();
SELECT pguom.convert(1.0, 'kg', 'lb');
SELECT pguom.convert(0.0, 'degC', 'degF');
```

### Operational Boundaries

The control does not require superuser-only installation, but object-creation privileges still apply. No preload is required. `seed_si` initializes units; `create_category`, `create_unit`, `list_units` and `compatible_units` manage the catalog. Temperature offsets work for simple conversions, while compound temperatures are rejected. Cross-dimension conversions fail. `pg_uom.enabled` can disable conversion and return the unchanged input, so applications must control that setting and not assume every successful call transformed the value. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
