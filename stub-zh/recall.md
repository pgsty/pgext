## 用法

来源：

- [Official README](https://github.com/mreithub/pg_recall/blob/87126883d6272a324b259b83f96436a53980dbd1/README.md)
- [Version 0.9.6 extension SQL](https://github.com/mreithub/pg_recall/blob/87126883d6272a324b259b83f96436a53980dbd1/recall--0.9.6.sql)
- [Official worked example](https://github.com/mreithub/pg_recall/blob/87126883d6272a324b259b83f96436a53980dbd1/examples/blog.md)

recall 为选定的 PostgreSQL 表保留完整行的时态历史。触发器把每个版本及其时间范围复制到逐表日志，从而可以在配置的保留周期内执行时点查询。它是用户数据的安全网，不是备份或防篡改审计日志。

### 核心流程

扩展依赖 `btree_gist`，目标表必须有主键。启用历史后照常修改数据，再让 recall 创建只在当前会话存在的历史视图：

```sql
CREATE EXTENSION recall CASCADE;

CREATE TABLE account_setting (
  account_id bigint PRIMARY KEY,
  value text NOT NULL
);

SELECT recall.enable('account_setting', '6 months');

INSERT INTO account_setting VALUES (1, 'old');
UPDATE account_setting SET value = 'new' WHERE account_id = 1;

SELECT recall.at('account_setting', now() - interval '1 minute');
SELECT * FROM account_setting_past WHERE account_id = 1;
```

也可以用 `_log_time @> timestamp` 直接查询日志。

### 重要对象

- `recall.enable` 创建模板表和日志表、复制既有行并安装历史触发器。
- `recall.at` 为一个时间点创建带 `_past` 后缀的临时视图。
- `recall.cleanup` 和 `recall.cleanup_all` 删除早于配置周期的版本。
- `recall.disable` 移除历史机制，并永久删除该表保留的日志。
- `_log_time` 是 `tstzrange`；GiST 排斥约束防止同一主键的版本区间重叠。

### 运维说明

每次变更都会存储完整行，增加存储、WAL 和写入开销。清理需要手工安排。模式变更必须通过生成的模板表执行，而日志表不会继承普通约束或外键。有权限的用户可以直接修改日志，因此它不能作为防篡改证据。需要谨慎测试主键变更和现代 PostgreSQL 兼容性，并始终保留独立备份。
