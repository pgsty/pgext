## Usage

Sources:

- [Official extension control file](https://github.com/eulerto/pg_inet_addr/blob/ac8bee1ed392e4a07ed75f24cb87414b297fe40c/pg_inet_addr.control)
- [Official upstream documentation](https://github.com/eulerto/pg_inet_addr/blob/ac8bee1ed392e4a07ed75f24cb87414b297fe40c/README.md)

`pg_inet_addr` — Lists the IPv4 and IPv6 addresses and netmasks assigned to the PostgreSQL server host network interfaces.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_inet_addr";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_inet_addr';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
