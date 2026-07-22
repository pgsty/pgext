## 用法

来源：

- [官方 README](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/README.md)
- [官方扩展控制文件](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/index_analyzer.control)
- [官方扩展 SQL](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/sql/index_analyzer--0.0.1.sql)

`index_analyzer` 是一组 PL/pgSQL 诊断函数，用于检查已有索引，以及发现引用表上缺少匹配索引的外键。其结果只是后续调查线索，不能直接作为自动删除或创建索引的决策。

### 核心流程

先对模式运行宽泛检查，再用以小数表示的选择率阈值检查具体表或索引。

```sql
CREATE EXTENSION index_analyzer;

SELECT analyze_tables('public', 0.05);
SELECT analyze_table('public.orders'::regclass::oid, 0.05);
SELECT analyze_index('public.orders_created_at_idx'::regclass::oid, 0.05);

SELECT analyze_fks('public');
SELECT analyze_fks('public.orders'::regclass::oid);
```

这些函数通过 notice 报告发现，而不是返回推荐结果表。在自动化中运行时应捕获客户端消息。

### 函数索引

- `analyze_tables`、`analyze_table` 与 `analyze_index` 负责在用户表和索引范围内组织检查。
- `analyze_index_selectivity` 根据系统目录统计估算不同值的选择率。
- `analyze_index_usage` 比较索引扫描与顺序扫描统计，并跳过包括主键在内的唯一索引。
- `analyze_index_count_distinct` 通过表采样估算不同值数量。
- `analyze_fks` 与 `analyze_fk` 可按模式、表 OID、约束名或约束 OID 检查外键索引。

### 解释与注意事项

上游明确说明这些答案并非定论。使用计数依赖工作负载历史；相关列会让目录估算高估选择率；采样检查可能代价很高，并且不擅长表达式索引。检查范围仅包含用户表。

该仓库已经归档。control 文件声明版本 `0.1.0`，但检入的安装脚本名为 `index_analyzer--0.0.1.sql`；安装前应确认打包产物已经处理这一上游不一致。对每个索引建议，都必须结合约束、低频关键查询和具有代表性的观察窗口复核。
