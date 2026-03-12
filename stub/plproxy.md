

## Usage

> [plproxy: Database partitioning implemented as procedural language](https://github.com/plproxy/plproxy)

PL/Proxy is a PostgreSQL procedural language handler that enables remote procedure calls
between PostgreSQL databases, with optional sharding.

### Create the Extension

```sql
CREATE EXTENSION plproxy;
```

### Language Statements

PL/Proxy functions use four types of statements:

**Cluster selection** -- connect to a pre-configured cluster:

```sql
CREATE FUNCTION get_user(i_id int) RETURNS SETOF users AS $$
    CLUSTER 'mycluster';
    RUN ON i_id;
$$ LANGUAGE plproxy;
```

**Direct connection** -- use a connection string:

```sql
CREATE FUNCTION get_config(key text) RETURNS text AS $$
    CONNECT 'host=remotehost dbname=config';
    SELECT val FROM config WHERE key = $1;
$$ LANGUAGE plproxy;
```

### Execution Modes

**RUN ON hash** -- route to a specific partition based on a hash:

```sql
CREATE FUNCTION get_user_settings(i_username text) RETURNS SETOF user_settings AS $$
    RUN ON namehash(i_username);
$$ LANGUAGE plproxy;
```

**RUN ON ALL** -- execute on all databases in parallel:

```sql
CREATE FUNCTION get_all_counts() RETURNS SETOF record AS $$
    RUN ON ALL;
    SELECT count(*) FROM users;
$$ LANGUAGE plproxy;
```

**RUN ON ANY** -- randomly select a server:

```sql
CREATE FUNCTION get_random_quote() RETURNS text AS $$
    RUN ON ANY;
    SELECT quote FROM quotes ORDER BY random() LIMIT 1;
$$ LANGUAGE plproxy;
```

### Cluster Configuration

Clusters are configured using SQL/MED (Management of External Data):

```sql
CREATE SERVER mycluster FOREIGN DATA WRAPPER plproxy
    OPTIONS (
        connection_lifetime '1800',
        p0 'host=node0 dbname=mydb',
        p1 'host=node1 dbname=mydb',
        p2 'host=node2 dbname=mydb',
        p3 'host=node3 dbname=mydb'
    );

CREATE USER MAPPING FOR CURRENT_USER
    SERVER mycluster
    OPTIONS (user 'proxy_user', password 'secret');
```
