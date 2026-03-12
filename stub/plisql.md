

## Usage

> [plisql: PL/iSQL procedural language](https://github.com/IvorySQL/IvorySQL/tree/master/src/pl/plisql)

PL/iSQL is an Oracle-compatible procedural language for PostgreSQL, provided by the IvorySQL project. It extends PL/pgSQL with Oracle PL/SQL syntax and semantics.

### Enabling

```sql
CREATE EXTENSION plisql;
```

### Creating Functions

```sql
CREATE OR REPLACE FUNCTION hello_world
RETURN VARCHAR2
AS
BEGIN
    RETURN 'Hello, World!';
END;
/
```

### Oracle-Style Procedures

```sql
CREATE OR REPLACE PROCEDURE update_salary(
    p_emp_id IN NUMBER,
    p_amount IN NUMBER
)
AS
BEGIN
    UPDATE employees SET salary = salary + p_amount WHERE emp_id = p_emp_id;
END;
/

CALL update_salary(100, 5000);
```

### Key Features

- Oracle-style `BEGIN...END` blocks
- `IN`, `OUT`, `IN OUT` parameter modes
- Oracle-style exception handling with named exceptions
- `%TYPE` and `%ROWTYPE` attribute references
- Oracle-compatible cursor syntax (`CURSOR...IS`, `OPEN`, `FETCH`, `CLOSE`)
- `RETURN` instead of `RETURNS` in function declarations
- Package-like variable scoping

### Differences from PL/pgSQL

- Uses `AS` keyword instead of `$$` delimiters
- Supports Oracle-style `/` as statement terminator
- `VARCHAR2`, `NUMBER`, and other Oracle types natively supported
- Oracle-compatible `DBMS_OUTPUT` integration
