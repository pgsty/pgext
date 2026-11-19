## 用法

来源：

- [Official source](https://github.com/munakoiso/logerrors/blob/v2.1.6/README.md)
- [Official source](https://github.com/munakoiso/logerrors/blob/v2.1.6/logerrors.control)
- [Official source](https://github.com/munakoiso/logerrors/releases/tag/v2.1.6)

`logerrors` 通过服务端钩子收集 WARNING、ERROR 和 FATAL 消息计数。2.1.6 修复 PostgreSQL 19 的 ABI/构建兼容问题，SQL 扩展版本仍为 2.1。

### 启用与查询

将库加入已有预加载列表，重启 PostgreSQL，然后创建 SQL 扩展：

```ini
shared_preload_libraries = 'logerrors'
```

```sql
CREATE EXTENSION logerrors;
SELECT * FROM pg_log_errors_stats();
SELECT * FROM pg_slow_log_stats();
SELECT pg_log_errors_reset();
```

`pg_log_errors_stats()` 返回统计区间、消息类型、计数、角色/数据库名称及 SQLSTATE。`pg_slow_log_stats()` 返回慢日志计数和重置时间；`pg_log_errors_reset()` 重置统计。

### 配置

`logerrors.interval` 以毫秒控制统计间隔，默认 5000、最大 60000。`logerrors.intervals_count` 默认保留 120 个区间，最多 360 个。`logerrors.excluded_errcodes` 接受逗号分隔的错误码。

这些是有限的内存统计，并非持久审计档案。应按数据库策略限制运行信息的访问。由于库在启动时预加载，替换库后需要重启；单凭 SQL extversion 2.1 无法识别运行库是否包含 2.1.6 修复。
