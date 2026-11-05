## 用法

来源：

- [上游来源 1](https://github.com/secp256k1-sha256/adaptive_autovacuum/blob/6712f2bc62cd7c4c7cc2c264fc8977c7f23b111f/README.md)
- [上游来源 2](https://github.com/secp256k1-sha256/adaptive_autovacuum/blob/6712f2bc62cd7c4c7cc2c264fc8977c7f23b111f/adaptive_autovacuum.control)

`adaptive_autovacuum` 1.3.1 是处于 Beta 阶段的实例级控制器，用于自动清理设置、统计修复及紧急维护。

### 核心流程

先预加载库并重启，再仅在 `adaptive_autovacuum.control_database` 指定的控制库中安装一次，默认库为 postgres。健康检查函数可由 `pg_monitor` 使用。`adaptive_autovacuum.policy` 中的策略是表字段，不是 GUC。表级建议包含应用与撤销 SQL，此版本不会执行这些表设置变更。

```sql
CREATE EXTENSION adaptive_autovacuum;
SELECT * FROM adaptive_autovacuum.doctor();
SELECT * FROM adaptive_autovacuum.status();
SELECT * FROM adaptive_autovacuum.table_recommendations;
```

### 运行边界

支持 PostgreSQL 17/18，其中 17 的工作进程数量调整仅提供建议。新策略默认启用实际的全局调优，若仅观察，应在预加载前关闭实例开关，将策略设为 dry_run，再启用控制器。需预留后台工作进程槽位，并与 Patroni 或运维控制器协调。暂停或删除不会撤销已应用的全局设置，也不会取消全部运行中任务，移除前应恢复需要保留的配置。
