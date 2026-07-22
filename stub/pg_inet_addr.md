## Usage

Sources:

- [Official README](https://github.com/eulerto/pg_inet_addr/blob/ac8bee911bd3ac3055f13a871bb2af34df254e29/README.md)
- [Extension SQL for version 1.0](https://github.com/eulerto/pg_inet_addr/blob/ac8bee911bd3ac3055f13a871bb2af34df254e29/pg_inet_addr--1.0.sql)
- [C implementation](https://github.com/eulerto/pg_inet_addr/blob/ac8bee911bd3ac3055f13a871bb2af34df254e29/pg_inet_addr.c)

`pg_inet_addr` lists IPv4 and IPv6 addresses assigned to the operating-system interfaces visible to the PostgreSQL backend. It is useful for local or Unix-socket connections where PostgreSQL's `inet_server_addr()` can return NULL.

### Core Workflow

The extension installs its function in `pg_catalog` and revokes its EXECUTE privilege from PUBLIC. A superuser or owner must grant access deliberately.

```sql
CREATE EXTENSION pg_inet_addr;

GRANT EXECUTE ON FUNCTION pg_catalog.pg_inet_addr()
TO monitoring_role;
```

Query all interfaces or filter by family and interface name.

```sql
SELECT interface, address, netmask, family
FROM pg_catalog.pg_inet_addr()
ORDER BY interface, family, address;

SELECT interface, address::inet
FROM pg_catalog.pg_inet_addr()
WHERE family = 'IPv4'
  AND interface <> 'lo';
```

### Returned Columns

- `interface`: operating-system interface name.
- `address`: interface address formatted as text.
- `netmask`: interface netmask formatted as text.
- `family`: `IPv4` or `IPv6`.

The function is declared VOLATILE and returns a fresh snapshot on each call. Both addresses and netmasks are text, so cast only after filtering or validating values. Link-local IPv6 addresses can include a scope suffix such as `%eth0`.

### Operational Notes

Results describe the network namespace of the PostgreSQL server process, not the client's network, DNS name, load-balancer address, or public endpoint. In containers, this normally means the container or pod namespace. Interfaces can appear or disappear between calls, and the function does not identify which address accepts PostgreSQL traffic.

The implementation relies on the host's `getifaddrs` interface and is intended for Unix-like systems. Because the output reveals host networking details, keep the default privilege boundary unless a monitoring role specifically needs it.
