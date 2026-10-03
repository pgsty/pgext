## Usage

Sources:

- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_sequence/pg_sequence.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_sequence/pg_sequence.control)
- [extensions/pg_sequence/src/definition.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_sequence/src/definition.rs)
- [extensions/pg_sequence/src/generate.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_sequence/src/generate.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)

`pg_sequence` 0.3.0 stores document-numbering rules in `pgsequence`, with prefixes, formatting, scopes and yearly/monthly reset policies.

### Core Workflow

```sql
CREATE EXTENSION pg_sequence;
SELECT pgsequence.create_sequence('invoice');
SELECT pgsequence.next_val('invoice', 'customer-a');
SELECT pgsequence.next_formatted('invoice', 'customer-a');
```

### Operational Boundaries

The non-relocatable control is not superuser-only; no preload is required. `create_sequence`, `alter_sequence` and `drop_sequence` manage definitions; `next_val`, `next_formatted`, `current_val` and `preview` expose counters. The returned `seq_id` type supports comparison and indexing. Restrict resets and definition changes, and design scope keys to match business uniqueness rules. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
