## Usage

Sources:

- [packages/utils/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/utils/README.md)
- [packages/utils/sql/pgpm-utils--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/utils/sql/pgpm-utils--0.47.0.sql)
- [packages/utils/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/utils/Makefile)
- [packages/utils/pgpm-utils.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/utils/pgpm-utils.control)

`pgpm-utils` installs helpers in the utility schema. `utils.mask_pad` keeps the rightmost requested characters or left-pads shorter input; `utils.bitmask_pad` does the same for bit strings. This operation is not a privacy-redaction function.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-utils";
SELECT utils.mask_pad('123456', 4, '0');
SELECT utils.bitmask_pad(B'101', 8, '0');
```

### Operational Boundaries

The pinned SQL defines `utils.throw()` as a trigger that raises an error, not a text-argument scalar function. `utils.ensure_singleton()` counts rows and raises if any exist; it does not itself provide concurrency-safe uniqueness or permit the README's generic UPDATE example. Other triggers include `utils.enforce_identity_providers_quota` and `utils.default_self_reference`. The versioned SQL takes precedence over conflicting README examples. Use real constraints and an appropriate locking strategy for data invariants. Dependencies are `plpgsql`, `pgpm-verify`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
