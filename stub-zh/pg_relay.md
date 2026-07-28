## 用法

来源：

- [官方上游 README](https://gitlab.com/pebble-it/pg_relay/-/blob/main/README.md)
- [官方扩展控制文件](https://gitlab.com/pebble-it/pg_relay/-/blob/main/pg_relay.control)
- [官方项目页面](https://gitlab.com/pebble-it/pg_relay)

`pg_relay` — pg_relay 允许你的 PostgreSQL 数据库在响应事件时运行 SQL 动作——可靠、自动，并且每个结果都会被记录。当应用程序需要这种特定的数据库功能时，请使用它。请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_relay;

-- Simple: pass an order ID
SELECT pgrelay.notify('new_order', NEW.id::text);
-- Action: SELECT warehouse.reserve_stock($1::bigint)

-- Complex: pass multiple values as JSON
SELECT pgrelay.notify('order_shipped',
    json_build_object('order_id', NEW.id, 'tracking', NEW.tracking_code)::text);
-- Action: SELECT notify_customer.send_shipment_email($1)
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以确保与固定源的一致性。
