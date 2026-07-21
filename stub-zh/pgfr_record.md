## 用法

来源：

- [pgfr_record v2.29.2 README](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_record/README.md)
- [pgfr_record control file](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_record/extension.control)
- [pg_flight_recorder v2.29.2 reference](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/REFERENCE.md)
- [v2.29.2 release notes](https://github.com/dventimisupabase/pg_flight_recorder/releases/tag/v2.29.2)

pgfr_record 是 PostgreSQL Flight Recorder 的记录部分。它定期采样 PostgreSQL 活动、等待、锁、复制、真空及相关健康数据，并将其存储在数据库内的有界缓冲区中，然后保留快照以供事故分析使用。当需要短生命周期的数据库条件能够存活足够长的时间以便后续诊断时，请使用此功能。

### 安装和启用记录

pgfr_record 需要 pg_cron：

    CREATE EXTENSION pg_cron;
    CREATE EXTENSION pgfr_record;
    SELECT pgfr_record.enable();

enable() 会安装并调度收集器任务。它还会报告配置警告；请务必查看这些警告，而不要将成功调用视为所有指标都被收集的证明。

### 检查记录器健康状况

    SELECT * FROM pgfr_record.health_check();
    SELECT * FROM pgfr_record.ring_buffer_health();
    SELECT * FROM pgfr_record.list_profiles();

在发生事故前使用 set_mode 或 apply_profile 选择所需的采集配置文件：

    SELECT pgfr_record.set_mode('normal');

可用的采集模式包括正常、轻量级、紧急和终止。配置文件名称和采样间隔可能会发生变化，因此请列出已安装的配置文件而不是硬编码未记录的名称。

### 记录数据索引

- deltas：累积 PostgreSQL 计数器的时间间隔变化。
- recent_activity 和 recent_waits：采样的会话和等待事件。
- recent_locks 和 recent_idle：锁和空闲会话观察结果。
- recent_replication 和 recent_vacuum：复制和维护状态。
- archiver_status：WAL 归档健康状况。
- snapshot 和 ring-buffer 表：用于 pgfr_analyze 的保留历史记录。

许多工作缓冲区是 UNLOGGED 的，以减少写入放大。它们在崩溃后不可用，并且不会像普通已记录的表那样进行复制；持久化的快照提供了更长生命周期的分析表面。

### 管理函数

- pgfr_record.enable()：创建或激活计划收集器。
- pgfr_record.disable()：停止计划采集。
- pgfr_record.health_check()：报告收集器和配置健康状况。
- pgfr_record.set_mode(...)：更改采集模式。
- pgfr_record.apply_profile(...)：应用预定义的配置文件。
- pgfr_record.list_profiles()：列出可用的配置文件。
- pgfr_record.ring_buffer_health()：检查容量和保留压力。
- pgfr_record.cleanup(...)：根据 API 删除保留的历史记录。

### 保留和开销

默认设计保持短周期环形缓冲区历史记录，并且更持久化的快照通常在安装的配置文件不同而有所变化，大约为7天到30天。请验证实际表大小、作业调度以及保留设置。

pgfr_record 创建约十个 pg_cron 任务。pg_cron.log_run 每天可能会生成数千行记录；当额外的审计跟踪不再需要时，请禁用该日志或清除 cron 历史记录。采样还会增加 SQL、存储和目录流量，因此请在目标工作负载上测量开销。

版本 2.29.2 处理无法更新 cron.job 的管理服务角色：任务仍然可以被调度，而可选的 nodename 正则化会在警告中跳过。

### 注意事项

- pg_stat_statements 可丰富多种分析但为可选项；当需要时，请单独启用并调整其大小。
- 采集无法重建从未采样的时间周期。在发生事故前请启用并验证记录器。
- UNLOGGED 缓冲区可能在崩溃恢复后被截断。
- 记录器表可以包含查询文本、角色名称、客户端数据和操作细节。请应用适当的权限和保留控制。
