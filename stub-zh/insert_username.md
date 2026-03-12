

## 用法

> [insert_username: 记录修改表行的用户](https://www.postgresql.org/docs/current/contrib-spi.html)

提供将当前用户名存储到指定文本列的触发器函数。

```sql
CREATE EXTENSION insert_username;
```

### 触发器函数

| 函数 | 说明 |
|---|---|
| `insert_username()` | 将当前用户名存储到指定列 |

参数：用于存储用户名的文本列名。

### 示例

```sql
CREATE TABLE audit_log (
  id serial PRIMARY KEY,
  data text,
  modified_by text
);

-- 记录插入者
CREATE TRIGGER set_insert_user
  BEFORE INSERT ON audit_log
  FOR EACH ROW
  EXECUTE FUNCTION insert_username('modified_by');

-- 记录更新者
CREATE TRIGGER set_update_user
  BEFORE UPDATE ON audit_log
  FOR EACH ROW
  EXECUTE FUNCTION insert_username('modified_by');

INSERT INTO audit_log (data) VALUES ('test');
SELECT modified_by FROM audit_log;  -- 返回当前用户
```
