


## 用法

> [omni_service: 服务管理](https://docs.omnigres.org/omni_service/management/)

`omni_service` 扩展为 `omni_httpd` 等数据库服务提供统一管理。它是一个模板化扩展。

### 初始化

```sql
SELECT omni_service.instantiate(schema => 'omni_service');
```

### 列出服务

```sql
SELECT * FROM omni_service.current_services;
```

| 列 | 类型 | 描述 |
|:---|:-----|:-----|
| `name` | text | 唯一的服务标识符 |
| `info` | jsonb | 服务元数据和配置 |
| `state` | service_state | `not-running`、`running` 或 `stopped` |

### 启动和停止服务

```sql
SELECT omni_service.service_start('service_name');
SELECT omni_service.service_stop('service_name');

-- 停止所有运行中的服务：
SELECT omni_service.service_stop(name)
FROM omni_service.current_services
WHERE state = 'running';
```

### 操作日志

`service_operations` 表记录所有服务操作，包含 `name`、`postmaster_start_time`、`operation`（`start`/`stop`）和 `applied_at`。
