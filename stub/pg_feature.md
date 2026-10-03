## Usage

Sources:

- [extensions/pg_feature/pg_feature.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_feature/pg_feature.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_feature/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_feature/Cargo.toml)
- [extensions/pg_feature/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_feature/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_feature/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_feature/README.md)

`pg_feature` 0.3.0 provides relational feature synthesis in `pgft`. Register table relationships and time indexes, then produce features using aggregations and temporal cutoffs.

### Core Workflow

```sql
CREATE EXTENSION pg_feature;
CREATE TABLE feature_events (id bigint PRIMARY KEY, created_at timestamptz, value numeric);
SELECT pgft.set_time_index('feature_events', 'created_at');
SELECT * FROM pgft.synthesize_features(
  'event_features', 'feature_events', 'id',
  cutoff_time => TIMESTAMPTZ '2026-10-01 00:00:00+00', max_depth => 1
);
SELECT * FROM pgft.list_features('event_features');
```

### Operational Boundaries

The control does not require superuser-only installation, but object-creation privileges still apply. No preload or Python runtime is required. `set_time_index` and `add_relationship` establish metadata; `synthesize_features` creates feature outputs, and `list_features` inspects definitions. Registered timestamps and cutoffs must prevent future-data leakage. Materialized outputs require explicit refresh, and generated SQL/relationships need review before use on large or sensitive tables. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
