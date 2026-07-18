## Usage

Sources:

- [Official extension control file](https://github.com/ntkme/postgresql-cidr/blob/0a79fac02992d1863b6f98a533e51abc25c3d69c/cidr.control)

`cidr` — Transform an IPv4 or IPv6 address range into the minimal set of CIDR blocks

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "cidr";
SELECT extversion
FROM pg_extension
WHERE extname = 'cidr';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
