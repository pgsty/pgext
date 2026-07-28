## Usage

Sources:

- [Official upstream README](https://github.com/cjavdev/believe-sql/blob/84bda689e85dd82e9d0ed044db139d86a703d519/README.md)
- [Official extension control file (believe.control)](https://github.com/cjavdev/believe-sql/blob/84bda689e85dd82e9d0ed044db139d86a703d519/believe.control)

`believe` — > [!NOTE] > > The Believe API PostgreSQL Extension is currently **experimental** and we're excited for you to experiment with it! Use it for the corresponding SQL or database utility workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION believe;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.11.0`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
