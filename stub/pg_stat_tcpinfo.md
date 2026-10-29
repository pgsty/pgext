## Usage

Sources:

- [README](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/README.md)
- [Installation SQL](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/pg_stat_tcpinfo--1.1.sql)
- [Extension control](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/pg_stat_tcpinfo.control)
- [Release 1.1](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/releases/tag/v1.1)
- [License](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/LICENSE)

`pg_stat_tcpinfo` 1.1 exposes Linux TCP socket statistics through SQL. It helps investigate slow client connections, replication links and foreign-server traffic using round-trip time, retransmissions, congestion control and queue sizes, without shell access to the host.

It requires Linux 4.18 or later, access to `/proc`, and the `NETLINK_INET_DIAG` interface. Install as a superuser; no preload or server restart is required.

### Core Workflow

```sql
CREATE EXTENSION pg_stat_tcpinfo;

SELECT a.pid, a.usename, a.application_name,
       t.dst_addr, t.dst_port, t.state, t.recvq, t.sendq,
       t.tcpinfo->>'rtt' AS rtt_ms,
       t.tcpinfo->>'total_retrans' AS retransmissions,
       t.tcpinfo->>'congestion' AS congestion_control
FROM pg_stat_activity a
JOIN pg_stat_tcpinfo t ON t.pid = a.pid
WHERE t.state = 'ESTABLISHED';
```

### Objects and Metrics

The `pg_stat_tcpinfo` view calls `pg_stat_get_tcpinfo()`. Each row contains a process and OS user ID, local and remote addresses and ports, connection state, receive and send queues, and a `tcpinfo` JSONB object. IPv4 and IPv6 sockets are supported. The `rtt` field is expressed in milliseconds; `total_retrans` is the cumulative retransmission counter.

The view includes every TCP socket owned by the server's OS user, including sockets from other database clusters or programs running under that account. Joining on the backend PID narrows the result to connections visible in the current cluster. Unix-domain connections are not TCP sockets and do not appear; TIME-WAIT sockets have no owner and are omitted. Metrics unavailable on an older kernel may be null.

### Permissions

The installation revokes public execution of the function and grants function and view access to `pg_read_all_stats`; `pg_monitor` inherits that access. Ordinary application roles cannot read the metrics by default. Additional grants expose all sockets owned by the server OS user, not only the caller's database connections.

### Operational Boundaries

Each call scans process descriptors and queries the kernel; it is an on-demand observation, not a historical sampler. Container PID and network namespaces and procfs restrictions determine which sockets and PIDs are visible. Choose a sampling interval appropriate to the number of processes and connections.

Use a package built for the server's PostgreSQL major version. The 1.1 release includes fresh-install SQL for versions 1.0 and 1.1 but no upgrade script between them, so an existing 1.0 installation has no upstream-provided `ALTER EXTENSION pg_stat_tcpinfo UPDATE` path.
