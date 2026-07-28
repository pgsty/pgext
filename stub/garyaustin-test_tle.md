## Usage

Sources:

- [Official database.dev package page](https://database.dev/garyaustin/test_tle)

`garyaustin-test_tle` — test custom-properties. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "garyaustin-test_tle";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `update_to_app_metadata` is an extension function.
- `user_has_property` is an extension function.
- `user_properties_match` is an extension function.
- `user_property_in` is an extension function.
- `property_names` is a table installed or managed by the extension.
- `user_properties` is a table installed or managed by the extension.
- `IF` is a schema created by the extension.
- `on_role_change` is an extension-defined trigger.

### Requirements and Caveats

- The catalog records version `0.0.2`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
