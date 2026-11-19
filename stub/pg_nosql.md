## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_nosql/pg_nosql.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_nosql/pg_nosql--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_nosql/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_nosql/nosql_query.c)

`pg_nosql` `1.0` provides json document collections and query helpers. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_nosql;
SELECT doc_create_collection('demo');
SELECT doc_insert('demo', '{"city":"Hilo","age":31}');
SELECT * FROM doc_search('demo', '{"city":"Hilo"}');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

doc_put, doc_patch and doc_remove modify documents; doc_search supports the documented filter operators and sort/field/limit options. Collection metadata and backing tables remain PostgreSQL objects. This API does not implement the MongoDB wire protocol. Review generated table privileges, schema validation rules and dynamic filters; do not expose unrestricted collection management to application roles.

This module corresponds to the former upstream name `alohadb_nosql`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
