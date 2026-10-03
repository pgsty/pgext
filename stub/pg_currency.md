## Usage

Sources:

- [extensions/pg_currency/pg_currency.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/pg_currency.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_currency/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/Cargo.toml)
- [extensions/pg_currency/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_currency/src/seed.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/src/seed.rs)
- [extensions/pg_currency/src/rate.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/src/rate.rs)
- [extensions/pg_currency/src/convert.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/src/convert.rs)

`pg_currency` 0.3.0 stores currency definitions and dated exchange rates in `pgcurrency`, and converts amounts by base-relative triangulation.

### Core Workflow

```sql
CREATE EXTENSION pg_currency;
SELECT pgcurrency.seed_iso();
SELECT pgcurrency.set_rate('EUR', 0.9, DATE '2026-10-01');
SELECT pgcurrency.convert(100, 'USD', 'EUR', DATE '2026-10-01');
```

### Operational Boundaries

The control does not require superuser-only installation, but object-creation privileges still apply. No preload is required. Rates are supplied by the operator; no live market feed is included. `pgcurrency.base_currency` defaults to USD, and the base rate is one. `set_rate` stores rate/date/type/source, while `get_rate`, `convert` and `round_to_currency` apply historical lookup and currency precision. Missing required rates raise errors. This is distinct from the existing currency-type extension with a similarly named source package. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
