## 用法

来源：

- [项目仓库](https://github.com/optionfactory/bgworker_segfault)
- [已复核 commit 的扩展 control 文件](https://github.com/optionfactory/bgworker_segfault/blob/5c239e488936efd3068c1679bc933693fed9dc42/bgworker_segfault.control)
- [Cargo 软件包与 PostgreSQL feature 矩阵](https://github.com/optionfactory/bgworker_segfault/blob/5c239e488936efd3068c1679bc933693fed9dc42/Cargo.toml)
- [后台工作进程测试源码](https://github.com/optionfactory/bgworker_segfault/blob/5c239e488936efd3068c1679bc933693fed9dc42/src/lib.rs)

`bgworker_segfault` 是一个最小化 pgrx 复现项目，而不是应用扩展。仓库描述表明，它用于演示动态后台工作进程分配耗尽 `max_worker_processes` 时发生的段错误。源码导出后台工作进程入口 `bgworker_sleep`，并包含一个比估算可用容量多分配一个工作进程的测试。

### 预期用途

只能在隔离测试集群中用该项目复现或调查这一 pgrx/PostgreSQL 故障模式。项目没有面向最终用户的 SQL API 文档；`bgworker_sleep` 是 C 后台工作进程入口，而不是 SQL 函数。

control 文件把该扩展标记为 `superuser`、不受信任且不可重定位。软件包版本为 `0.0.0`，使用 pgrx `0.11.0`，声明了 PostgreSQL 11 至 16 的构建 feature，并默认选择 PostgreSQL 13。

### 注意事项

- 该测试会故意把后台工作进程容量推到极限，并且用于复现服务器崩溃。绝不能在生产或共享集群上运行。
- 安装扩展不会提供可运营的后台工作进程服务；有意义的行为位于测试框架中。
- 源码没有 README、release、用户配置或 Cargo feature 列表以外的兼容性承诺。应把它视为冻结的诊断夹具，而不是受支持的 PostgreSQL 组件。
