## Usage

Sources:

- [Official upstream README](https://github.com/xiayingyin/atree/blob/8d5f94d51d02c90114ae47beb6cd1a2fc8f83ca2/README)
- [Official extension control file (ftree.control)](https://github.com/xiayingyin/atree/blob/8d5f94d51d02c90114ae47beb6cd1a2fc8f83ca2/ftree.control)
- [Official implementation source](https://github.com/xiayingyin/atree/blob/8d5f94d51d02c90114ae47beb6cd1a2fc8f83ca2/ftree.c)

`ftree` — We have made the following changes in order to incorporate the L&Y algorithm into Postgres:. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION ftree;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
