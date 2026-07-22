## Usage

Sources:

- [Official package page](https://database.dev/danahartweg/pg_protect_columns)
- [Official database.dev installation guide](https://database.dev/docs/install-a-package)

`pg_protect_columns` prevents selected columns from changing during UPDATE by means of a row trigger. It is useful for application-level immutability rules, but it is not a substitute for privileges, row-level security, or constraints. The public package page currently exposes artifact version 0.0.1 even though catalog metadata reports 0.0.2.

### Core Workflow

Install the package through database.dev/dbdev in an environment with the required trusted-language-extension support, then create a trigger on every protected table.

```sql
SELECT dbdev.install('pg_protect_columns');
CREATE EXTENSION IF NOT EXISTS pg_protect_columns;

CREATE TABLE account (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tenant_id bigint NOT NULL,
    display_name text NOT NULL
);

CREATE TRIGGER _protect_columns_before_update
BEFORE UPDATE ON account
FOR EACH ROW
EXECUTE PROCEDURE protect_columns('tenant_id');
```

Changing a protected value raises restrict_violation. The underscore prefix helps this trigger run before later alphabetically named triggers, but trigger order must still be reviewed explicitly.

### Controlled Bypass

- `disable_protection_on_column(text)` sets a transaction-local bypass for one named column.
- `re_enable_column_protection()` clears that bypass. Call it as soon as the controlled update finishes, inside a tightly permissioned routine.
- The bypass setting is not a general authorization system. Do not grant these helpers to roles that must remain unable to change protected values.

### Caveats

- Trigger enforcement applies only where the trigger exists and is enabled. Owners, superusers, trigger-disabling operations, or replication paths may bypass it.
- Schema-qualify the functions if the package is installed outside the active search path.
- database.dev installation depends on pg_tle/dbdev support. Follow its backup guidance for trusted-language extensions, and test logical restore before relying on it for disaster recovery.
