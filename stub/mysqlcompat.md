## Usage

Sources:

- [mysqlcompat 0.0.7 README on PGXN](https://pgxn.org/dist/mysqlcompat/0.0.7/README.html)
- [Official upstream README](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/README.md)
- [Version 0.0.7 build manifest](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/Makefile)
- [Aggregate SQL entry point](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/all.sql)

`mysqlcompat` 0.0.7 is a pure SQL/PLpgSQL collection of MySQL-like functions, aggregates, operators, and casts intended to reduce application-porting work. It provides selected compatibility helpers; it does not change PostgreSQL grammar or make PostgreSQL behave exactly like MySQL.

### Core Workflow

```sql
CREATE EXTENSION mysqlcompat;

SELECT true && false;
SELECT format(1234.432, 4);
SELECT ADDTIME(
  '01:00:00.999999'::interval,
  '02:00:00.999998'::interval
);
```

The PGXS build assembles the extension from topic files under `sql_bits`. Major groups include aggregate, bit, cast, control-flow, date/time, information, mathematical, miscellaneous, operator, and string helpers. Names beginning with `_` are implementation helpers and should not be called directly.

### Compatibility Boundaries

- Upstream reports tests only on PostgreSQL 8.2 through 9.5. Validate installation, overload resolution, casts, and results on every modern target major.
- Named MySQL operators `XOR`, `DIV`, and `MOD` cannot be reproduced as grammar; upstream suggests PostgreSQL `#`, integer division, and `%` alternatives.
- PostgreSQL strings remain case-sensitive by default, unlike many MySQL collations.
- Several MySQL time-like values are represented as PostgreSQL `interval`, and MySQL interval syntax must be rewritten with quoted PostgreSQL interval literals.
- Overloaded functions may require explicit casts, as the `ADDTIME` example shows.

The extension creates many broadly named objects in its target schema. Install it in an isolated schema when possible, control `search_path`, and inventory conflicts with core or application functions before exposing it. Ported code still needs semantic tests for NULL handling, overflow, collation, date arithmetic, aggregate behavior, and implicit casts; successful object creation is not proof of MySQL compatibility.
