

## Usage

> [refint: referential integrity trigger functions](https://www.postgresql.org/docs/current/contrib-spi.html)

Provides trigger functions for implementing referential integrity checks (largely superseded by built-in foreign key constraints).

```sql
CREATE EXTENSION refint;
```

### Trigger Functions

| Function | Description |
|---|---|
| `check_primary_key()` | Check referencing table -- ensures referenced row exists |
| `check_foreign_key()` | Check referenced table -- handles cascading actions on delete/update |

### check_primary_key

Trigger on the referencing table (`AFTER INSERT OR UPDATE`). Parameters: referencing column name(s), referenced table name, referenced column name(s).

```sql
CREATE TRIGGER check_fk
  AFTER INSERT OR UPDATE ON orders
  FOR EACH ROW
  EXECUTE FUNCTION check_primary_key('customer_id', 'customers', 'id');
```

### check_foreign_key

Trigger on the referenced table (`AFTER DELETE OR UPDATE`). Parameters: number of referencing tables, action (`cascade`, `restrict`, or `setnull`), primary key column(s), then referencing table and column pairs.

```sql
CREATE TRIGGER check_ref
  AFTER DELETE OR UPDATE ON customers
  FOR EACH ROW
  EXECUTE FUNCTION check_foreign_key(1, 'cascade', 'id', 'orders', 'customer_id');
```
