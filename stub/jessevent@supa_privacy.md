## Usage

Sources:

- [Official database.dev package page](https://database.dev/jessevent/supa_privacy)

`jessevent@supa_privacy` — Formatting-preserving anonymisation and data masking. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "jessevent@supa_privacy";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `create_masked_view` is an extension function.
- `generalize_date` is an extension function.
- `generalize_numeric` is an extension function.
- `mask_email` is an extension function.
- `mask_phone` is an extension function.
- `mask_phone_flexible` is an extension function.
- `mask_text` is an extension function.
- `partial_mask` is an extension function.
- `perturb_numeric` is an extension function.
- `perturb_numeric_deterministic` is an extension function.
- `salted_hash` is an extension function.
- `shift_date_deterministic` is an extension function.

### Requirements and Caveats

- The catalog records version `1.0.1`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
