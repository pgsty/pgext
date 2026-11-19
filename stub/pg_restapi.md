## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_restapi/pg_restapi.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_restapi/pg_restapi--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_restapi/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_restapi/restapi_http.c)

`pg_restapi` `1.0` provides experimental sql and worker-based http crud handlers. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_restapi;
SELECT * FROM pg_restapi_endpoints();
SELECT * FROM pg_restapi_status();
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

SQL handlers work separately from the network worker. The worker requires pg_restapi in shared_preload_libraries and a restart; database, schema, port and default_role configure exposure. The source accepts an X-PG-Role header and switches SQL roles, and the worker connects without a dedicated database role. It does not establish a safe public authentication boundary. Keep the listener isolated and use only disposable test data until authentication and privilege handling have been independently reviewed.

This module corresponds to the former upstream name `alohadb_restapi`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
