## Usage

Sources:

- [Official README](https://github.com/zilder/pg_udp/blob/ab0f84f8820996bf58eb597eeb5c299476f9137d/README.md)
- [Extension SQL](https://github.com/zilder/pg_udp/blob/ab0f84f8820996bf58eb597eeb5c299476f9137d/pg_udp--0.1.sql)
- [UDP implementation](https://github.com/zilder/pg_udp/blob/ab0f84f8820996bf58eb597eeb5c299476f9137d/pg_udp.c)

`pg_udp` 0.1 exposes one C function that sends a UDP datagram from a PostgreSQL backend. It can emit a small best-effort notification to a trusted host, but it provides no response, delivery confirmation, retry, queue, or transactional integration.

### Core Workflow

```sql
CREATE EXTENSION pg_udp;

SELECT udp_send('127.0.0.1', 9999, 'cache-invalidate:customer-42');
```

The only object is `udp_send(host cstring, port int, data cstring) RETURNS void`. It resolves an IPv4 host with `gethostbyname`, opens a datagram socket, and calls `sendto` with the bytes up to the first NUL character.

### Safety and Reliability Boundaries

UDP is lossy and unordered. A successful SQL call does not prove the receiver obtained or processed the message, and a later transaction rollback cannot retract a datagram already sent. Never use it as the sole audit, replication, payment, or cache-coherency mechanism.

The pinned implementation does not check DNS failure before dereferencing the resolver result, does not inspect the `sendto` return value, and does not close the socket. An invalid host can crash a backend, send errors can be silently lost, and repeated calls can exhaust file descriptors. It also uses legacy IPv4-only resolver APIs and `cstring` pseudo-type arguments. Restrict execution privileges and network egress, accept only trusted constant hosts/ports/data, and use a maintained transactional outbox or message client for production workloads.
