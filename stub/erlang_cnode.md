## Usage

Sources:

- [Official documentation](https://github.com/samrose/pg-erl/blob/03e81b7f10c083246456dc798f6228539c04ad05/README.md)
- [Extension control file](https://github.com/samrose/pg-erl/blob/03e81b7f10c083246456dc798f6228539c04ad05/erlang_cnode.control)
- [Official repository](https://github.com/samrose/pg-erl)

`erlang_cnode` Erlang C-node bridge for connecting PostgreSQL sessions to Erlang or Elixir nodes.

### Enablement

Install the files for the intended server, then create `erlang_cnode` in the target database:

```sql
CREATE EXTENSION erlang_cnode;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Connect to an Erlang node
SELECT erlang_connect('testnode@127.0.1.1', 'cookie123');

-- Call a remote function
SELECT erlang_call('testnode@127.0.1.1', 'erlang', 'node', '[]'::jsonb);

-- Disconnect
SELECT erlang_disconnect('testnode@127.0.1.1');
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `erlang_call` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_connect` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_disconnect` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_cast` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_check_connection` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_pending_requests` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_ping` | FUNCTION | Callable function from the reviewed install surface. |
| `erlang_receive_async` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
