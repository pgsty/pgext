## 用法

来源：

- [官方 README](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/README.md)
- [1.0 版本 SQL 函数](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/pg_sysdatetime--1.0.sql)
- [扩展实现](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/pg_sysdatetime.c)

`pg_sysdatetime` 为 PostgreSQL 9.4 及更早版本提供 SQL Server 风格的当前时间函数，主要用于提高 Windows 上的时间戳精度。PostgreSQL 9.5 已在可用时原生使用 Windows 精确时钟，因此现代 PostgreSQL 不需要这个已弃用扩展。

### 函数

```sql
CREATE EXTENSION pg_sysdatetime;

SELECT sysutcdatetime(),
       sysdatetime(),
       sysdatetimeoffset();
```

`sysutcdatetime()` 以 `timestamp without time zone` 返回 UTC。`sysdatetime()` 同样返回不带时区标记的本地挂钟时间，并使用会话 `TimeZone`。`sysdatetimeoffset()` 以 `timestamp with time zone` 返回当前时刻。比较前两个值时，必须考虑其隐含时区不同。

### Windows 计时器设置

在旧 Windows 系统上，超级用户可以请求更细的系统计时器间隔：

```conf
pg_sysdatetime.adjust_timer_resolution = on
```

该设置可以重载，也可由超级用户按会话修改。它会影响操作系统计时器行为与功耗，启用前应测量对全系统的影响。

该扩展不会调用 `GetSystemTimePreciseAsFileTime`。在 Windows 8 与 Windows Server 2012 之前，底层时钟可能仍然每毫秒才变化一次，因此显示更多数字不代表准确度更高。迁移时可用 PostgreSQL 原生时钟函数编写小型 SQL 包装器来保留这些名称，并验证应用所需的时区与易变性语义。
