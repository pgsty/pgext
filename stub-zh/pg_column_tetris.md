## 用法

来源：

- [项目 README](https://github.com/rogerwelin/pg_column_tetris/blob/e70f9867c63e932cdaf87b2d34b6504adad9ce12/README.md)
- [扩展 control 文件](https://github.com/rogerwelin/pg_column_tetris/blob/e70f9867c63e932cdaf87b2d34b6504adad9ce12/pg_column_tetris.control)
- [0.1.0 版 SQL 实现](https://github.com/rogerwelin/pg_column_tetris/blob/e70f9867c63e932cdaf87b2d34b6504adad9ce12/pg_column_tetris--0.1.0.sql)

`pg_column_tetris` 0.1.0 是适用于 PostgreSQL 14 至 18 的纯 SQL 扩展。它通过事件触发器在 `CREATE TABLE` 后估算对齐填充，并可对低效列顺序发出警告或拒绝建表；还提供检查和重写建议函数。

### 检查并选择执行模式

默认模式为 `warn`；`strict` 会拒绝估算结果不佳的新表，`off` 则禁用事件触发器检查。

```sql
CREATE EXTENSION pg_column_tetris;

SELECT column_tetris.mode();
SELECT * FROM column_tetris.check('public.measurement'::regclass);
SELECT column_tetris.padding_wasted('public.measurement'::regclass);

SELECT column_tetris.set_mode('warn');
```

对不应检查的表使用 `column_tetris.exclude()`。临时表和系统表会被跳过，事件触发器检查建表操作，而不是每一次后续修改。

### 将估算与重写视为建议

估算器会模拟元组头和类型对齐，但无法完整预测 null 位图、变长或 TOAST 值、压缩以及工作负载特定行分布的实际存储。因此，报告的字节节省是设计信号，不是实测磁盘回收量。

`column_tetris.suggest_rewrite()` 返回迁移脚本；它不会保留每个外键、索引、触发器或默认值。生成的流程会重命名原表、创建并复制替代表，最后删除旧表，可能需要排他锁和停机。审查依赖对象、权限、identity 与序列行为、复制、回退并在真实的预演环境测试之前，绝不能执行该输出。列顺序还可能属于应用契约，例如按位置插入和行解码。
