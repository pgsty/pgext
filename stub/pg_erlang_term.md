## Usage

Sources:

- [Official documentation](https://github.com/flambard/pg_erlang_term/blob/32715a42e25ad75da709cdac61bb9af6c8617f61/README.md)
- [Extension control file](https://github.com/flambard/pg_erlang_term/blob/32715a42e25ad75da709cdac61bb9af6c8617f61/pg_erlang_term.control)
- [Official repository](https://github.com/flambard/pg_erlang_term)

`pg_erlang_term` PostgreSQL type for encoding, storing, and decoding Erlang External Term Format values.

### Enablement

Install the files for the intended server, then create `pg_erlang_term` in the target database:

```sql
CREATE EXTENSION pg_erlang_term;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE TABLE messages (payload erlang_term);
INSERT INTO messages VALUES ('{ok,42}'::erlang_term);
SELECT payload FROM messages;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `erlang_term` | TYPE | User-facing data type created by the extension. |
| `erlang_term_decode` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_term_encode` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_term_input` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_term_output` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_term_receive` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_term_send` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
