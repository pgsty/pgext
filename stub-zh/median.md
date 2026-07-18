## 用法

来源：

- [已复核 commit 的 Median README](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/README.md)
- [已复核 commit 的 median.control](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median.control)
- [已复核 commit 的聚合定义](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median--1.0.sql)
- [已复核 commit 的 Median 实现](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median.c)

`median` 定义了名为 `median(anyelement)` 的并行安全聚合。它忽略空值，使用输入类型的比较支持对其余值排序，然后返回中间值。输入数量为偶数时，它会计算两个中间值的平均值。

### 基本示例

```sql
CREATE EXTENSION median;

SELECT median(value)
FROM (VALUES
  (1.0::numeric),
  (9.0::numeric),
  (3.0::numeric),
  (7.0::numeric)
) AS sample(value);
```

该示例返回两个居中数值的中点。

### 注意事项

- 输入类型必须提供比较支持。偶数基数的输入还需要兼容的 `+` 和 `/` 操作符，因为实现会计算平均值。
- 聚合会在内存中保留并排序所有非空值，因此它不是常量内存的百分位数实现。
- control 文件与目录都标识版本 `1.0`。上游没有发布当前 PostgreSQL 兼容矩阵；请针对实际使用的服务器大版本和输入类型进行测试。
