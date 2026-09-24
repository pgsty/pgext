## 用法

来源：

- [README v0.1.0](https://github.com/PG-Circuit/pg-circuit/blob/v0.1.0/README.md)

`pg_circuit` 在 PostgreSQL 16–18 中检查有风险的 DML 和 DDL。社区版可以观察、告警或阻止语句。先将其加入 `shared_preload_libraries` 并重启 PostgreSQL，再以超级用户创建扩展。

### 基础保护

```conf
shared_preload_libraries = 'pg_circuit'
```

```sql
CREATE EXTENSION pg_circuit;
CREATE TABLE circuit_demo (id integer);
INSERT INTO circuit_demo VALUES (1);
SET pg_circuit.mode = 'enforce';
DELETE FROM circuit_demo; -- blocked
DELETE FROM circuit_demo WHERE id = 1;
SELECT * FROM pg_circuit_status();
```

### 配置与诊断

`pg_circuit.mode` 默认为告警模式；观察模式只计算风险，强制模式阻止评分达到配置阈值的语句。`pg_circuit_runtime_state()` 报告压力信号，`pg_circuit_events()` 展示近期事件。

社区版的有效运行模式始终为 NORMAL，压力读数不会自动提升拦截力度。扩展是策略辅助工具，数据保护仍取决于应用事务、授权和备份。启用强制模式前应使用代表性查询检查规则和阈值。
