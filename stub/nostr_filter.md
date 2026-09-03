## Usage

Sources:

- [Official documentation](https://github.com/chebizarro/nostrpgx/blob/25ba7715e58dca4486cc6a5a81f8ad1493f58b2b/README.md)
- [Extension control file](https://github.com/chebizarro/nostrpgx/blob/25ba7715e58dca4486cc6a5a81f8ad1493f58b2b/nostr_filter.control)
- [Official repository](https://github.com/chebizarro/nostrpgx)

`nostr_filter` Nostr event types, insertion helpers, and JSON filter evaluation inside PostgreSQL.

### Enablement

Install the files for the intended server, then create `nostr_filter` in the target database:

```sql
CREATE EXTENSION nostr_filter;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT test_insert_nostr_event();
SELECT test_retrieve_nostr_event();
SELECT test_nostr_event_type();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `insert_nostr_event` | FUNCTION | Callable function from the reviewed install surface. |
| `insert_nostr_event_type` | FUNCTION | Callable function from the reviewed install surface. |
| `nostr_filter_search` | FUNCTION | Callable function from the reviewed install surface. |
| `nostr_event` | TYPE | User-facing data type created by the extension. |
| `nostr_events` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `nostr_events` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `insert_nostr_event` | FUNCTION | Callable function from the reviewed install surface. |
| `insert_nostr_event_type` | FUNCTION | Callable function from the reviewed install surface. |
| `nostr_event` | TYPE | User-facing data type created by the extension. |
| `nostr_filter_search` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
