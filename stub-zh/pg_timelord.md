## 用法

来源：

- [官方 README](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/README.md)
- [扩展 SQL](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/pg_timelord--0.0.1.sql)
- [快照、工作进程与钩子实现](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/pg_timelord.c)
- [扩展控制文件](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/pg_timelord.control)

`pg_timelord` 是一个概念验证项目，会依据事务提交时间戳替换历史堆可见性快照。它通过阻止清理来刻意保留旧元组版本，并不适用于生产、时间点恢复或可靠的时态查询。

### 配置与核心流程

模块只能通过 `shared_preload_libraries` 加载。启用提交时间戳跟踪，配置保留工作进程使用的唯一数据库，然后重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_timelord'
track_commit_timestamp = on
pg_timelord.database = 'postgres'
pg_timelord.keep_xact = 200000000
```

```sql
CREATE EXTENSION pg_timelord;

SELECT * FROM pg_timelord_oldest_xact;
SET pg_timelord.ts = '2016-10-31 23:59:59 GMT';

SELECT * FROM public.example;

RESET pg_timelord.ts;
```

应在显式事务之外设置或重置 `pg_timelord.ts`。请求的时间戳不能早于具有可用提交时间戳的最老保留事务。

### 对象与配置

- `pg_timelord.ts` 是用户可设置的时间戳字符串，用于在当前会话启用历史可见性。
- `pg_timelord.database` 是 postmaster 参数，默认为 `postgres`，选择保留后台工作进程连接的数据库。
- `pg_timelord.keep_xact` 是超级用户参数，默认为 200,000,000，控制工作进程尝试阻止清理的事务数量。
- `pg_timelord_oldest_xact()` 返回最老的保留事务 ID。
- `pg_timelord_oldest_xact` 视图报告配置的数据库、保留的 `xmin` 及其提交时间戳。

### 严重运维风险

工作进程会维持旧事务视界，使 VACUUM 无法删除实验所需的版本。默认保留窗口极大，可能造成无界的表与索引膨胀、事务 ID 回卷压力、磁盘耗尽、性能下降和可用性损失。在任何一次性测试中，都要持续监控事务年龄、存储、autovacuum 与回卷风险。

历史可见性仅覆盖尚未删除、且事务提交时间戳仍可用的堆元组版本。它不会重建 DDL、序列、外部副作用、非堆状态，也不是事务一致的全库镜像。启用后，规划器会禁用位图扫描和仅索引扫描，钩子会尝试强制只读并拒绝大多数实用命令；这些概念验证检查不是安全边界。

实现使用过时的 PostgreSQL 内部快照 API，上游仍明确指出其行为需要验证。即使版权年份后来更新，版本 `0.0.1` 也应视为已废弃的实验代码。不要在有价值的集群加载；只能在匹配且隔离的 PostgreSQL 构建上复现，并准备完整备份及丢弃实例的方案。
