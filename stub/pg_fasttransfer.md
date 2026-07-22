## Usage

Sources:

- [Official README](https://github.com/arpe-io/pg_fasttransfer/blob/a06073f50427935d1ef9ecfea101b4b77b323a98/README.md)
- [Version 1.0 SQL definitions](https://github.com/arpe-io/pg_fasttransfer/blob/a06073f50427935d1ef9ecfea101b4b77b323a98/pg_fasttransfer--1.0.sql)
- [Version 1.0 C implementation](https://github.com/arpe-io/pg_fasttransfer/blob/a06073f50427935d1ef9ecfea101b4b77b323a98/pg_fasttransfer.c)
- [FastTransfer product documentation](https://fasttransfer.arpe.io/)

`pg_fasttransfer` lets a PostgreSQL backend launch the separately installed FastTransfer program and return its process result as rows. It is an administrative bridge to an external licensed data-transfer tool, not an in-database replication protocol.

### Core Workflow

Install `pgcrypto` and the extension, make the FastTransfer executable and license readable by the PostgreSQL operating-system account, then invoke the administrative function with explicit source, target, and binary paths:

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_fasttransfer;

SELECT *
FROM xp_RunFastTransfer_secure(
    sourceconnectiontype := 'pgsql',
    sourceserver := '127.0.0.1:5432',
    sourceuser := 'transfer_user',
    sourcepassword := pg_fasttransfer_encrypt('source-secret'),
    sourcedatabase := 'source_db',
    sourceschema := 'public',
    sourcetable := 'orders',
    targetconnectiontype := 'msbulk',
    targetserver := '127.0.0.1,1433',
    targetuser := 'transfer_user',
    targetpassword := pg_fasttransfer_encrypt('target-secret'),
    targetdatabase := 'target_db',
    targetschema := 'dbo',
    targettable := 'orders',
    fasttransfer_path := '/opt/fasttransfer',
    license := '/opt/fasttransfer/FastTransfer.lic'
);
```

### Objects and Result

- `pg_fasttransfer_encrypt(text)` produces the encoded password format accepted by the runner.
- `xp_RunFastTransfer_secure(...)` maps its named arguments to FastTransfer command-line options and returns `exit_code`, `output`, `error_message`, `total_rows`, `total_columns`, `transfer_time_ms`, and `total_time_ms`.
- Version `1.0` requires `pgcrypto`. Upstream reports tests on PostgreSQL 15 through 17 on Linux and 16 through 17 on Windows; other combinations are not documented as validated.

### Security and Operations

The reviewed C source constructs a shell command from SQL arguments and calls `popen`; values are enclosed in double quotes but embedded quotes and shell metacharacters are not escaped. Revoke execution from application and untrusted roles, and allow only vetted, fixed inputs from administrators. The same source embeds the password-encryption key as the constant `key`, so `pg_fasttransfer_encrypt(text)` is obfuscation rather than protection against anyone who can inspect the binary or source. Process output may also contain sensitive connection or transfer details.

The transfer runs with the database server account's filesystem, process, and network access and can perform remote bulk writes. Use a least-privileged source account, a narrowly scoped target account, protected license and settings files, explicit timeouts outside the extension, and a rehearsed recovery plan before enabling destructive load modes.
