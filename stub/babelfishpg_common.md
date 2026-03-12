

## Usage

> [babelfishpg_common: SQL Server Transact SQL Datatype Support](https://babelfishpg.org/)

The `babelfishpg_common` extension provides SQL Server-compatible data type support for PostgreSQL as part of the Babelfish project. It enables PostgreSQL to understand and work with Microsoft SQL Server data types.

### Enabling

```sql
CREATE EXTENSION babelfishpg_common;
```

### SQL Server Data Types

The extension adds the following SQL Server-compatible data types:

- **TINYINT** - 1-byte unsigned integer (0 to 255)
- **SMALLMONEY** - Small monetary value
- **MONEY** - Monetary value (see also `babelfishpg_money`)
- **DATETIME** - SQL Server-style datetime
- **DATETIME2** - Extended precision datetime
- **SMALLDATETIME** - Reduced precision datetime
- **DATETIMEOFFSET** - Date and time with timezone offset
- **BIT** - SQL Server-compatible boolean
- **NCHAR** / **NVARCHAR** - Unicode character types
- **UNIQUEIDENTIFIER** - SQL Server-style UUID
- **VARBINARY** - Variable-length binary data
- **IMAGE** - Legacy binary data type
- **SQL_VARIANT** - Generic data type container
- **XML** - SQL Server-compatible XML type
- **SYSNAME** - System name type (nvarchar(128))

### Key Features

- Provides implicit and explicit type casting between SQL Server and PostgreSQL types
- Supports SQL Server-style collation behavior
- Handles SQL Server-specific type coercion rules
- Works in conjunction with `babelfishpg_tsql` for full T-SQL compatibility

This extension is typically deployed as part of a full Babelfish for PostgreSQL installation and is a prerequisite for `babelfishpg_tsql`.
