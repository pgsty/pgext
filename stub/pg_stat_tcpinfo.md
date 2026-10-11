## Usage

Sources:

- [README.md](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/46be6d5f55330a6f890a093963110bf477c90fdc/README.md)
- [pg_stat_tcpinfo--1.0.sql](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/46be6d5f55330a6f890a093963110bf477c90fdc/pg_stat_tcpinfo--1.0.sql)
- [LICENSE](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/46be6d5f55330a6f890a093963110bf477c90fdc/LICENSE)
- [pg_stat_tcpinfo.control](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/46be6d5f55330a6f890a093963110bf477c90fdc/pg_stat_tcpinfo.control)

`pg_stat_tcpinfo` exposes TCP connections owned by the server OS user, including client, replication and foreign-server traffic. Version 1.0 requires Linux 4.18 or later and access to `/proc` and `NETLINK_INET_DIAG`.

### Core Workflow

```sql
CREATE EXTENSION pg_stat_tcpinfo;
SELECT pid, dst_addr, dst_port, state, recvq, sendq, tcpinfo
FROM pg_stat_tcpinfo;
```

### Operational Boundaries

The view calls `pg_stat_get_tcpinfo()` and returns process/user IDs, local and remote addresses and ports, state, queue sizes and a `tcpinfo` JSON object. Native installation requires a superuser. Access is granted to `pg_read_all_stats`, inherited by `pg_monitor`; review any additional grants because all sockets of the server OS user are visible. TIME-WAIT sockets have no owner and are omitted. There is no documented PostgreSQL-major support matrix or preload requirement.
