## Usage

Sources:

- [database/extensions/efelant_grpc/efelant_grpc.control](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/extensions/efelant_grpc/efelant_grpc.control)
- [docs/core.md](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/docs/core.md)
- [docs/api.md](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/docs/api.md)
- [database/extensions/efelant_grpc/efelant_grpc.c](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/extensions/efelant_grpc/efelant_grpc.c)
- [database/extensions/efelant_grpc/efelant_grpc--1.0.sql](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/extensions/efelant_grpc/efelant_grpc--1.0.sql)
- [database/migrations/015_api.sql](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/migrations/015_api.sql)

`efelant_grpc` is the Efelant Connect/gRPC-JSON transport worker. It delegates requests to `api.handle_grpc` in the application database; the extension itself does not install the business API or its authorization model.

### Core Workflow

```conf
shared_preload_libraries = 'efelant_grpc'
efelant_grpc.database = 'efelant'
efelant_grpc.listen_addresses = '127.0.0.1'
efelant_grpc.port = 18081
```

```sql
CREATE EXTENSION efelant_grpc;
```

### Operational Boundaries

Use the matching PostgreSQL 17 Efelant schema/migrations before enabling the listener. Superuser installation, preload and restart are required. `efelant_grpc.database` defaults to the application database, `efelant_grpc.port` defaults to 18081, and `efelant_grpc.listen_addresses` defaults to all addresses. Limit binding and network access, and use the application deployment’s TLS proxy. Authentication and authorization live in SQL handlers. Installing only this control/SQL pair creates no standalone service API.
