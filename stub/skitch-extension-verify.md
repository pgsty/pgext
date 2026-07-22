## Usage

Sources:

- [Extension control file](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/verify/skitch-extension-verify.control)
- [Version 0.0.7 extension SQL](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/verify/sql/skitch-extension-verify--0.0.7.sql)
- [Official verify package tree](https://github.com/airpage-app/pg-utils/tree/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/verify)

`skitch-extension-verify` 0.0.7 is a collection of PL/pgSQL deployment checks. Each function returns `true` when its catalog query finds the requested object and raises an exception when the check fails, making the extension suitable for controlled migration verification rather than interactive discovery.

### Core Workflow

The extension requires `plpgsql`, `uuid-ossp`, and `skitch-extension-utils`:

```sql
CREATE EXTENSION "skitch-extension-verify" CASCADE;

SELECT verify_schema('app');
SELECT verify_table('app.orders');
SELECT verify_constraint('app.orders', 'orders_pkey');
SELECT verify_table_grant('app.orders', 'SELECT', 'reporter');
```

Use schema-qualified object strings for checks that split a name into schema and entity components. A successful call returns `true`; an absent object normally aborts the current statement with the function's `Nonexistent ...` exception.

### Verification Functions

- Objects: `verify_schema`, `verify_table`, `verify_view`, `verify_type`, `verify_domain`, `verify_extension`, and `verify_function`.
- Table details: `verify_constraint`, `verify_index`, `verify_trigger`, `verify_policy`, and `verify_security`.
- Roles and privileges: `verify_role`, `verify_membership`, and `verify_table_grant`.

The functions depend on helper routines such as `get_schema_from_str`, `get_entity_from_str`, `list_indexes`, and `list_memberships` supplied by `skitch-extension-utils`.

### Correctness Boundaries

These functions are not a complete security or schema-diff engine. In the reviewed SQL, `verify_extension` queries `pg_available_extensions`, so it proves that an extension is available to install rather than installed in the current database. `verify_function` checks that a matching function row exists but does not require the returned `has_function_privilege` value to be true.

Several routines use broad or fragile catalog joins, and all are incorrectly declared `IMMUTABLE` despite reading mutable catalogs and privileges. The failure branch of `verify_security` also references an undefined `_name` variable. Treat these as migration assertions whose SQL has been reviewed for the exact check and PostgreSQL version, not as compliance evidence. Avoid prepared-plan reuse or constant expressions, and independently verify security-sensitive results with direct catalog queries.
