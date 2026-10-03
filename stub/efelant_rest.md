## Usage

Sources:

- [database/extensions/efelant_rest/efelant_rest.control](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/extensions/efelant_rest/efelant_rest.control)
- [docs/core.md](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/docs/core.md)
- [docs/api.md](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/docs/api.md)
- [database/extensions/efelant_rest/efelant_rest.c](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/extensions/efelant_rest/efelant_rest.c)
- [database/extensions/efelant_rest/efelant_rest--1.0.sql](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/extensions/efelant_rest/efelant_rest--1.0.sql)
- [database/migrations/015_api.sql](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/migrations/015_api.sql)

`efelant_rest` is the Efelant REST transport worker. It delegates requests to `api.handle_http` in the application database; the extension itself does not install the business API or its authorization model.

### Core Workflow

```conf
shared_preload_libraries = 'efelant_rest'
efelant_rest.database = 'efelant'
efelant_rest.listen_addresses = '127.0.0.1'
efelant_rest.port = 18080
```

```sql
CREATE EXTENSION efelant_rest;
```

### Operational Boundaries

Use the matching PostgreSQL 17 Efelant schema/migrations before enabling the listener. Superuser installation, preload and restart are required. `efelant_rest.database` defaults to the application database, `efelant_rest.port` defaults to 18080, and `efelant_rest.listen_addresses` defaults to all addresses. Limit binding and network access, and use the application deployment’s TLS proxy. Authentication and authorization live in SQL handlers. Installing only this control/SQL pair creates no standalone service API.
