## Usage

Sources:

- [Official extension control file](https://github.com/RekGRpth/pg_grpc/blob/adbdf3a00c68581b6691c48bb45e81f8295e442e/pg_grpc.control)

`pg_grpc` — PostgreSQL gRPC

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_grpc";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_grpc';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
