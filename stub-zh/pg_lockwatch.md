## 用法

来源：

- [README](https://github.com/kxtxr/pg_lockwatch/blob/93c2d377c855ab0a39d9b2742d5fb022769454fe/README.md)
- [Control file / 控制文件](https://github.com/kxtxr/pg_lockwatch/blob/93c2d377c855ab0a39d9b2742d5fb022769454fe/pg_lockwatch.control)
- [Cargo.toml](https://github.com/kxtxr/pg_lockwatch/blob/93c2d377c855ab0a39d9b2742d5fb022769454fe/Cargo.toml)
- [src/lib.rs](https://github.com/kxtxr/pg_lockwatch/blob/93c2d377c855ab0a39d9b2742d5fb022769454fe/src/lib.rs)

`pg_lockwatch` 采样阻塞后端和等待链，估算锁竞争风险并发出告警。审核的 0.1.0 源码仍属实验性实现，分数是监控信号，不具有经过验证的预测准确率保证。

### 核心工作流

将 `pg_lockwatch` 加入 `shared_preload_libraries`，通过 `lockwatch.database` 选择工作数据库并重启 PostgreSQL，再以超级用户在该库安装。

```sql
CREATE EXTENSION pg_lockwatch;
SELECT lockwatch_sample_now();
SELECT * FROM lockwatch_risks;
SELECT * FROM lockwatch_history;
LISTEN lockwatch_alert;
```

### 配置

`lockwatch.sample_interval_ms` 设置采样间隔，`lockwatch.risk_threshold` 设置告警阈值。评分权重包括 `lockwatch.weight_velocity`、`lockwatch.weight_duration`、`lockwatch.weight_lock_mode` 和 `lockwatch.weight_cascade_depth`。发出通知时也会写入历史表。上游发布流程的目标版本是 PostgreSQL 15–18。

### 限制

共享内存数组容量在编译时确定，提高 `lockwatch.max_tracked_blockers` 或 `lockwatch.history_window` 不能扩容。指纹对原始查询文本取哈希，不做查询归一化。持续时间基线来自首次观测到的阻塞者，`lockwatch_history.resolved_as` 不会自动补填。

上游要求在实际负载前验证真实集成。审核提交中未找到许可证文件或 Cargo 许可声明，不能因源码公开而推断可再分发。
