## Usage

Sources:

- [Upstream manual at the reviewed commit](https://github.com/uptimejp/sql_firewall/blob/39fdec885d748975e7475e05156b91eddff9d7a7/README.sql_firewall)
- [Version 0.8 SQL objects](https://github.com/uptimejp/sql_firewall/blob/39fdec885d748975e7475e05156b91eddff9d7a7/sql_firewall--0.8.sql)
- [Extension control file](https://github.com/uptimejp/sql_firewall/blob/39fdec885d748975e7475e05156b91eddff9d7a7/sql_firewall.control)

`sql_firewall` is an abandoned PostgreSQL 9.4 query allowlist prototype. It learns pairs of user ID and parse-tree-derived query ID, then warns about or rejects queries that are not in the learned set. It must be in `shared_preload_libraries` and requires a restart.

```conf
shared_preload_libraries = 'sql_firewall'
sql_firewall.firewall = 'learning'
```

After creating the extension in an administration database, learn a representative workload, inspect `sql_firewall.sql_firewall_statements`, switch first to `permissive`, and only then consider `enforcing`. The hook applies across the cluster even though management views and functions exist only in the database where the SQL extension was created.

Rule export and import use superuser-supplied server file paths and are allowed only while disabled. Query IDs depend on PostgreSQL internal parse trees, so rule files are not portable across major versions. This code supports PostgreSQL 9.4 only and must not be deployed on current PostgreSQL without a full port and security review. It is defense in depth, not a substitute for parameterized SQL, least privilege, or application validation.
