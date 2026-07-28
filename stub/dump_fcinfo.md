## Usage

Sources:

- [Official upstream README](https://github.com/evancarroll/pg-dump-fcinfo/blob/3defa8fe6769e15fdca9185bbfda3eb707a9962d/README.md)
- [Official extension control file (dump_fcinfo.control)](https://github.com/evancarroll/pg-dump-fcinfo/blob/3defa8fe6769e15fdca9185bbfda3eb707a9962d/dump_fcinfo.control)
- [Official extension SQL (dump_fcinfo--0.0.1.sql)](https://github.com/evancarroll/pg-dump-fcinfo/blob/3defa8fe6769e15fdca9185bbfda3eb707a9962d/dump_fcinfo--0.0.1.sql)

`dump_fcinfo` — Have you ever wondered what is in fcinfo the FunctionCallInfoBaseData that stores a C function's context in PostgreSQL. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION dump_fcinfo;

SELECT dump_fcinfo();
SELECT * FROM dump_fcinfo();
SELECT * FROM dump_fcinfo() LIMIT 5;
SELECT dump_fcinfo() FROM ( VALUES (1) ) AS t(x);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `dump_fcinfo()` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
