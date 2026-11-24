## 用法

来源：

- [上游来源 1](https://github.com/secp256k1-sha256/adaptive_autovacuum/blob/v1.4.0/README.md)
- [上游来源 2](https://github.com/secp256k1-sha256/adaptive_autovacuum/blob/v1.4.0/adaptive_autovacuum.control)

`adaptive_autovacuum` 1.4.0 是实例级控制器，用于自动清理设置、统计修复及紧急维护。

### 核心流程

先预加载库并重启，再仅在 `adaptive_autovacuum.control_database` 指定的控制库中安装一次，默认库为 postgres。健康检查函数可由 `pg_monitor` 使用。`adaptive_autovacuum.policy` 中的策略是表字段，不是 GUC。表级建议包含应用与撤销 SQL，此版本不会执行这些表设置变更。

```sql
CREATE EXTENSION adaptive_autovacuum;
SELECT * FROM adaptive_autovacuum.doctor();
SELECT * FROM adaptive_autovacuum.status();
SELECT * FROM adaptive_autovacuum.table_recommendations;
```

### 运行边界

支持 PostgreSQL 17/18/19，其中 17 的工作进程数量调整仅提供建议。新策略默认启用实际的全局调优，若仅观察，应在预加载前关闭实例开关，将策略设为 dry_run，再启用控制器。需预留后台工作进程槽位，并与 Patroni 或运维控制器协调。暂停或删除不会撤销已应用的全局设置，也不会取消全部运行中任务，移除前应恢复需要保留的配置。

### 1.4.0 的维护变化

紧急接管现在会终止已确认停滞的防事务回卷自动清理工作进程，并立即排队安排替代维护。它不会接管人工发起或仍在推进的清理，也会排除 PostgreSQL 19 中进入 failsafe 模式的清理。启用前应审查这种干预策略。

工作进程在 `pg_upgrade` 的二进制升级模式中保持空闲。升级前先安装与新主版本匹配的库，升级后在控制库执行 `ALTER EXTENSION adaptive_autovacuum UPDATE`。从 1.1.0 升级需要重建扩展，会丢失旧策略和历史；应遵循对应版本的流程。
