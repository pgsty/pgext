## Usage

Sources:

- [Upstream README](https://github.com/AbdulYadi/pgsocket/blob/c52b088ae1bf9cb356e55b5c63715147c4f2e18b/README.md)
- [Extension SQL](https://github.com/AbdulYadi/pgsocket/blob/c52b088ae1bf9cb356e55b5c63715147c4f2e18b/pgsocket--1.0.sql)
- [Extension control file](https://github.com/AbdulYadi/pgsocket/blob/c52b088ae1bf9cb356e55b5c63715147c4f2e18b/pgsocket.control)

`pgsocket` exposes a TCP client inside the PostgreSQL backend. `pgsocketsend` sends bytes, while `pgsocketsendrcvstxetx` sends bytes and returns a response framed by STX and ETX:

```sql
CREATE EXTENSION pgsocket;

SELECT pgsocketsend(
  '127.0.0.1', 9090, 30,
  decode('48656c6c6f', 'hex')
);

SELECT pgsocketsendrcvstxetx(
  '127.0.0.1', 9090, 30, 40,
  decode('48656c6c6f', 'hex')
);
```

### Compatibility and security

Version `1.0` also defines `pgsocketgetimage`, a specialized request/response function. The README says the project was compiled on Linux against PostgreSQL 10 and provides no newer compatibility matrix, so port testing is required. These functions allow database sessions to initiate outbound network connections; restrict `EXECUTE` privileges, review network egress policy, and avoid untrusted destination or payload inputs.
