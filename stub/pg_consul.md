## Usage

Sources:

- [Official README](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/README.md)
- [Extension SQL for version 0.1.0](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/sql/pg_consul--0.1.0.sql)
- [Official Consul KV notes](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/doc/consul-kv.md)
- [C implementation and GUC definitions](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/src/pg_consul.cpp)

`pg_consul` lets a PostgreSQL backend make synchronous HTTP requests to a Consul agent. Version `0.1.0` can ping an agent, read leader and peer status, and fetch KV entries; it does not expose the documented command-line tool's KV write, delete, CAS, or lock operations as SQL functions.

### Core Workflow

After installing the shared library and its cURL dependency, create the extension and point the current session at a reachable Consul agent.

```sql
CREATE EXTENSION pg_consul;

SET consul.agent_host = '127.0.0.1';
SET consul.agent_port = 8500;
SET consul.agent_timeout = 1000;

SELECT consul_agent_ping();
SELECT consul_status_leader();
SELECT * FROM consul_status_peers();
```

Read one key, or pass `true` to recurse through a key prefix. The optional third argument selects a Consul datacenter.

```sql
SELECT *
FROM consul_kv_get('app/config', false, NULL);

SELECT *
FROM consul_kv_get('services/', true, 'dc1');
```

### Important Objects

- `consul_agent_ping()`: pings the configured agent.
- `consul_agent_ping(text, integer)`: pings an explicit host and port.
- `consul_status_leader()`: returns the Consul leader endpoint as text.
- `consul_status_peers()`: returns peer `host`, `port`, and `leader` status rows.
- `consul_kv_get(text, boolean, text)`: returns key, decoded value, flags, create/modify/lock indexes, and session information.
- `consul.agent_host`: user-settable agent host, default `127.0.0.1`.
- `consul.agent_port`: user-settable agent port, default `8500`.
- `consul.agent_timeout`: user-settable request timeout in milliseconds, default `1000`.

### Security and Failure Behavior

Every call runs inside the PostgreSQL backend and waits for the network request. A timeout, malformed response, or unavailable agent can therefore fail or delay the SQL statement. Keep the timeout bounded and do not place these calls on latency-critical query paths without testing.

The exposed interface has no SQL parameter or GUC for a Consul ACL token, HTTPS CA, or client certificate. Use it only with a trusted local agent or another endpoint protected outside the extension, and restrict EXECUTE privileges because KV values and cluster topology may be sensitive. The README title says `0.10.0`, while the control file, extension SQL, and PGXN metadata identify this reviewed release as `0.1.0`.
