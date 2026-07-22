## 用法

来源：

- [官方上游文档](https://github.com/nuko-yokohama/vmatch/blob/6fd255384005a84a3bd8e05dc2e099eb60fa1bdf/readme_jp.txt)
- [1.0 版本 SQL 定义](https://github.com/nuko-yokohama/vmatch/blob/6fd255384005a84a3bd8e05dc2e099eb60fa1bdf/vmatch--1.0.sql)
- [扩展控制文件](https://github.com/nuko-yokohama/vmatch/blob/6fd255384005a84a3bd8e05dc2e099eb60fa1bdf/vmatch.control)

`vmatch` 提供文本相似度函数与模糊布尔操作符 `/=`。其自定义评分会对可能的 QWERTY 键盘输入错误赋予特殊权重，它适用于小型近似匹配实验，而不是经过校准的搜索排序。

### 核心流程

```sql
CREATE EXTENSION vmatch;

SELECT data, get_similar_rate(data, 'postgresql') AS similarity
FROM candidate
WHERE data /= 'postgresql'
ORDER BY similarity DESC;
```

上游示例展示了不区分大小写的匹配，以及 `PostgresSQL`、`pastgres` 和 `pstgres` 等拼写变体。应在应用真正关心的语言与数据上验证分数分布。

### 对象与语义

- `get_similar_rate(text, text)` 返回 `real` 相似度分数。
- `vmatch(text, text)` 返回 `/=` 使用的布尔判定。
- 在版本 `1.0` 中，`/=` 使用固定阈值 0.75；没有文档化的 GUC 可以更改它，也没有索引操作符类。

上游明确把日文支持、改进公式和可配置阈值列为未完成工作，并称算法的实用性尚未验证。扩展被声明为 `STRICT` 和 `IMMUTABLE`，但这些声明不会让其模糊结果在语言学上正确。不要把 `/=` 用作等值、唯一性或授权测试；应明确地规范化文本，并在使用前测试多字节输入、重音、区域设置、长字符串、NULL、假阳性和假阴性。
