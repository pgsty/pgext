## 用法

来源：

- [官方 README](https://github.com/MasahikoSawada/pg_simula/blob/532efbc57ebd57de8308af8278e77581bb6be885/README.md)
- [官方扩展 SQL](https://github.com/MasahikoSawada/pg_simula/blob/532efbc57ebd57de8308af8278e77581bb6be885/pg_simula--1.0.sql)
- [官方 C 实现](https://github.com/MasahikoSawada/pg_simula/blob/532efbc57ebd57de8308af8278e77581bb6be885/pg_simula.c)

`pg_simula` 是用于 PostgreSQL 测试的故障注入扩展。它可以延迟指定命令标签、抛出 `ERROR`、`FATAL` 或 `PANIC`，也可以拒绝新连接。只能在可丢弃的测试系统中使用：部分动作会终止会话或使服务器崩溃。

### 设置与核心流程

该库会安装规划器、工具命令与连接钩子，因此必须预加载。将其加入 `shared_preload_libraries`，重启 PostgreSQL，然后创建扩展：

```sql
CREATE EXTENSION pg_simula;

SELECT add_simula_event('INSERT', 'WAIT', 10);
SELECT add_simula_event('TRUNCATE TABLE', 'ERROR', 0);

SET pg_simula.enabled = on;
INSERT INTO test_table VALUES (1);
TRUNCATE test_table;

SELECT clear_all_events();
```

规则保存在 `simula_events(operation, action, sec)` 中，`operation` 是 PostgreSQL 命令标签。`WAIT` 睡眠 `sec` 秒；`ERROR`、`FATAL` 与 `PANIC` 注入对应的服务器错误级别。`clear_all_events()` 删除全部规则。应优先使用管理函数，而不是直接修改表；后者可能递归触发模拟行为。

### 连接拒绝与安全

`pg_simula.enabled` 控制命令模拟。`pg_simula.connection_refuse` 会以认证错误拒绝新连接；启用前必须通过现有会话或配置重载准备并验证恢复路径。

上游仅报告在 Linux 与 PostgreSQL 10 上测试。钩子及内部 API 对版本敏感，不能据此推断支持更新版本。`PANIC` 规则可能迫使 PostgreSQL 进入恢复，`FATAL` 规则会断开会话。不要在生产环境中保留该扩展、其配置或事件表，并在每次测试后清除所有规则。
