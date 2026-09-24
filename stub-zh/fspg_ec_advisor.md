## 用法

来源：

- [Official documentation](https://github.com/hqakhtar/pg_benchmark/blob/ea3d29f9c50ebf286a40606a9c2d2ef53ec98402/workload_analysis/fspg_ec_advisor/README.md)
- [Control file](https://github.com/hqakhtar/pg_benchmark/blob/ea3d29f9c50ebf286a40606a9c2d2ef53ec98402/workload_analysis/fspg_ec_advisor/fspg_ec_advisor.control)
- [Version 1.0 SQL](https://github.com/hqakhtar/pg_benchmark/blob/ea3d29f9c50ebf286a40606a9c2d2ef53ec98402/workload_analysis/fspg_ec_advisor/fspg_ec_advisor--1.0.sql)
- [Decision guide](https://github.com/hqakhtar/pg_benchmark/blob/ea3d29f9c50ebf286a40606a9c2d2ef53ec98402/workload_analysis/fspg_ec_advisor/docs/sql-decision-guide.md)

`fspg_ec_advisor` 1.0 是面向 PostgreSQL 17 的纯 SQL 工作负载顾问，保存测量基线和操作反馈，再给出资源配置与扩缩容建议。SQL 例程生成建议记录和待投递队列，不会自行调整基础设施规格。

### 启用与核心流程

在允许部署扩展文件的服务器上安装该扩展。配置并预加载 `pg_stat_statements`，修改预加载配置后重启，再先启用依赖扩展：

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION fspg_ec_advisor;
SELECT fspg_ec_advisor.capture();
SELECT * FROM fspg_ec_advisor.query_scores()
ORDER BY total_exec_ms DESC;
SELECT fspg_ec_advisor.latest_advice();
```

对象位于固定的 `fspg_ec_advisor` 模式。JSONB 采集结果与 `capture_history` 保存判断依据；`record_feedback()` 和 `recommendation_calibration` 将操作结果与建议关联。应在有意义的观察时间窗口内采样，不宜把单次瞬时结果直接当作扩缩容决策。

### 调度与集成

`run_advisory_cycle()` 采集数据、评估策略、更新建议事件并排队投递，可由外部调度器调用。可选的 `pg_cron` 支持通过 `schedule_advisory_cycle()` 及对应的取消调度例程提供；Citus 遥测也是可选项。策略保存在 `advisory_policy` 中，外部采集器可补充基础设施遥测。Azure Monitor 投递需要外部发送程序和已配置的云资源，不是在 SQL 扩展内嵌入云凭据。

### 权限与重置边界

控制文件不要求超级用户安装，但调用者仍需相应的数据库、模式权限，以及底层统计与操作权限。会话设置 `fspg_ec_advisor.reset_stats` 默认关闭，只有显式开启后，`capture_and_reset()` 才会在采集后重置语句统计；数据库、I/O 与检查点统计不会被重置。

上游目标版本为 PostgreSQL 17，扩展自身没有共享库，也不需要预加载。Azure 托管 PostgreSQL 并不自动允许任意扩展文件；只有通过服务允许的安装路径，才能在那里使用此源码包。
