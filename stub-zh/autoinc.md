

## 用法

> [autoinc: 自增触发器函数](https://www.postgresql.org/docs/current/contrib-spi.html)

提供使用序列自增指定列的触发器函数。

```sql
CREATE EXTENSION autoinc;
```

### 触发器函数

| 函数 | 说明 |
|---|---|
| `autoinc()` | 在插入时（以及可选的更新时）使用序列自增列值 |

参数为（列名, 序列名）的配对。仅当初始值为零或 NULL 时才会替换该值。

### 示例

```sql
CREATE SEQUENCE next_id;
CREATE TABLE ids (id int4, idesc text);

CREATE TRIGGER ids_autoinc
  BEFORE INSERT ON ids
  FOR EACH ROW
  EXECUTE FUNCTION autoinc('id', 'next_id');

INSERT INTO ids (idesc) VALUES ('first');  -- id 自动从 next_id 序列分配

-- 多个自增列
CREATE TRIGGER multi_autoinc
  BEFORE INSERT ON my_table
  FOR EACH ROW
  EXECUTE FUNCTION autoinc('col1', 'seq1', 'col2', 'seq2');
```
