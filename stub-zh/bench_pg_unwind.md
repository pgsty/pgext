## 用法

来源：

- [扩展控制文件](https://github.com/JLockerman/bench_pg_unwind/blob/ff7c71b7c8935810521d7e58cd11d9643c41f1c3/bench_pg_unwind.control)
- [SQL API](https://github.com/JLockerman/bench_pg_unwind/blob/ff7c71b7c8935810521d7e58cd11d9643c41f1c3/sql/bench_pg_unwind.sql)
- [Rust 基准实现](https://github.com/JLockerman/bench_pg_unwind/blob/ff7c71b7c8935810521d7e58cd11d9643c41f1c3/src/lib.rs)
- [C 边界实现](https://github.com/JLockerman/bench_pg_unwind/blob/ff7c71b7c8935810521d7e58cd11d9643c41f1c3/src/bench.c)

bench_pg_unwind 0.1 是一个范围很窄的 2020 年微基准，用于比较带与不带 PostgreSQL 错误展开保护的重复 Rust 到 C 调用。它是诊断代码，不是应用扩展。

### 一次性基准

应在隔离服务器上使用较小的正迭代次数：

```sql
CREATE EXTENSION bench_pg_unwind;
SELECT bench_cross_fn(1000, 0, 1);
SELECT bench_try_fn(1000, 0, 1);
```

两个函数都返回重复加法结果。每次调用还会发出包含总耗时与单次迭代耗时的 WARNING，因此应将服务器和客户端消息作为基准输出收集。

### 注意事项

- 迭代次数直接控制后端中的紧密 CPU 循环，且没有上限。绝不能暴露给不可信输入。
- 零或负迭代次数会使计时代码用非正数除以持续时间，并可能进入 Rust 恐慌路径。
- 加法使用没有溢出检查的有符号 64 位 C 算术。应选择能使起始值、增量与迭代结果保持在范围内的参数。
- 结果包含 SQL 调用、日志、时钟、编译器与构建影响，不能作为 PostgreSQL 或 Rust 的通用性能指标。
- 构建依赖历史 Timescale Rust 支持分支与 2020 年时代的 PostgreSQL 头文件。上游没有测试、兼容矩阵、README 或完整的仓库许可证文件。
