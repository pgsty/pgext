## 用法

来源：

- [官方 README](https://github.com/glynastill/table_log_pl/blob/564007f69efb810d0dc2123a8080037810283fff/README.md)
- [官方扩展 SQL](https://github.com/glynastill/table_log_pl/blob/564007f69efb810d0dc2123a8080037810283fff/sql/table_log_pl--0.1.sql)
- [官方扩展控制文件](https://github.com/glynastill/table_log_pl/blob/564007f69efb810d0dc2123a8080037810283fff/table_log_pl.control)

`table_log_pl` 是旧 C 扩展 `table_log` 的 PL/pgSQL 替代实现。它把 INSERT、UPDATE 与 DELETE 的行映像记录到影子表，并可按时间点重建一个独立表；这是兼容性辅助工具，而不是防篡改审计系统。

### 核心流程

创建目标模式，然后为表初始化日志：

```sql
CREATE EXTENSION table_log_pl;

CREATE TABLE public.account (
  id bigint PRIMARY KEY,
  balance numeric
);
CREATE SCHEMA audit;

SELECT table_log_pl_init(
  5, 'public', 'account', 'audit', 'account_log'
);
```

`table_log_pl_init(integer, text, text, text, text)` 创建日志表，以及 AFTER INSERT/UPDATE/DELETE 行触发器。级别 3 记录行元数据，但不带事件 ID 或用户；级别 4 添加 `trigger_id`；级别 5 同时添加 `trigger_id` 与 `trigger_user`。

### 日志与恢复语义

影子表复制原始列，并添加 `trigger_mode`、`trigger_tuple`、`trigger_changed` 及取决于级别的可选字段。INSERT 记录新行，DELETE 记录旧行，UPDATE 同时记录旧行与新行映像。

`table_log_pl_restore_table(...)` 按时间点重建一个独立表。方法 0 从完整历史向前重放，方法 1 从当前源表向后推演。默认结果是临时表，函数不会覆盖原始表。

日志要求行结构完全兼容；后续列变更可能破坏触发器或恢复逻辑。TRUNCATE 与 DDL 不会被捕获。触发器会增加同步写入成本，能修改触发器或日志表的所有者也能改变历史。应保护审计模式，测试模式迁移与恢复流程；需要证据完整性时，应采用专用审计设计。
