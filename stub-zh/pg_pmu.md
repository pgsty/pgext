## 用法

来源：

- [官方资料](https://github.com/darthunix/pg_pmu/blob/fc7defb2a9cf7f287db0b216f29d17d362eab4c5/README.md)
- [官方资料](https://github.com/darthunix/pg_pmu/blob/fc7defb2a9cf7f287db0b216f29d17d362eab4c5/LICENSE)
- [官方资料](https://github.com/darthunix/pg_pmu/blob/fc7defb2a9cf7f287db0b216f29d17d362eab4c5/pg/pg_pmu.c)

pg_pmu 是面向 PostgreSQL 19/master 的无 SQL 扩展对象模块，模块身份声明版本 0.1.0。它在 Linux 和 macOS 上为 EXPLAIN ANALYZE 添加硬件指令数与周期计数，并包含实际参与执行的并行工作进程。

### 启用与使用

将库加入现有预加载列表并重启，无需 CREATE EXTENSION。PMU 要求 ANALYZE，这会真正执行语句；应使用只读示例，修改语句则需采用正常的执行保护措施。

```ini
shared_preload_libraries = 'pg_pmu'
```

```sql
EXPLAIN (ANALYZE, PMU, TIMING off) SELECT count(*) FROM pg_class;
EXPLAIN (ANALYZE, PMU, VERBOSE, FORMAT JSON, TIMING off)
SELECT count(*) FROM pg_class;
```

### 解释与限制

Linux 要求 kernel.perf_event_paranoid 不大于 2 或 CAP_PERFMON，容器与虚拟机仍可能拒绝计数器。无法取得或维持计数器时会让查询报错，而不返回似是而非的计数。Linux 用户态计数与包含内核的 macOS 计数不能直接比较。工作进程计数已包含在执行总量中，VERBOSE 可单独显示其贡献。

计数按语句汇总，不按计划节点拆分，并包含测量开销。缓存的预备计划没有规划计数，auto_explain 不在接口范围内。pg_pmu.max_nested 默认为 16，在服务启动时固定；槽位或硬件资源耗尽会报错。
