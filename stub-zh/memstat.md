## 用法

来源：

- [Official README](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/README.md)
- [Extension SQL](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/memstat--1.0.sql)
- [Collector implementation](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/memstat.c)

memstat 报告 PostgreSQL 内存上下文的分配统计。无需预加载即可检查当前后端；在服务器启动时加载后，还能收集整个实例的逐后端汇总。它适合定向内存诊断，不适合作为永久高频监控。

### 核心流程

对于当前会话，创建扩展并检查其上下文树：

```sql
CREATE EXTENSION memstat;

SELECT name, level, totalspace, freespace
FROM local_memory_stats()
ORDER BY totalspace DESC;
```

如需查看所有后端，把 `memstat` 放在 `shared_preload_libraries` 最后一项，重启 PostgreSQL，再查询实例视图：

```ini
shared_preload_libraries = 'other_library,memstat'
memstat.period = 10s
```

```sql
SELECT * FROM memory_stats ORDER BY totalspace DESC;
```

### 重要对象

- `local_memory_stats` 返回当前后端的上下文名称、嵌套层级、块与空闲块计数、总字节及空闲字节。
- `instance_memory_stats` 增加后端 PID，并读取存活后端统计；它要求预加载。
- `memory_stats` 把实例数据汇总为逐后端摘要。
- `memstat.period` 设置两次实例统计收集之间的最小秒数。

### 运维说明

预加载钩子会在查询开始时收集内存统计，在繁忙服务器上可能代价较高。应调大 `memstat.period`、短时采样，并在诊断后移除模块。上游要求它位于预加载列表末尾，因为它会包装共享内存设置。上下文名称和布局属于 PostgreSQL 实现细节，可能随大版本变化。
