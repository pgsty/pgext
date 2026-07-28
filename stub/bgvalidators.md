## Usage

Sources:

- [Official extension control file (bgValidators.control)](https://api.pgxn.org/src/bgvalidators/bgvalidators-0.1.2/bgValidators.control)
- [Official extension SQL (bgValidators--1.0.sql)](https://api.pgxn.org/src/bgvalidators/bgvalidators-0.1.2/bgValidators--1.0.sql)

`bgvalidators` — Validation functions for IBAN, Bulgarian EGN personal IDs, BULSTAT VAT numbers, and LNCh foreigner IDs. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION bgvalidators;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `validate_bulstat` is an extension function.
- `validate_egn` is an extension function.
- `validate_iban` is an extension function.
- `validate_lnch` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
