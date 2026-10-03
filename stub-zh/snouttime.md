## 用法

来源：

- [README.md](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/README.md)
- [snouttime.control](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/snouttime.control)
- [Cargo.toml](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/Cargo.toml)
- [src/series.sql](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/src/series.sql)
- [src/worker.rs](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/src/worker.rs)
- [src/jobs.sql](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/src/jobs.sql)
- [src/columnar/mod.rs](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/src/columnar/mod.rs)
- [sql/snouttime--0.1.5--0.1.6.sql](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/sql/snouttime--0.1.5--0.1.6.sql)

`snouttime` 0.1.6 为 PostgreSQL 17 和 18 提供时序分区、列式存储、汇总与时间关联查询。它属于早期源码版本；将已有数据交由其管理前，应评估存储与升级行为。

### 核心工作流

由超级用户安装 SQL 扩展，时序表管理操作需要相应的表所有权权限。此例从新表开始；转换已有表会改变表结构，并非只添加元数据标签。基本函数和手工执行的任务无需预加载。

```sql
CREATE EXTENSION snouttime;
CREATE TABLE metrics (ts timestamptz NOT NULL, value double precision);
SELECT snouttime.create_series('metrics', 'ts',
  partition_interval => interval '1 day');
INSERT INTO metrics VALUES (now(), 42);
SELECT snouttime.bucket('1 hour', ts) AS bucket, avg(value)
FROM metrics GROUP BY 1 ORDER BY 1;
SELECT snouttime.run_due_job();
```

### 分区与查询对象

`snouttime.create_series` 创建原生分区父表，并将已有行保留在默认分区中。时间列使用 `partition_interval`，整数时间使用 `partition_width`。已有约束和依赖可能限制转换；用于业务表前，应检查上游转换要求。

`snouttime.bucket` 按时间分桶；`snouttime.gapfill`、`snouttime.locf` 与 `snouttime.interpolate` 生成或填充缺失桶。`snouttime.first` 和 `snouttime.last` 按时间选取值，相同时间的取值次序不确定。`snouttime.counter_delta` 与 `snouttime.counter_rate` 要求输入按时间排序。分位数与去重计数草图均为近似值。`snouttime.asof_join` 选择右侧最近的匹配观测，`snouttime.window_join` 对有限时间窗口聚合；这些查询使用调用者权限执行。

### 任务、保留与汇总

要在服务器启动时运行任务，需将 `snouttime` 加入 `shared_preload_libraries`，把 `snouttime.databases` 配为逗号分隔的数据库列表，然后重启。`snouttime.interval` 默认为 10 秒。超级用户也可在未预加载时调用 `snouttime.start_worker()`，为当前数据库启动动态工作进程；该进程不会在服务器重启后自动恢复。手工运行器可反复调用 `snouttime.run_due_job()`，直到返回 false。任务以所管理表的所有者权限执行。

`snouttime.create_rollup` 创建结合物化桶与原始变更的视图，`snouttime.refresh_rollup` 可手工刷新。直接写入子分区会绕过变更跟踪。保留策略会删除旧分区，启用前应核对保留与备份政策。

### 列式存储与对象存储

`snouttime.seal` 将分区改写为 `snouttime_columnar` 访问方法，`snouttime.unseal` 将其还原为堆表。迟到写入进入内部增量存储。非唯一索引通常只覆盖这些后续写入；需要完整覆盖时，应使用文档中的 `keep_indexes` 选项。唯一索引与排斥约束仍覆盖全部行。封存分区不接受 BRIN 索引、`CREATE INDEX CONCURRENTLY` 与 `TABLESAMPLE`，并发更新可能需要客户端重试。

分层存储将封存数据移至兼容 S3 的对象存储。凭据 GUC 要求预加载，并由超级用户管理；服务端环境凭据是另一种上游支持的方式。`snouttime.tier_gc` 默认关闭，因为恢复出的数据库仍可能引用旧对象；启用垃圾回收前应考虑这些副本。

### 模式与升级

扩展不可作为可信扩展安装，也不可重定位。公开对象位于 `snouttime`，列式辅助表使用受保护的 `snouttime_internal` 模式。0.1.6 修复了非超级用户读取封存分区的问题。安装匹配的库与脚本后，通过 `ALTER EXTENSION snouttime UPDATE` 升级。上游不提供降级脚本，回退需要恢复升级前的备份。
