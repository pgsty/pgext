## 用法

来源：

- [官方 README](https://github.com/dariubs/town/blob/3944f2f34bd509588907ec845398e18796939379/readme.md)
- [官方使用示例](https://github.com/dariubs/town/blob/3944f2f34bd509588907ec845398e18796939379/docs/example.md)
- [官方扩展 SQL](https://github.com/dariubs/town/blob/3944f2f34bd509588907ec845398e18796939379/sql/town--0.1.0.sql)

`town` 是一个小型 PL/pgSQL 辅助扩展，用于创建带有时间戳、标签数组和 JSONB 载荷列的约定式时序表。它适合需要这一固定布局的原型；它并不是时序数据库引擎，也不提供分区、保留策略、压缩、汇总或后台维护。

### 核心流程

安装 `town` 后，以表名调用 `create_town_table()`：

```sql
CREATE EXTENSION town;
SELECT create_town_table('metrics');

INSERT INTO metrics(ts, tags, data)
VALUES (now(), '{town,db,timeseries}', '{"val":21324}');

SELECT ts, tags, data->>'val' AS val
FROM metrics
WHERE 'town' = ANY(tags)
  AND ts BETWEEN timestamptz '2026-07-01' AND now();
```

生成的表包含 `id BIGSERIAL PRIMARY KEY`、`ts TIMESTAMP WITH TIME ZONE`、`tags TEXT[]` 和 `data JSONB DEFAULT '{}'::jsonb`。辅助函数还会在需要时安装 `btree_gist`，并为 `tags` 创建 GIN 索引、为 `ts` 创建 GiST 索引。

### 运维边界

调用 `create_town_table()` 的角色必须有足够权限在当前数据库和模式中创建扩展、表、序列及索引。表标识符会通过 `format('%I', ...)` 正确引用，但函数只接受最长 30 个字符的名称，且不接受带模式限定的名称。

虽然创建表使用了 `IF NOT EXISTS`，创建索引却没有，因此对已有表重复调用时可能因索引重名而失败。应用所需的约束、非空规则、额外索引、分区、保留策略和访问控制都需要自行添加；还应针对目标负载评估 GIN/GiST 的写入成本与 JSONB 结构。
