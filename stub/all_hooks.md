## Usage

Sources:

- [Official upstream README](https://github.com/frbn/all_hooks/blob/c66e7b101f6fd7b216ea06aed91d9acb438447d5/README.md)
- [Official extension control file (all_hooks.control)](https://github.com/frbn/all_hooks/blob/c66e7b101f6fd7b216ea06aed91d9acb438447d5/all_hooks.control)
- [Official extension SQL (all_hooks--0.1.sql)](https://github.com/frbn/all_hooks/blob/c66e7b101f6fd7b216ea06aed91d9acb438447d5/all_hooks--0.1.sql)

`all_hooks` — Extension for PostgreSQL emitting a message for each of all available hooks. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION all_hooks;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- The control file marks the extension as trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
