

## 用法

> [tcn: 通过 LISTEN/NOTIFY 触发变更通知](https://www.postgresql.org/docs/current/tcn.html)

提供在行变更时发送 NOTIFY 事件的触发器函数，携带变更行信息，实现异步变更跟踪。

```sql
CREATE EXTENSION tcn;
```

### 触发器函数

| 函数 | 说明 |
|---|---|
| `triggered_change_notification()` | 行变更时发送携带主键信息的 NOTIFY 通知 |

可选参数：自定义频道名称（默认为 `tcn`）。

### 通知载荷格式

```
"table_name",operation,"column"='value',"column"='value'
```

操作类型：`I`（INSERT）、`U`（UPDATE）、`D`（DELETE）。

### 示例

```sql
CREATE TABLE tcndata (
  a int NOT NULL,
  b date NOT NULL,
  c text,
  PRIMARY KEY (a, b)
);

-- 附加触发器
CREATE TRIGGER tcndata_tcn
  AFTER INSERT OR UPDATE OR DELETE ON tcndata
  FOR EACH ROW
  EXECUTE FUNCTION triggered_change_notification();

-- 监听通知
LISTEN tcn;

-- 变更会触发通知：
INSERT INTO tcndata VALUES (1, '2024-01-01', 'test');
-- 通知: "tcndata",I,"a"='1',"b"='2024-01-01'

UPDATE tcndata SET c = 'updated' WHERE a = 1;
-- 通知: "tcndata",U,"a"='1',"b"='2024-01-01'

DELETE FROM tcndata WHERE a = 1;
-- 通知: "tcndata",D,"a"='1',"b"='2024-01-01'

-- 使用自定义频道名称
CREATE TRIGGER my_trigger
  AFTER INSERT OR UPDATE OR DELETE ON my_table
  FOR EACH ROW
  EXECUTE FUNCTION triggered_change_notification('my_channel');
```
