## 用法

来源：

- [官方原生组件 control 文件](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/native/sql_saga_native.control)
- [官方原生 SQL 绑定](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/src/2n_temporal_merge_native.sql)
- [官方架构决策](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/doc/decision-log.md#decision-companion-rust-extension-for-performance-critical-paths-2026-02-22)
- [官方 Rust 实现](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/native/src/lib.rs)

`sql_saga_native` 是 `sql_saga` 的 Rust/pgrx 伴生库，用于性能关键的 `temporal_merge` 规划器和执行器内省缓存。它是无头组件，没有独立创建的扩展对象。主扩展 `sql_saga` 会登记由匹配的 `$libdir/sql_saga_native` 库实现的 SQL 函数。

### 核心流程

安装同时提供匹配 `sql_saga` 与 `sql_saga_native` 产物的软件包。主扩展安装后，可通过其登记的函数检查当前连接的原生缓存：

```sql
SELECT stat_name, stat_value
FROM sql_saga.temporal_merge_native_cache_stats();

SELECT sql_saga.temporal_merge_native_cache_reset();
```

普通应用调用 `sql_saga.temporal_merge(...)`，不直接调用原生规划器。主过程默认使用原生扫描线规划器。诊断时，可在当前会话改用 PL/pgSQL 规划器：

```sql
SET sql_saga.temporal_merge.use_plpgsql_planner = true;
```

### 原生接口

- `temporal_merge_plan_native(...)`：内部规划器，填充 `pg_temp.temporal_merge_plan` 并返回行数。
- `temporal_merge_executor_introspect(...)`：目标/源元数据与生成 DML 片段的内部单连接缓存。
- `temporal_merge_native_cache_stats()`：报告当前连接的规划器/执行器缓存条目、预备读取语句、命中与未命中。
- `temporal_merge_native_cache_reset()`：清除当前连接的原生缓存、预备语句和计数器。

规划器批量读取源与目标行，在 Rust 中用扫描线算法执行时间分段与合并，再生成 PL/pgSQL 执行器所消费的相同计划结构。缓存采用线程本地状态，在 PostgreSQL 中对应单个后端连接。

### 注意事项

`sql_saga_native` 没有独立 DDL 生命周期；其版本和 ABI 必须匹配命名 C 符号的 `sql_saga` SQL 绑定。不要创建单独扩展，也不要把内部规划器暴露为应用 API。主扩展与伴生库必须原子升级，然后在目标 PostgreSQL 18 构建上测试时间合并与缓存重置/统计。

缓存按连接隔离，不会跨会话共享。源字段名和类型参与缓存校验，但操作性 DDL 与角色变更仍应使用新会话测试。上游给出的性能倍数来自其自身基准，不是对其他模式或负载的保证。
