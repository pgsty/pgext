## 用法

来源：

- [AWS aurora_wait_report 文档](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora_wait_report.html)
- [AWS aurora_stat_system_waits 文档](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora_stat_system_waits.html)
- [AWS Aurora PostgreSQL 扩展版本矩阵](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`aurora_stat_utils` 是 AWS 提供的 Aurora PostgreSQL 扩展；其文档化的用户接口 `aurora_wait_report([time])` 用于测量一个时间区间内的实例级等待活动。它仅由云厂商提供，应安装在 Aurora PostgreSQL 上，而不是社区 PostgreSQL 服务端。

### 核心流程

```sql
CREATE EXTENSION aurora_stat_utils;

-- Use the default 10-second observation interval.
SELECT *
FROM aurora_wait_report();

-- Observe a specific 60-second interval.
SELECT *
FROM aurora_wait_report(60);
```

可选参数单位为秒，默认值是 10。调用先取得统计快照，等待指定区间，再取得第二个快照并返回差值。

### 结果与解释

报告返回 `type_name`、`event_name`、`waits`、`wait_time`、`ms_per_wait`、`waits_per_xact` 与 `ms_per_xact`。区间数据由 `aurora_stat_system_waits()` 和 `pg_stat_database` 统计计算得到。

结果覆盖当前连接的 DB 实例，并包含实例上所有数据库的活动；它不局限于当前会话或查询。在写入节点上，AWS 还会通过 notice 报告已提交事务数和 TPS。判定原因前，应把热点等待与工作负载、Performance Insights 及 PostgreSQL 活动关联起来。

`aurora_stat_system_waits()` 提供累计等待及微秒单位的等待时间，而 `aurora_wait_report([time])` 以毫秒报告区间等待时间。比较数值列时必须考虑单位和范围差异。

### 运维边界

该函数会有意占用调用后端直至观察区间结束。应使用专用诊断会话、限定时长，避免放入延迟敏感的应用事务或高频监控抓取。

实例重启会重置累计系统等待统计，短区间也可能有较大噪声。保存报告时应一并记录观察窗口、引擎和实例上下文。

AWS 当前在受支持的 Aurora PostgreSQL 版本中列出扩展版本 `1.0`，但实际可用性取决于确切 Aurora 引擎小版本和区域发布进度。自动化前应检查当前 AWS 扩展矩阵以及目标集群上的 `pg_available_extension_versions`。
