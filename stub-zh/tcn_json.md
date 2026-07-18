## 用法

来源：

- [项目 README](https://github.com/ArturFormella/tcn_json/blob/b4d5bcd733905a28c7d3401242861e1333202fea/README.md)
- [扩展 control 文件](https://github.com/ArturFormella/tcn_json/blob/b4d5bcd733905a28c7d3401242861e1333202fea/tcn_json.control)
- [1.1 版 SQL API](https://github.com/ArturFormella/tcn_json/blob/b4d5bcd733905a28c7d3401242861e1333202fea/tcn_json--1.1.sql)
- [触发器实现](https://github.com/ArturFormella/tcn_json/blob/b4d5bcd733905a28c7d3401242861e1333202fea/tcn_json.c)

`tcn_json` 1.1 是 PostgreSQL triggered-change-notification 扩展的修改版本。AFTER 行级触发器发送事务性通知，其中包含标签、操作码以及由记录主键值组成的类 JSON 对象。

### 发布记录变更通知

目标表必须具有有效主键。第一个触发器参数选择 channel，可选的第二个参数替换表名标签：

```sql
CREATE EXTENSION tcn_json;

CREATE TRIGGER orders_tcn_json
AFTER INSERT OR UPDATE OR DELETE ON orders
FOR EACH ROW
EXECUTE FUNCTION triggered_change_notification_json('orders_changes', 'orders');

LISTEN orders_changes;
```

提供两个触发器参数时，典型载荷的 envelope 形如 `"orders",I,{"id":"42"}`。不提供参数时，channel 为 `tcn_json`，标签使用表名。操作码为 `I`、`U` 或 `D`。

### 通知边界

投递遵循 PostgreSQL LISTEN/NOTIFY 语义：通知只在提交后可见，同一事务中的重复载荷可能合并，断开连接的消费者无法取得历史记录，而且载荷大小受限。消费者需要重放、确认或可靠投递时，应使用持久 outbox 或逻辑解码。

触发器只发送主键值，不发送变更列或完整记录。在没有主键的表上执行时会报错，并中止修改语句。应把载荷视为此扩展的文本 envelope，而非独立 JSON 文档；解析前必须验证，并针对主键更新和特殊字符串值编写明确的应用测试。已审查源码的构建说明面向 PostgreSQL 9.5 和 9.6 时代，且没有发布当前兼容矩阵。
