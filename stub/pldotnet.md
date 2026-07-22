## Usage

Sources:

- [Official README](https://github.com/Brick-Abode/pldotnet/blob/7edb597e296f69609bb80054202c731861e2d9d8/README.md)
- [Official installation guide](https://github.com/Brick-Abode/pldotnet/blob/7edb597e296f69609bb80054202c731861e2d9d8/INSTALL.md)
- [Official extension SQL for version 0.9](https://github.com/Brick-Abode/pldotnet/blob/7edb597e296f69609bb80054202c731861e2d9d8/pldotnet--0.9.sql)

`pldotnet` embeds .NET in PostgreSQL and installs the untrusted procedural languages `plcsharp` and `plfsharp`. Functions, procedures, `DO` blocks, triggers, set-returning functions, records, and SPI database access can be implemented in C# or F#.

### Prerequisites and Core Workflow

The reviewed 0.9 guide requires PostgreSQL 10 or later, .NET 6 or later, and matching native packages/libraries. Install a package built for the exact PostgreSQL major version before creating the extension.

```sql
CREATE EXTENSION pldotnet;

CREATE FUNCTION add_one(value integer)
RETURNS integer
AS $$
    return value + 1;
$$ LANGUAGE plcsharp;

SELECT add_one(41);

CREATE FUNCTION double_value(value bigint)
RETURNS bigint
AS $$
    value * 2L
$$ LANGUAGE plfsharp;

SELECT double_value(21);
```

The extension SQL creates call handlers, inline handlers, and validators for both languages. SQL arguments and return values are converted to .NET/Npgsql-compatible types. The project also supports code loaded from external assemblies and SPI calls from managed code; consult the official tests for feature-specific signatures and type mappings.

### Security and Runtime Boundaries

`plcsharp` and `plfsharp` are untrusted languages. Managed code executes inside a PostgreSQL backend and can consume process memory, block a session, access the server filesystem or network through .NET APIs, or crash the backend through native interop. Only superusers should be allowed to create such functions, and database roles must not be allowed to replace trusted function bodies.

The .NET runtime and managed dependencies become part of the database server's ABI and patching surface. Validate startup, function compilation, SPI behavior, nullable/type conversions, parallel use, upgrades, and crash recovery with the exact PostgreSQL/.NET/package combination. The catalog labels version 0.9 as preview; do not treat upstream benchmark statements as operational guarantees.
