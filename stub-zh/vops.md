## 用法

来源：

- [固定提交的上游架构与使用指南](https://github.com/postgrespro/vops/blob/8259b306c6d71b605fd0ac4e1c0b64846358e4c2/README.md)
- [固定提交的扩展控制文件](https://github.com/postgrespro/vops/blob/8259b306c6d71b605fd0ac4e1c0b64846358e4c2/vops.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/vops/)

`vops` 为分析型 PostgreSQL 负载提供向量化图块类型、运算符、聚合、投影和 FDW。`vops_int4`、`vops_float8`、`vops_bool` 等类型在每个图块中保存最多 64 个标量值，在保留普通 PostgreSQL 堆存储的同时降低逐行执行器开销。

应在启动时加载查询重写钩子；无法预加载时，指南也记录了显式调用 `vops_initialize()` 的方式：

```conf
shared_preload_libraries = 'vops'
```

创建并填充 VOPS 图块表后，向量谓词使用 `filter`、`betwixt` 和重载的布尔运算符 `&`：

```sql
CREATE EXTENSION vops;

SELECT sum(l_extendedprice * l_discount) AS revenue
FROM vops_lineitem
WHERE filter(
  betwixt(l_shipdate, '1996-01-01', '1997-01-01')
  & betwixt(l_discount, 0.08, 0.1)
  & (l_quantity < 24)
);
```

图块列并非标量列：空值掩码、类型转换、过滤、聚合、投影刷新和批量加载都有 VOPS 专用语义。已复核指南说明向量运算符不支持连接，运算符优先级也要求给复合向量谓词加括号。信任性能收益前，应将查询计划与结果同标量参照表对照。

固定提交的控制文件版本为 `1.1`，而 PGXN 列出了较新的 `2.0.1` 发行版。应固定一条源码和版本线，核实服务器兼容性与重写钩子行为，并把投影刷新一致性、崩溃恢复、并发更新和不受支持的 SQL 形状视为明确运行边界。
