## 用法

来源：

- [Official README](https://github.com/okbob/autoreindex/blob/363e5299bda99f9e5166c228fd82c8845bf77500/README.md)
- [Extension control file](https://github.com/okbob/autoreindex/blob/363e5299bda99f9e5166c228fd82c8845bf77500/autoreindex.control)
- [Reindex worker source](https://github.com/okbob/autoreindex/blob/363e5299bda99f9e5166c228fd82c8845bf77500/reindexdb.c)

autoreindex 是一个未完成的后台工作进程原型，原计划用于发现膨胀索引并重建它们。所核验的上游版本并不会真正执行重建索引命令：它只查找并记录候选项，随后执行一个带休眠的测试代码块。不要把它当作自动索引维护系统。

### 核心流程

这是无 SQL 对象的模块，并非 SQL 扩展。只能在隔离测试实例中把 `autoreindex` 加入 `shared_preload_libraries`，然后重启 PostgreSQL：

```ini
shared_preload_libraries = 'autoreindex'
```

控制进程会遍历可连接数据库并启动工作进程。当前工作进程会排除主键及系统索引、记录候选行，然后进入休眠；并不存在完整的 `CREATE INDEX CONCURRENTLY` 或等效维护路径。

### 重要对象

- `autoreindex.diagnostics` 控制诊断日志。
- `autoreindex.hide_context_log_messages` 控制是否隐藏上下文消息。
- `autoreindex.max_workers` 和 `autoreindex.max_workers_per_database` 限制后台工作进程数量。
- `autoreindex.naptime` 控制工作进程间隔。

### 运维说明

由于该库必须预加载，即便它不能提供有效的索引重建，安装也会改变 postmaster 的启动行为。这个 2016 年的原型有 control 文件，但没有扩展 SQL 脚本，并含有硬编码测试行为。不要用于生产环境；索引膨胀策略和并发重建应使用仍受维护的 PostgreSQL 运维工具。
