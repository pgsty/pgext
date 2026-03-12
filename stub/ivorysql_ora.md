

## Usage

> [ivorysql_ora: Oracle Compatible extension on Postgres Database](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora)

The `ivorysql_ora` extension provides Oracle compatibility features for PostgreSQL as part of the IvorySQL project. It adds Oracle-compatible data types, functions, and PL/SQL behavior.

### Enabling

```sql
CREATE EXTENSION ivorysql_ora;
```

### Oracle-Compatible Data Types

The extension adds Oracle-style data types including:

- `NUMBER` / `NUMBER(p,s)` - Oracle-compatible numeric type
- `VARCHAR2(n)` - Oracle-compatible variable-length string
- `DATE` - Oracle-style DATE with time component
- `BINARY_FLOAT` / `BINARY_DOUBLE` - IEEE floating point types

### Oracle-Compatible Functions

Provides Oracle-style built-in functions for string manipulation, date arithmetic, numeric operations, and type conversion that behave consistently with Oracle semantics.

### Compatibility Mode

IvorySQL supports an Oracle compatibility mode that changes parser behavior:

```sql
SET compatible_mode TO oracle;  -- enable Oracle compatibility
SET compatible_mode TO pg;      -- revert to standard PostgreSQL
```

In Oracle mode, the SQL parser accepts Oracle-style syntax including:

- Oracle-style outer joins (`(+)` syntax)
- `CONNECT BY` hierarchical queries
- Oracle-style sequences (`sequence.NEXTVAL`)
- Package-style object references
