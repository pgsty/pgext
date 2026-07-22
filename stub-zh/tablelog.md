## 用法

来源：

- [官方 README](https://github.com/snaga/tablelog/blob/0981b7efe69a477c06cce53f0b2151d6b1dd38ad/README.md)
- [0.1 版 SQL 实现](https://github.com/snaga/tablelog/blob/0981b7efe69a477c06cce53f0b2151d6b1dd38ad/tablelog--0.1.sql)
- [扩展控制文件](https://github.com/snaga/tablelog/blob/0981b7efe69a477c06cce53f0b2151d6b1dd38ad/tablelog.control)

`tablelog` 0.1 是一个基于 PL/v8 触发器的 `INSERT`、`UPDATE` 与 `DELETE` 变更日志。它将变更列名和新旧值以文本数组写入共享的 `__table_logs__` 表，同时记录事务、用户、表、事件和键信息。

### 启用并检查日志

目标表必须具有主键或唯一约束。先安装 PL/v8，按模式名和表名启用目标，再检查中央日志。

```sql
CREATE EXTENSION plv8;
CREATE EXTENSION tablelog;

SELECT tablelog_enable_logging('public', 'account');

SELECT ts, txid, dbuser, schemaname, tablename,
       event, col_names, old_vals, new_vals, key_names, key_vals
FROM __table_logs__
WHERE schemaname = 'public' AND tablename = 'account'
ORDER BY ts;

SELECT tablelog_disable_logging('public', 'account');
```

安装的触发器是 `tablelog_logging_trigger`。新日志行以 `status = 0` 开始；其他状态含义由应用定义。UPDATE 行包含发生变化的列，而不保证是完整行镜像。

### 审计与重放边界

触发器写入参与源事务，因此回滚会同时删除业务变更与日志。所有者或特权用户可以禁用触发器、编辑日志或绕过它，且 `TRUNCATE` 不在文档所列事件中。这适合变更历史，但不是独立且防篡改的审计轨迹。

值会转换为文本数组，从而丢失原生类型身份，并可能受格式化、排序规则、时区、编码和扩展类型输出变化影响。不要把它们直接当作 SQL 重放；应保留模式/版本上下文，并验证引用、NULL、二进制、生成列、标识列及大值行为。PL/v8 会在每个写入后端中执行 JavaScript，而单一中央表可能成为写入、存储、清理和保留瓶颈。应限制日志访问、尽量排除密钥、衡量工作负载开销，并制定明确的归档与删除策略。
