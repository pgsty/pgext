

## 用法

> [moddatetime: 记录修改时间戳](https://www.postgresql.org/docs/current/contrib-spi.html)

提供在行被修改时存储当前时间戳的触发器函数。

```sql
CREATE EXTENSION moddatetime;
```

### 触发器函数

| 函数 | 说明 |
|---|---|
| `moddatetime()` | 在 UPDATE 时将当前时间戳存储到指定列 |

参数：要更新的 `timestamp` 或 `timestamp with time zone` 列名。

### 示例

```sql
CREATE TABLE documents (
  id serial PRIMARY KEY,
  content text,
  modified_at timestamp with time zone
);

CREATE TRIGGER set_modified
  BEFORE UPDATE ON documents
  FOR EACH ROW
  EXECUTE FUNCTION moddatetime('modified_at');

UPDATE documents SET content = 'new content' WHERE id = 1;
-- modified_at 自动设置为当前时间戳
```
