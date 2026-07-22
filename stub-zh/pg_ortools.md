## 用法

来源：

- [官方 pg_ortools README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/README.md)
- [扩展 control 文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/pg_ortools.control)
- [SQL API 实现](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/src/lib.rs)

`pg_ortools` 版本 `0.2.0` 用 SQL 定义混合整数与线性优化问题，并通过 HiGHS 求解。它适合有界指派、资源分配、可行性及目标优化工作负载，这些问题通常可表示为整数或布尔变量和线性约束。

### 核心流程

```sql
CREATE EXTENSION pg_ortools;

SELECT pgortools.create_problem('example');
SELECT pgortools.add_int_var('example', 'x', 0, 100);
SELECT pgortools.add_int_var('example', 'y', 0, 100);
SELECT pgortools.add_constraint('example', 'x + y <= 150');
SELECT pgortools.maximize('example', '2*x + 3*y');

SELECT pgortools.solve_sync('example');
SELECT pgortools.get_solution('example');
```

对于异步任务，`solve` 返回作业 ID；使用 `solve_status`、`cancel_solve`、`LISTEN pgortools_solve` 和 `get_solution` 管理作业。`solve_greedy` 返回第一个可行解，而 `solve_local`、`solve_with_strategy` 与 `solve_auto` 提供其他搜索策略。声明式函数 `solve_assignment` 和 `parse_assignment` 可构建常见的表驱动指派模型。

### 运维说明

扩展在 `pgortools` 中保存问题、变量、约束、作业和解。约束表达式支持线性比较，但官方 README 明确不支持 `!=`。默认求解时限由 `pg_ortools.solver_time_limit` 控制；上游把 worker 启用、轮询间隔和 worker 数据库描述为对重启敏感的服务器设置。依赖异步作业前应确认后台工作进程已经运行。control 文件不可重定位但不限定超级用户安装，安装 SQL 还向 public 授予其模式表的 DML 权限；多租户使用前应审查权限。
