

## 用法

> [refint: 引用完整性触发器函数](https://www.postgresql.org/docs/current/contrib-spi.html)

提供用于实现引用完整性检查的触发器函数（大部分已被内置外键约束取代）。

```sql
CREATE EXTENSION refint;
```

### 触发器函数

| 函数 | 说明 |
|---|---|
| `check_primary_key()` | 检查引用表 -- 确保被引用行存在 |
| `check_foreign_key()` | 检查被引用表 -- 处理删除/更新时的级联操作 |

### check_primary_key

引用表上的触发器（`AFTER INSERT OR UPDATE`）。参数：引用列名、被引用表名、被引用列名。

```sql
CREATE TRIGGER check_fk
  AFTER INSERT OR UPDATE ON orders
  FOR EACH ROW
  EXECUTE FUNCTION check_primary_key('customer_id', 'customers', 'id');
```

### check_foreign_key

被引用表上的触发器（`AFTER DELETE OR UPDATE`）。参数：引用表数量、操作（`cascade`、`restrict` 或 `setnull`）、主键列，然后是引用表和列的配对。

```sql
CREATE TRIGGER check_ref
  AFTER DELETE OR UPDATE ON customers
  FOR EACH ROW
  EXECUTE FUNCTION check_foreign_key(1, 'cascade', 'id', 'orders', 'customer_id');
```
