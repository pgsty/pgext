## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/postgrespro/aqo/blob/4a7fcd7291a8c4209c9102d1736f9f754d311cf2/README.md)
- [扩展控制文件](https://github.com/postgrespro/aqo/blob/4a7fcd7291a8c4209c9102d1736f9f754d311cf2/aqo.control)

`aqo` 使用先前执行的规范化查询形状训练模型，增强 PostgreSQL 基数估算。它面向因错误行数估算而选择差计划的昂贵查询；这是实验性优化器研究，并不保证加速。

所复核上游要求应用 PostgreSQL 版本专用核心补丁并完整重编服务器，同时进行集群级预加载：

```conf
shared_preload_libraries = 'aqo'
```

重启 PostgreSQL、创建扩展，并只在受控事务中训练选定查询：

```sql
CREATE EXTENSION aqo;
BEGIN;
SET aqo.mode = 'learn';
EXPLAIN ANALYZE SELECT * FROM fact JOIN dim USING (dim_id) WHERE dim.kind = 'x';
SET aqo.mode = 'controlled';
COMMIT;
```

`controlled` 是上游建议的生产默认值，会忽略未知查询形状；`learn` 记录查询；`intelligent` 自动调优；`forced` 共用模型；`disabled` 绕过 AQO。可检查 `aqo_queries` 和 `aqo_query_texts`，而 `aqo.show_details` 与 `aqo.show_hash` 只应用于诊断。

训练会真实执行查询，可能增加 CPU、内存、存储和延迟；数据分布变化后计划也可能退化。AQO 不支持临时对象，副本不能收集训练统计。必须固定精确的补丁服务器分支，用代表性参数做基准，保留快速禁用路径，并在集群级使用前测试故障切换、升级、模型清理和备份恢复。
