## Usage

Sources:

- [Official README](https://github.com/AbdulYadi/pgsocket/blob/master/README.md)
- [Extension control file](https://github.com/AbdulYadi/pgsocket/blob/master/pgsocket.control)
- [Version 1.0 extension SQL](https://github.com/AbdulYadi/pgsocket/blob/master/pgsocket--1.0.sql)
- [Version 1.0 C implementation](https://github.com/AbdulYadi/pgsocket/blob/master/pgsocket.c)

pgsocket opens outbound IPv4 TCP connections from a PostgreSQL backend. It can send raw bytes, exchange a message framed by the STX and ETX control bytes, or use an application-specific length-prefixed image protocol. It provides no TLS, authentication, connection pooling, or asynchronous I/O.

### Core Workflow

Version 1.0 installs all functions explicitly in the public schema. The basic send function accepts a numeric IPv4 address, port, send timeout in seconds, and byte payload.

```sql
CREATE EXTENSION pgsocket;

SELECT public.pgsocketsend(
    '192.0.2.10',
    9090,
    30,
    convert_to('Hello', 'UTF8')
);
```

For a service whose reply starts with byte 0x02 and ends with byte 0x03, use the framed request/response function:

```sql
SELECT public.pgsocketsendrcvstxetx(
    '192.0.2.10',
    9090,
    30,
    40,
    convert_to('Hello', 'UTF8')
);
```

The returned byte array excludes the framing bytes. The image helper implements a separate binary protocol with native integer fields and should be used only with a peer built for that exact format.

### Implementation Boundaries

Although the README shows a host-name example, the cataloged C source passes the address through numeric IPv4 conversion and does not perform DNS resolution. Use a numeric address and verify connectivity from the database host.

Each call opens and blocks one PostgreSQL backend until completion or socket timeout. The functions are declared stable in the extension SQL even though they perform network I/O, so do not rely on that volatility label for optimization or replay safety. Calls are not transactional: a later SQL rollback cannot undo bytes already sent.

### Security and Operations

Outbound sockets from SQL create a server-side request and data-exfiltration surface. Revoke execution from public roles, allow only fixed destinations through host firewall policy, and avoid passing user-controlled addresses, ports, or payloads. The protocol is plaintext and does not authenticate the peer.

The control file says relocatable, but the SQL hard-codes the public schema, so schema relocation is not effective for these functions. No preload or restart is declared. The source uses POSIX sockets and must be built and tested against the exact PostgreSQL and operating-system target.
