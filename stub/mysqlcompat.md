## Usage

Sources:

- [mysqlcompat 0.0.7 README on PGXN](https://pgxn.org/dist/mysqlcompat/0.0.7/README.html)
- [mysqlcompat build rules at the reviewed commit](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/Makefile)
- [mysqlcompat aggregate install script at the reviewed commit](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/all.sql)

`mysqlcompat` 0.0.7 is a pure SQL/PLpgSQL collection of MySQL-like functions, aggregates, operators, and casts intended to reduce porting work. Its PGXS build assembles the versioned extension SQL from the files under `sql_bits`.

```sql
CREATE EXTENSION mysqlcompat;

SELECT true && false;
SELECT format(1234.432, 4);
SELECT ADDTIME('01:00:00.999999'::interval, '02:00:00.999998');
```

Objects are installed in the extension's target schema; upstream examples assume they are visible through the active search path.

### Caveats

- Upstream reports testing only on PostgreSQL 8.2 through 9.5. Verify object creation, overload resolution, and behavior on every supported modern major version.
- This is partial compatibility, not MySQL grammar emulation. Named XOR, DIV, and MOD operators are not implemented, and MySQL interval syntax must be rewritten.
- PostgreSQL strings remain case-sensitive, and several time-like MySQL results are represented with PostgreSQL intervals.
- Functions whose names start with an underscore are internal helpers and should not be called directly.
