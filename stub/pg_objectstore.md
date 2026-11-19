## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_objectstore/pg_objectstore.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_objectstore/pg_objectstore--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_objectstore/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_objectstore/objectstore_http.c)

`pg_objectstore` `1.0` provides cloud object storage access from sql. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_objectstore;
SELECT * FROM objectstore_credential_list();
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

objectstore_credential_register stores credentials; objectstore_get, objectstore_put, objectstore_list and objectstore_delete access configured providers. External writes/deletes are not undone by SQL rollback. Restrict function execution and credential-table access to trusted operators, validate destinations and protect network egress. Provider support and authentication must be checked against the shipped transport implementation before using real secrets.

This module corresponds to the former upstream name `alohadb_objectstore`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
