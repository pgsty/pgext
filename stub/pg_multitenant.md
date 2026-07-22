## Usage

Sources:

- [Official README](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/README.md)
- [Extension implementation](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/src/pgmt.rs)
- [Upstream example](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/scripts/test.sql)

`pg_multitenant` is an experimental helper for shared-database, shared-schema multitenancy. It creates PostgreSQL row-level security policies that compare a designated tenant column with the current role name, and records marked columns in `pgmt.table_tenant_column`.

### Core Workflow

The usable upstream example employs the `user` strategy: each tenant connects or assumes a role whose name equals the tenant key converted to text.

```sql
CREATE EXTENSION pg_multitenant;

CREATE TABLE public.tenant_item (
    tenant_id uuid NOT NULL,
    item_id bigint GENERATED ALWAYS AS IDENTITY,
    payload jsonb,
    PRIMARY KEY (tenant_id, item_id)
);

SET pgmt.tenant_strategy = 'user';
SELECT pgmt.mark_tenant_column(
    'public',
    'tenant_item',
    'tenant_id'
);

CREATE ROLE "11111111-1111-1111-1111-111111111111";
GRANT SELECT, INSERT, UPDATE, DELETE
ON public.tenant_item
TO "11111111-1111-1111-1111-111111111111";

SET ROLE "11111111-1111-1111-1111-111111111111";
SELECT * FROM public.tenant_item;
RESET ROLE;
```

To remove the policy and metadata entry:

```sql
SELECT pgmt.unmark_tenant_column('public', 'tenant_item');
```

`pgmt.mark_tenant_column(schema_name, table_name, tenant_column)` enables RLS, creates a policy named `tenant_isolation_policy`, and records the table OID and column. `pgmt.unmark_tenant_column(schema_name, table_name)` disables RLS, drops that policy, and deletes the metadata row.

### Security and Implementation Caveats

This reviewed revision should not be treated as a complete tenant-isolation framework. The source registers both GUC settings under the name `pgmt.tenant_strategy`; consequently, the advertised `value` strategy and `pgmt.tenant_value` path are not safely usable as written. The helper also interpolates supplied identifiers into dynamic SQL without identifier quoting, uses a fixed policy name, and does not make repeated marking idempotent.

PostgreSQL table owners, superusers, and roles with `BYPASSRLS` can bypass ordinary RLS unless the surrounding design accounts for them. The generated policy compares `tenant_column::TEXT` with `current_user`, so role naming is part of the security contract. Use only trusted identifiers, test owner and privileged-role behavior, inspect the generated policy in `pg_policy`, and perform an independent security review before production use. The control version is `1.0`, while the frozen Cargo manifest reports `0.1.0` and targets PostgreSQL `16`.
