## 用法

来源：

- [官方资料](https://github.com/jobin2005/postgresql-apt-behaviour-analysis/blob/bc6e7eacefefd2ff2859bf19b28d9d01663016d3/README.md)
- [官方资料](https://github.com/jobin2005/postgresql-apt-behaviour-analysis/blob/bc6e7eacefefd2ff2859bf19b28d9d01663016d3/src/apt_guard.control)
- [官方资料](https://github.com/jobin2005/postgresql-apt-behaviour-analysis/blob/bc6e7eacefefd2ff2859bf19b28d9d01663016d3/src/apt_guard--1.0.sql)
- [官方资料](https://github.com/jobin2005/postgresql-apt-behaviour-analysis/blob/bc6e7eacefefd2ff2859bf19b28d9d01663016d3/Dockerfile.postgres)

`apt_guard` 1.0 为研究型行为安全流程记录数据库活动。上游容器使用 PostgreSQL 15，原生模块负责采集事件，外部 Python 与机器学习服务提供外围分析和缓解流程。它不是经过独立验证的数据库防火墙。

### 采集事件

配置共享动态库并重启 PostgreSQL：

```ini
shared_preload_libraries = 'apt_guard'
```

由超级用户在被监控数据库中安装，并检查已采集事件：

```sql
CREATE EXTENSION apt_guard;
SELECT * FROM apt_events ORDER BY event_time DESC LIMIT 10;
SELECT * FROM apt_alerts ORDER BY id DESC LIMIT 10;
```

扩展安装用户、会话、事件、序列模式与告警表。执行器和工具语句钩子采集命令文本、耗时、结果与关系信息。外部分析代理和仪表盘需要单独部署，创建扩展不会启动它们。

### 边界与维护

查询文本可能包含敏感值，应限制采集表和外部服务的访问，配置保留策略，并在用于繁忙数据库前测量采集开销。内部采集表查询会被有意跳过，因此这些事件不能作为完整、独立的审计轨迹。错误和事务行为可能影响持久化，需要测试工作负载相关的失败路径。

容器示例凭据仅用于开发，不应作为生产默认值。核对的仓库没有许可证文件，这里不假定已授予复用许可。升级时应匹配动态库和 SQL 对象，并重启已经加载旧模块的后端。
