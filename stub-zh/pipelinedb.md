## 用法

来源：

- [官方 README](https://github.com/pipelinedb/pipelinedb/blob/5cc6ef58ab5b1c84bb7b4e932f99bf5c347f46d8/README.md)
- [扩展控制文件](https://github.com/pipelinedb/pipelinedb/blob/5cc6ef58ab5b1c84bb7b4e932f99bf5c347f46d8/pipelinedb.control)
- [运行时配置](https://github.com/pipelinedb/pipelinedb/blob/5cc6ef58ab5b1c84bb7b4e932f99bf5c347f46d8/src/config.c)
- [匿名更新检查实现](https://github.com/pipelinedb/pipelinedb/blob/5cc6ef58ab5b1c84bb7b4e932f99bf5c347f46d8/src/update.c)

`pipelinedb` 1.1.0 会持续运行 SQL 查询，增量聚合流上的事件。连续视图保存聚合状态而不是原始输入行，适合高吞吐时序汇总，但其持久性、重放和查询语义与普通表不同。

### 启用运行时

模块必须通过 `shared_preload_libraries` 加载，并重启 PostgreSQL。这个已归档版本默认启用匿名更新检查；启动前应将其关闭，因为代码会通过明文 HTTP 向已失效的 `anonymous.pipelinedb.com` 服务发送安装及运行汇总数据。

```conf
shared_preload_libraries = 'pipelinedb'
pipelinedb.anonymous_update_checks = off
```

```sql
CREATE EXTENSION pipelinedb;
```

仓库最终代码只支持 64 位系统上的 PostgreSQL 10.1 到 10.5 和 PostgreSQL 11.0，并依赖 ZeroMQ。不要在未经测试的大版本上加载该二进制。

### 流与连续视图

```sql
CREATE FOREIGN TABLE event_stream (
    key integer,
    value integer
)
SERVER pipelinedb;

CREATE VIEW event_counts
WITH (action = materialize) AS
SELECT key, count(*) AS count
FROM event_stream
GROUP BY key;

INSERT INTO event_stream (key, value)
SELECT (random() * 10)::integer, g
FROM generate_series(1, 100000) AS g;

SELECT *
FROM event_counts
ORDER BY count DESC;
```

外部表是流入口，不是持久行存储。插入操作把事件发送给读取该流的连续查询；除非连续查询显式把所需信息物化到其他位置，原始行会被丢弃。

### 主要对象与控制

- 使用 `action = materialize` 创建的视图是带内部物化表的连续视图。
- 使用 `action = transform` 创建的视图可过滤或转换流，并通过 `pipelinedb.insert_into_stream(...)` 等函数发送输出。
- `pipelinedb.get_views()` 和 `pipelinedb.get_streams()` 查看已注册的连续对象。
- `pipelinedb.activate(text)`、`pipelinedb.deactivate(text)` 和 `pipelinedb.truncate_continuous_view(text)` 控制连续查询。
- `pipelinedb.stream_insert_level` 选择异步或同步确认行为；摄取前必须测试准确的丢失与延迟权衡。

不要直接修改内部物化表；`pipelinedb.matrels_writable` 默认关闭正是为了这一点。工作进程、合并器、队列、批次、提交间隔和内存 GUC 是需要一起规划的集群级容量参数。

### 生命周期与持久性

PipelineDB 团队已加入 Confluent，上游也宣布旧产品线不再发布新版本；仓库处于废弃状态。流接收、工作进程处理和物化聚合提交是三个不同的持久性节点。应在流之外保留可重放的事实来源，测试崩溃重启和备份恢复行为，监控 `pipelinedb.query_stats` 与 `pipelinedb.stream_stats`，并在任何存量系统采用前制定迁移计划。
