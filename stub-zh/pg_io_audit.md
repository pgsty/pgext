## 用法

来源：

- [官方资料](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/extensions/pg_io_audit/README.md)
- [官方资料](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/extensions/pg_io_audit/pg_io_audit.control)
- [官方资料](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/LICENSE)
- [官方资料](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/extensions/pg_io_audit/pg_io_audit--1.0.sql)
- [官方资料](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/patches/README.md)
- [官方资料](https://github.com/maryyywq/pg_io_hooks/blob/0985958929dd4ca220b512040f8893f787c9f7ab/docs/limitations.md)

pg_io_audit 是面向 PostgreSQL 17 的研究模块，要求应用项目提供的存储管理器钩子补丁，不能加载到普通未修补服务端。计数表示发往存储管理器的请求，不代表物理 I/O 已成功完成。

### 启用与检查

```ini
shared_preload_libraries = 'pg_io_audit'
pg_io_audit.enabled = on
pg_io_audit.trace_size = 1024
```

重启修补后的服务端，再由管理员安装 SQL 接口。监控函数和视图仅对 pg_monitor 或超级用户开放，重置操作仅允许超级用户执行。

```sql
CREATE EXTENSION pg_io_audit;
SELECT * FROM pg_io_audit_stats;
SELECT * FROM pg_io_audit_trace;
SELECT pg_io_audit_reset();
```

### 结果解释

统计按操作与 fork 汇总调用次数和块数；可选的有限跟踪区尽力记录近期关系文件事件。enabled 可以重载，trace_size 要求重启。重启或崩溃后共享内存中的计数与跟踪均丢失。WAL、SLRU、临时文件及直接目录操作不在这些钩子覆盖范围内。
