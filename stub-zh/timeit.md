## 用法

来源：

- [官方 README](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/README.md)
- [基准调度 SQL](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/FUNCTIONS/measure.sql)
- [查询计时 SQL](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/FUNCTIONS/time_query.sql)
- [C 实现](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/timeit.c)
- [扩展控制文件](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/timeit.control)

`timeit` 以高分辨率墙上时间测量 PostgreSQL 内置 C 函数，并能在受支持的 x86/Linux 系统上测量 CPU 周期。它重复调用指定内部符号，不断把迭代次数加倍，并用线性回归估算单次调用成本。这是用于隔离测试服务器的侵入式诊断工具，不是应用计时 API。

### 测量内置函数

在允许任何其他角色连接前先限制模式权限，然后以参数的文本形式测量一个已知安全的内置符号：

```sql
CREATE EXTENSION timeit;

REVOKE ALL ON SCHEMA timeit FROM PUBLIC;
REVOKE EXECUTE ON ALL FUNCTIONS IN SCHEMA timeit FROM PUBLIC;
GRANT USAGE ON SCHEMA timeit TO performance_lab;
GRANT EXECUTE ON FUNCTION timeit.t(text, text[], integer, double precision, integer, interval, integer)
    TO performance_lab;

SELECT timeit.t('numeric_add', ARRAY['1.5', '2.5']);
SELECT timeit.t('pg_sleep', ARRAY['0.01'], 2, 0.99, 10, interval '1 second');
```

默认测量从一次调用开始并逐次加倍，直到最近样本达到要求的回归阈值，或检测到超时。超时只在一个批次完成后检查；它不能中断一次很慢或卡住的调用，总调用次数也可能快速增长。

### 主要对象

- `timeit.t(...)`：以易读文本返回估算的单次调用墙上时间。
- `timeit.c(...)`：以 `bigint` 返回估算的单次调用 CPU 周期。
- `timeit.measure(...)`：返回样本数组、R 平方、斜率、截距和最终迭代次数；`measure_type` 取 `time` 或 `cycles`。
- `timeit.eval(text, text[])`：调用内部函数一次并把结果作为文本返回。
- `timeit.time_query(text)`：规划并执行只读 SQL，报告规划/执行时间与 Linux 性能计数器。
- `timeit.pretty_time`、`timeit.round_to_sig_figs` 和 `timeit.compute_regression_metrics`：格式化与回归辅助函数。

`core_id` 参数默认为 `-1`，由操作系统调度后端；指定核心会调用 Linux CPU 亲和性接口。周期测量需要兼容 CPU，而 `timeit.time_query` 需要主机内核与容器策略允许 Linux `perf_event_open`。

### 安全与测量注意事项

C 代码通过 `fmgr_internal_function` 解析内置符号，使用函数声明的参数类型转换输入文本，再通过 PostgreSQL 函数管理器直接调用。它不走普通 SQL 名称解析，也没有对选中的内置函数做显式 `EXECUTE` 权限检查。扩展函数默认可由 `PUBLIC` 执行；必须撤销该权限，只把严格选择的包装器授予可信实验角色。

不要传入任意函数名。内置函数可能易变、休眠、大量分配内存、修改状态、依赖特殊调用上下文、使用伪类型，或在直接调用时失败。重复执行会放大所有副作用，并可能阻塞或崩溃后端。`timeit.time_query` 要求 SPI 只读执行，但只读 `SELECT` 仍可调用易变函数；只能接受已审查的查询。

测量会受到缓存状态、CPU 频率、中断、后端内存分配与插桩开销影响。固定核心会改变后端亲和性，性能计数器可能不可用或被复用，并发活动也会扭曲回归结果。请使用一次性数据库，有意识地预热，多次重复，检查返回的拟合，并把结果视为该主机上的对比观察，而非普适性能保证。
