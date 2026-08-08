## 用法

来源：

- [pg_disorder 0.1.0 README](https://api.pgxn.org/src/pg_disorder/pg_disorder-0.1.0/README.md)
- [pg_disorder 0.1.0 元数据](https://api.pgxn.org/src/pg_disorder/pg_disorder-0.1.0/META.json)
- [pg_disorder 0.1.0 Makefile](https://api.pgxn.org/src/pg_disorder/pg_disorder-0.1.0/Makefile)

`pg_disorder` 是一个仅用于测试的 PostgreSQL 可加载模块，它会有意改变符合条件的 `SELECT` 查询输出顺序，用于发现无意中依赖未指定行顺序的应用和测试。它是一个无扩展对象的模块：没有控制文件、SQL 安装脚本，也不需要执行 `CREATE EXTENSION pg_disorder`。

### 为测试数据库启用

在会话启动时加载该模块，以便其规划器钩子可用：

```sql
ALTER DATABASE regression_db
  SET session_preload_libraries = 'pg_disorder';

ALTER DATABASE regression_db
  SET pg_disorder.mode = 'reverse';
```

修改 `session_preload_libraries` 后应重新连接。不要将此模块加入生产环境全局的 `shared_preload_libraries` 设置。

### 模式

```sql
SET pg_disorder.mode = 'off';
SET pg_disorder.mode = 'reverse';
SET pg_disorder.mode = 'shuffle';
SET pg_disorder.seed = 42;
SET pg_disorder.force_serial = on;
```

- `off` 不改变执行计划。
- `reverse` 以确定性方式反转符合条件的输出。
- `shuffle` 在会话种子、提交的查询文本和执行计划固定时产生确定性排列。使用默认种子零时，每个会话会先选择并记录一个随机种子。
- `force_serial` 禁止并行计划，使乱序测试能够复现。

修复失败查询时，应添加语义正确的 `ORDER BY`；不要编码在 `off` 模式下偶然观察到的顺序。

### 适用条件与注意事项

该钩子面向没有 `ORDER BY` 的顶层 `SELECT` 语句。它会有意跳过那些重新排序不安全或会改变 SQL 语义的查询形态，包括聚合、分组、`DISTINCT`、集合操作、窗口函数、递归查询、行锁，以及没有 `FROM` 关系的查询。

- `pg_disorder` 是故障注入工具，而不是生产查询功能。
- 乱序测试通过并不能证明每个无序查询都安全；被排除的查询形态和规划器路径不会被重写。
- 软件包仅安装服务器模块。应通过 GUC 或模块加载状态验证是否启用，而不是查看 `pg_extension`。
