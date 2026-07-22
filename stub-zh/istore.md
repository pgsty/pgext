## 用法

来源：

- [官方 README](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/README.md)
- [官方扩展 SQL](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/istore--0.1.12.sql)
- [官方扩展控制文件](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/istore.control)

`istore` 为 PostgreSQL 提供紧凑的整数键映射。`istore` 类型存储整数值，`bigistore` 存储 bigint 值；两者适用于以整数标识符为键的稀疏计数器、分布与分析聚合。

### 核心流程

文本输入类似 hstore，数组构造函数可以聚合重复键：

```sql
CREATE EXTENSION istore;

SELECT istore(ARRAY[1, 2, 1, 3, 2, 2]);
SELECT '1=>4,2=>5'::istore -> 1;
SELECT '1=>4,2=>5'::istore + '1=>4,3=>6'::istore;

CREATE INDEX metrics_keys_idx
  ON metrics USING gin (values istore_key_ops);
```

键查询使用 `->`；`?`、`?&` 和 `?|` 检查键是否存在。`||` 连接映射，算术操作符则合并匹配键。`istore_key_ops` GIN 操作符类可加速键存在谓词。

### 重要对象与注意事项

构造与转换工具包括 `istore(...)`、`bigistore(...)`、`akeys`、`avals`、`each`、`slice`、`compact`、`fill_gaps`、`accumulate`、`clamp` 和删除辅助函数。`sum_up` 汇总值，SUM 聚合返回 `bigistore`，为总计提供更宽的值类型。扩展也提供逐键最小值和最大值聚合。

从 `istore` 到 `bigistore` 的转换是隐式的；窄化赋值可能溢出。算术与除法仍需应用检查溢出和零除数。大型值可能进入 TOAST，GIN 索引也会增加写入和维护成本，因此应针对实际分布与更新负载进行基准测试，不能把上游示例当作性能保证。
