

## Usage

> [babelfishpg_money: SQL Server Money Data Type](https://babelfishpg.org/)

The `babelfishpg_money` extension provides SQL Server-compatible `MONEY` and `SMALLMONEY` data type implementations for PostgreSQL as part of the Babelfish project.

### Enabling

```sql
CREATE EXTENSION babelfishpg_money;
```

### Data Types

- **MONEY** - 8-byte monetary value ranging from -922,337,203,685,477.5808 to 922,337,203,685,477.5807, with fixed 4 decimal places
- **SMALLMONEY** - 4-byte monetary value ranging from -214,748.3648 to 214,748.3647, with fixed 4 decimal places

### Behavior

The extension implements SQL Server's monetary arithmetic rules:

- Fixed-point representation with exactly 4 decimal digits
- SQL Server-compatible rounding behavior for monetary calculations
- Proper casting between MONEY and other numeric types
- Arithmetic operations follow SQL Server semantics (e.g., money / money = money, not float)

### Notes

- Part of the Babelfish for PostgreSQL project
- Works in conjunction with `babelfishpg_common` which provides the base type infrastructure
- The PostgreSQL built-in `money` type has different precision and behavior; this extension provides the SQL Server-compatible variant
