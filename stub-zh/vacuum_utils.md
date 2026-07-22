## 用法

来源：

- [官方 README](https://github.com/asotolongo/vacuum_utils/blob/8de352f1960ad0d32c9d0d8cf24fa7417ef6c057/README.md)
- [0.1.0 版本扩展 SQL](https://github.com/asotolongo/vacuum_utils/blob/8de352f1960ad0d32c9d0d8cf24fa7417ef6c057/vacuum_utils--0.1.0.sql)
- [PGXN 发行页](https://pgxn.org/dist/vacuum_utils/)

`vacuum_utils` 提供 SQL 辅助函数，用于估计自动清理/分析阈值、读取最近维护时间与次数、切换表级自动清理，以及运行 `ANALYZE`。规范函数名 `diff_to_autovaccum` 有意保留了上游拼写错误。

### 核心流程

```sql
CREATE EXTENSION vacuum_utils;

SELECT vacuum_utils.get_table_vacuum_threshold('public', 'customers');
SELECT vacuum_utils.diff_to_autovaccum('public', 'customers');
SELECT * FROM vacuum_utils.last_autovacuum_count('public', 'customers');

SELECT vacuum_utils.execute_analyze('public', 'customers');
```

检查辅助函数读取 `pg_stat_user_tables`、`pg_class.reltuples` 和当前全局自动清理设置。变更辅助函数以调用者权限执行 `ALTER TABLE` 或 `ANALYZE`。

### 对象索引

- `diff_to_autovaccum(schema, table)` 估算距离全局清理阈值还差多少死元组。
- `get_table_vacuum_threshold` 和 `get_table_analyze_threshold` 基于全局设置计算阈值。
- `last_vacuum_count`、`last_autovacuum_count`、`last_analyze_count` 和 `last_autoanalyze_count` 返回 `datetext_and_int` 复合值。
- `disable_autovacuum` 和 `enable_autovacuum` 修改表的 `autovacuum_enabled` 存储参数。
- `execute_analyze` 对指定表运行 `ANALYZE`。

### 运维说明

上游 README 明确说明版本 `0.1.0` 存在已知 bug。虽然控制文件声明可重定位，SQL 脚本实际创建并使用固定的 `vacuum_utils` 模式。阈值计算使用当前全局 GUC，没有考虑表级存储参数、基于插入的触发条件、事务 ID 回卷清理、分区行为或规划器估计陈旧；结果只能作为粗略诊断。

不要把变更函数授予不受信角色。其动态 SQL 直接拼接模式名与表名，没有引用标识符，存在语法和 SQL 注入风险；函数以调用者权限运行。只能使用管理员提供的简单标识符，另行确认目标表；生产自动化应优先使用原生系统目录查询以及安全引用的 `ALTER TABLE`/`ANALYZE` 语句。
