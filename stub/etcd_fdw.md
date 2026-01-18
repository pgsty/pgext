

## Usage

- [Intro Blog](https://www.cybertec-postgresql.com/en/bringing-etcd-to-the-database-with-rust-and-pgrx/)
- [Github Repo](https://github.com/cybertec-postgresql/etcd_fdw)


-----------

## Quick Start

### 1. Enable Extension

```sql
CREATE EXTENSION etcd_fdw;
```

### 2. Create Foreign Data Wrapper

```sql
CREATE FOREIGN DATA WRAPPER etcd_fdw
  HANDLER etcd_fdw_handler
  VALIDATOR etcd_fdw_validator;
```

### 3. Create Server

```sql
-- Basic connection
CREATE SERVER etcd_plain
  FOREIGN DATA WRAPPER etcd_fdw
  OPTIONS (connstr '127.0.0.1:2379');

-- Production etcd with authentication and SSL
CREATE SERVER etcd FOREIGN DATA WRAPPER etcd_fdw OPTIONS (
    connstr '127.0.0.1:2379',
    username 'root',
    password 'Etcd.Root',
    ssl_ca '/pg/cert/ca.crt',
    ssl_cert '/pg/cert/server.crt',
    ssl_key '/pg/cert/server.key'
);
```

### 4. Create Foreign Table

```sql
-- Basic table mapping all keys
CREATE FOREIGN TABLE etcd_kv (key TEXT, value TEXT) SERVER etcd OPTIONS (rowid_column 'key');

-- Table with prefix filter (only keys starting with '/config/')
CREATE FOREIGN TABLE etcd_config (key TEXT, value TEXT)
  SERVER etcd OPTIONS (rowid_column 'key', prefix '/config/');
```

### 5. Query Data

```sql
-- Read all keys
SELECT * FROM etcd_kv;

-- Filter by key pattern (pushdown supported)
SELECT * FROM etcd_kv WHERE key LIKE '/app/%';

-- Range query
SELECT * FROM etcd_kv WHERE key >= '/a' AND key < '/b';

-- Insert new key
INSERT INTO etcd_kv (key, value) VALUES ('/mykey', 'myvalue');

-- Delete key
DELETE FROM etcd_kv WHERE key = '/mykey';
```

### 6. Real-time Sync with etcd

Changes made outside PostgreSQL are immediately visible:

```bash
# Insert via etcdctl
etcdctl put '/config/db_pool_size' '20'
```

```sql
-- Instantly visible in PostgreSQL
SELECT * FROM etcd_config;
     key               | value
-----------------------+-------
 /config/db_pool_size  | 20
(1 row)
```

-----------

## Reference

### Server Options

| Option            | Required | Description                            |
|-------------------|----------|----------------------------------------|
| `connstr`         | Yes      | etcd endpoint (e.g., `127.0.0.1:2379`) |
| `username`        | No       | Authentication username                |
| `password`        | No       | Authentication password                |
| `ssl_ca`          | No       | CA certificate file path               |
| `ssl_cert`        | No       | Client certificate file path           |
| `ssl_key`         | No       | Client private key file path           |
| `ssl_servername`  | No       | Domain name for TLS verification       |
| `connect_timeout` | No       | Connection timeout (default: `10s`)    |
| `request_timeout` | No       | Request timeout (default: `30s`)       |

### Foreign Table Options

| Option         | Default  | Description                              |
|----------------|----------|------------------------------------------|
| `rowid_column` | Required | Column used as unique row identifier     |
| `prefix`       | None     | Restrict to keys with this prefix        |
| `keys_only`    | `false`  | Fetch only keys, skip values             |
| `revision`     | `0`      | Read at specific etcd revision           |
| `key`          | `\0`     | Starting key for range scan              |
| `range_end`    | None     | Exclusive end key for range scan         |
| `consistency`  | `l`      | `l` (linearizable) or `s` (serializable) |

### Query Pushdown

The following operations are pushed down to etcd for better performance:

- **WHERE**: `=`, `>=`, `>`, `<=`, `<`, `BETWEEN`, `LIKE 'prefix%'`
- **ORDER BY**: Remote sorting
- **LIMIT/OFFSET**: Remote pagination

### Limitations

- `UPDATE` on key column is not supported. Workaround: `INSERT` new key, then `DELETE` old key.
- Requires etcd v3 API.

