## Usage

Sources:

- [Official README](https://github.com/sabino/openapi_fdw/blob/v0.4.1/README.md)
- [Extension control file](https://github.com/sabino/openapi_fdw/blob/v0.4.1/openapi_fdw.control)
- [pgrx manifest](https://github.com/sabino/openapi_fdw/blob/v0.4.1/Cargo.toml)

`openapi_fdw` maps OpenAPI 3.0/3.1 JSON services to live PostgreSQL foreign tables. It imports typed columns from the published contract and can retain the complete source object in `attrs jsonb`.

### Enablement

Release 0.4.1 publishes builds for PostgreSQL 14–18. Install the exact-major native library, then create the superuser-only, relocatable extension:

```sql
CREATE EXTENSION openapi_fdw;
```

No preload or restart is required. The optional control-plane service helps discover and apply definitions but is not required after configuration.

### Import an API

Create a foreign server that points to a trusted OpenAPI document, then import selected operations into a schema.

```sql
CREATE SERVER pokeapi
FOREIGN DATA WRAPPER openapi_fdw
OPTIONS (
  spec_url 'https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml'
);

CREATE SCHEMA poke;

IMPORT FOREIGN SCHEMA api
  LIMIT TO (pokemon_list)
  FROM SERVER pokeapi
  INTO poke
  OPTIONS (methods 'GET', include_attrs 'true');

SELECT name, attrs ->> 'url'
FROM poke.pokemon_list
LIMIT 5;
```

Every scan performs bounded live HTTP requests; materialize results explicitly if a local snapshot is required. Path-parameter endpoints issue no request until equality predicates bind every path variable.

### Authentication and Writes

Prefer environment-backed options such as `bearer_token_env`, `api_key_env`, and `headers_env`, because the PostgreSQL server process resolves them without storing literal secrets in catalog options. Tables are read-only unless explicit insert/update/delete endpoints are configured.

Remote HTTP writes are not transactional with PostgreSQL: a later rollback cannot undo an accepted request, and a multi-row statement can partially succeed. Version 0.4.1 sends one request per row, does not automatically retry POST/PATCH, and does not support `RETURNING`; design idempotency and compensation at the remote API boundary.
