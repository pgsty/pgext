


## Usage

Sources: [README](https://github.com/credativ/db2fce/blob/master/README.md), [SQL objects](https://github.com/credativ/db2fce/blob/master/db2fce.sql), [control file](https://github.com/credativ/db2fce/blob/master/db2fce.control)

`db2fce` provides a DB2 compatibility environment for PostgreSQL. It creates DB2-style functions, types, operators, and the `SYSIBM.SYSDUMMY1` compatibility view so SQL originally written for IBM Db2 can be adapted with fewer changes.

### Enable

```sql
CREATE EXTENSION db2fce;

SET search_path = db2, sysibm, public;
```

The extension creates most compatibility objects in the `db2` schema and creates `sysibm.sysdummy1` for DB2 queries that expect a dummy single-row table.

```sql
SELECT * FROM sysibm.sysdummy1;
```

### Compatibility Functions

The `db2` schema includes date/time helpers such as `microsecond`, `second`, `minute`, `hour`, `day`, `month`, `year`, `days`, `months_between`, `date`, `time`, and `timestamp_format`.

String and conversion helpers include `locate`, `translate`, `lcase`, `upper`, `lower`, `strip`, `char`, `integer`, `int`, `double`, `decimal`, `dec`, `hex`, `round`, `digits`, and `value`.

### Operators

The SQL layer also defines DB2-style operators such as `^=` for inequality and `!!` for concatenation across several data types.

```sql
SELECT db2.int('42');
SELECT db2.days(current_date);
SELECT 'db' !! '2';
```

### Notes

Adding `db2` to `search_path` lets many DB2 function calls work without schema qualification. Some names that conflict with PostgreSQL syntax or built-in behavior may still need explicit `db2.` qualification.
