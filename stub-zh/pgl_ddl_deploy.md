

## 用法

> [pgl_ddl_deploy: 使用 pglogical 自动化 DDL 部署](https://github.com/enova/pgl_ddl_deploy)

支持 pglogical 和原生逻辑复制的 PostgreSQL 9.5+ 透明 DDL 复制。自动将 DDL 变更（CREATE TABLE、ALTER TABLE 等）传播到订阅者。

### 启用

```sql
CREATE EXTENSION pgl_ddl_deploy;
```

### 配置

向 `pgl_ddl_deploy.set_configs` 表中插入配置：

```sql
-- 复制所有用户模式的 DDL 并自动添加新表
INSERT INTO pgl_ddl_deploy.set_configs (set_name, include_schema_regex, driver)
VALUES ('default', '.*', 'native'::pgl_ddl_deploy.driver);

-- 或使用 pglogical 驱动
INSERT INTO pgl_ddl_deploy.set_configs (set_name, include_schema_regex, driver)
VALUES ('default', '.*', 'pglogical'::pgl_ddl_deploy.driver);

-- 仅维护已在复制中的特定表（仅 ALTER TABLE）
INSERT INTO pgl_ddl_deploy.set_configs (set_name, include_only_repset_tables, driver)
VALUES ('my_tables', TRUE, 'native'::pgl_ddl_deploy.driver);
```

### 部署事件触发器

配置后，部署事件触发器：

```sql
SELECT pgl_ddl_deploy.deploy(set_config_id) FROM pgl_ddl_deploy.set_configs;
```

### 关键配置选项

- `driver`：`native` 或 `pglogical`
- `set_name`：发布名称或 pglogical 复制集名称
- `include_schema_regex`：匹配 DDL 复制模式的正则表达式
- `include_only_repset_tables`：为 true 时，仅对已在复制中的表执行 ALTER TABLE
- `lock_safe_deployment`：为 true 时，DDL 在订阅者上以低 lock_timeout 循环执行
- `allow_multi_statements`：为 true 时，可以传播多语句 DDL
- `queue_subscriber_failures`：为 true 时，订阅者上失败的 DDL 会排队重试
- `ddl_only_replication`：仅复制模式，不自动将表添加到数据复制

### 监控

```sql
-- 查看未处理的 DDL 事件
SELECT * FROM pgl_ddl_deploy.unhandled;

-- 查看订阅者失败的 DDL
SELECT * FROM pgl_ddl_deploy.subscriber_logs WHERE NOT succeeded;

-- 在订阅者上重试失败的 DDL
SELECT pgl_ddl_deploy.retry_all_subscriber_logs();
```

### 检查已解析的模式

```sql
SELECT pgl_ddl_deploy.resolved_regex_include_schemas(set_config_id)
FROM pgl_ddl_deploy.set_configs;
```
