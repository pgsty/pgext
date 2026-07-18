## 用法

来源：

- [已复核 commit 的 jsonb_deep_sum README](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/README.md)
- [已复核 commit 的 jsonb_deep_sum.control](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/jsonb_deep_sum.control)
- [PGXN 上的 jsonb_deep_sum](https://pgxn.org/dist/jsonb_deep_sum/)

`jsonb_deep_sum` 递归合并 JSONB 对象中的数值字段。它提供用于相加两个对象的 `jsonb_deep_add`，以及按行归约一列、同时保留嵌套对象结构的 `jsonb_deep_sum` 聚合。

### 相加与聚合 JSONB 值

```sql
CREATE EXTENSION jsonb_deep_sum;

SELECT jsonb_deep_add(
  '{"a": 1, "nested": {"b": 2}}'::jsonb,
  '{"a": 4, "nested": {"b": 3}}'::jsonb
);

CREATE TABLE measurements (data jsonb);
INSERT INTO measurements VALUES
  ('{"a": 1}'),
  ('{"a": 2, "b": 1}'),
  ('{"a": 5}');

SELECT jsonb_deep_sum(data) FROM measurements;
```

本例最后一个聚合会产生等价于 `{"a": 8, "b": 1}` 的对象。

### 注意事项

- 0.0.2 版支持 JSON 对象和数值。上游说明，JSON 值内部出现空值、字符串、数组或布尔值时会报错。
- 该算法遍历 JSONB 内部排序的树并执行归并；它没有定义数组合并或非数值冲突语义。
- control 文件与 PGXN 发行版都标识版本 0.0.2，与 catalog 版本一致。
