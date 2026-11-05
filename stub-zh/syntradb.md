## 用法

来源：

- [上游来源 1](https://github.com/fikrifirmanf/syntradb/blob/1c8e61c1e0a1573df8dace235946d5a887d34527/README.md)
- [上游来源 2](https://github.com/fikrifirmanf/syntradb/blob/1c8e61c1e0a1573df8dace235946d5a887d34527/crates/syntra-pg/syntradb.control)

`syntradb` 0.12.0 在 PostgreSQL 堆存储之外提供显式分析函数与持久化列段。

### 核心流程

`syntra_heap_int4_summary` 按调用者快照读取受支持的堆表列。段函数编码和扫描固定数据，捕获源与检查点需显式管理生命周期。`syntra_apply_prune_capture` 原子应用并清理捕获源。这些接口不会自动替换普通查询计划。

```sql
CREATE EXTENSION syntradb;
SELECT syntra_version();
SELECT rows, non_nulls, sum, min, max
FROM syntra_int4_summary(ARRAY[1, NULL, -2, 8]::int4[]);
```

### 运行边界

测试基线为 PostgreSQL 18.6，安装需要超级用户。这是原型实现，自动分析规划路径尚未完成。可选后台检查点维护需按文档配置预加载、重启及数据库列表。应遵守所有权、列类型和资源限制，不能将已存储的段视为自动保持最新的副本。
