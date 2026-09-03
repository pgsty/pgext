## 用法

来源：

- [官方文档](https://github.com/codarn/pg_mqtt_pub/blob/36ed725dc7aa9a621749a6c378caed84e64bc76e/README.md)
- [扩展控制文件](https://github.com/codarn/pg_mqtt_pub/blob/36ed725dc7aa9a621749a6c378caed84e64bc76e/pg_mqtt_pub.control)
- [官方仓库](https://github.com/codarn/pg_mqtt_pub)

`pg_mqtt_pub` 通过后台工作进程从 SQL、触发器与定时任务发布 MQTT 消息。

### 启用

将 `pg_mqtt_pub` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_mqtt_pub`：

```ini
shared_preload_libraries = 'pg_mqtt_pub'
```

```sql
CREATE EXTENSION pg_mqtt_pub;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Broker connection status and message metrics
SELECT * FROM mqtt_status();

-- Example output:
--  host      | port | connected | messages_sent | messages_failed | dead_lettered | queue_depth | connected_since | disconnected_since | worker_pid
-- -----------+------+-----------+---------------+-----------------+---------------+-------------+-----------------+--------------------+----------
--  localhost | 1883 | true      |         12847 |               3 |             1 |           0 | 2026-02-18 09:15:00 | (null)             | 1234

-- Dead letter diagnostics
SELECT * FROM mqtt_pub.dead_letters ORDER BY failed_at DESC;

-- Summary view
SELECT * FROM mqtt_pub.dead_letter_summary;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `mqtt_publish` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `mqtt_pub.dead_letters` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `mqtt_status` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `mqtt_pub.dead_letter_summary` | VIEW | 扩展创建的检查或查询视图。 |

### 运维与边界

- 预加载 `pg_mqtt_pub` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
- 扩展会在 `mqtt_pub` 下固定或创建模式对象；权限与备份审查应包含这些对象。
