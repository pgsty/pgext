## 用法

来源：

- [官方 README](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/README.md)
- [扩展 SQL API](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/jsonb_deep_sum--0.0.2.sql)
- [官方回归测试](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/sql/jsonb_deep_sum_test.sql)

`jsonb_deep_sum` 0.0.2 递归相加 JSON 对象中的数值叶子。它提供双值函数和聚合，适合对不同行键集合不同的稀疏嵌套指标映射求和。

### 核心流程

```sql
CREATE EXTENSION jsonb_deep_sum;

SELECT jsonb_deep_add(
  '{"requests":{"ok":4},"bytes":10}'::jsonb,
  '{"requests":{"ok":3,"failed":1},"bytes":5}'::jsonb
);

CREATE TABLE metric_batch (metrics jsonb);
INSERT INTO metric_batch VALUES
  ('{"a":1,"nested":{"x":2}}'),
  ('{"a":4,"nested":{"x":3,"y":8}}'),
  (NULL);

SELECT jsonb_deep_sum(metrics) FROM metric_batch;
```

`jsonb_deep_add(jsonb, jsonb) RETURNS jsonb` 递归合并对象，并在相同路径相加数值。只存在于一侧的键会被保留。`jsonb_deep_sum(jsonb)` 使用该函数作为转换步骤，并以 `{}` 为初始状态；SQL NULL 输入行会被跳过。

### 数据契约与限制

遍历到的每个 JSON 值都必须是对象或数字。字符串、布尔值、JSON null 与数组都会报错，同一键上的不兼容结构也不会自动转换。聚合前应验证文档，并确保每条重复路径使用一致单位。

实现会在 C 中直接遍历 PostgreSQL 内部 `jsonb` 表示。上游示例测试过 PostgreSQL 12 时代的构建，但没有发布当前大版本矩阵；应针对精确服务器重新构建并运行回归测试。聚合会在后端内存中物化完整结果对象，因此高键基数或深层嵌套文档可能成本很高。对于模式稳定的指标，带类型列通常能提供更清晰的验证、索引和溢出处理。
