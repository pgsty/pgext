## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/simon-engledew/postgresql-rpc/blob/9a21b7825e7c0879b1f6b20a92ec6e4a722af350/README.md)
- [Version 1.0 SQL declaration](https://github.com/simon-engledew/postgresql-rpc/blob/9a21b7825e7c0879b1f6b20a92ec6e4a722af350/extension/rpc--1.0.sql)
- [Unix-socket C client](https://github.com/simon-engledew/postgresql-rpc/blob/9a21b7825e7c0879b1f6b20a92ec6e4a722af350/extension/rpc.c)
- [Extension control file](https://github.com/simon-engledew/postgresql-rpc/blob/9a21b7825e7c0879b1f6b20a92ec6e4a722af350/extension/rpc.control)

`rpc` sends a JSON document to a Unix-domain socket and returns the JSON document sent back by the peer. Each frame is prefixed by a big-endian 32-bit length.

```sql
CREATE EXTENSION rpc;
SET rpc.path = '/var/run/postgresql/rpc.sock';
SET rpc.timeout = 5;
SELECT rpc('{"operation":"validate","value":42}'::json);
```

The default socket is `/var/run/postgresql/rpc.sock`; both `rpc.path` and the zero-to-sixty-second `rpc.timeout` are user-settable. The peer must implement the exact framing protocol and return valid JSON.

Version 1.0 is a small old native client with no authentication, authorization, release history, tests, or compatibility matrix. Any function caller can select another accessible Unix socket. The response length supplied by the peer drives an unchecked memory allocation, so a malicious peer can exhaust a backend. Revoke public execution and setting privileges, constrain filesystem access, and use only with a trusted local service.
