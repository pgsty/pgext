## 用法

来源：

- [官方 README](https://github.com/frost242/monitoring_role/blob/f76f0629de7a8af4e28b8ca92336a5b050ef4e39/README.md)
- [0.0.1 扩展 SQL](https://github.com/frost242/monitoring_role/blob/f76f0629de7a8af4e28b8ca92336a5b050ef4e39/monitoring_role--0.0.1.sql)

`monitoring_role` 创建一组范围受限的监控包装函数，超级用户可以把它们授予原本无特权的角色。这些包装器以定义者权限运行，因此只能由可信管理员安装和授权。

### 核心流程

将扩展安装在专用模式中，通过辅助函数授予监控登录角色访问权，并把该模式放在该角色搜索路径中 `pg_catalog` 的前面。

```sql
CREATE ROLE monitor LOGIN;
CREATE SCHEMA monitoring;
CREATE EXTENSION monitoring_role SCHEMA monitoring;

SELECT monitoring.grant_monitor('monitor');
ALTER ROLE monitor SET search_path = monitoring, pg_catalog, public;
```

以被授权角色重新连接后，未限定模式的监控调用会解析到这些包装器。

```sql
SELECT * FROM pg_stat_activity;
SELECT * FROM pg_ls_dir('.');
SELECT * FROM pg_stat_file('PG_VERSION');
SELECT pg_read_file('PG_VERSION');
```

### 暴露对象与边界

- `pg_stat_activity()` 及同名视图：暴露完整活动视图。
- `pg_ls_dir(text)` 与 `pg_stat_file(text)`：包装对应的系统目录函数。
- 两个 `pg_read_file` 重载只接受字面路径 `PG_VERSION`，其他路径会抛出异常。
- `grant_monitor(text)`：授予模式使用权、视图查询权和包装函数执行权。

创建 0.0.1 版本需要超级用户权限。安装脚本会撤销 public 的执行权，但 `grant_monitor(text)` 本身不是安全定义者函数；应由扩展所有者或其他能够执行授权的角色调用。授权前应根据安全策略审查被暴露的目录与活动信息。
