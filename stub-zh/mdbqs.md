## 用法

来源：

- [已复核 commit 的 mdbqs 0.1 安装 SQL](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/mdbqs--0.1.sql)
- [已复核 commit 的 mdbqs 回归测试 SQL](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/sql/mdbqs_test.sql)
- [已复核 commit 的 mdbqs.control](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/mdbqs.control)

`mdbqs` 增加 `mdbquery` 类型和 `<=>` 运算符，用于让 JSONB 值匹配 MongoDB 风格的谓词。回归测试覆盖 `$lt`、`$gte` 和 `$in` 等比较与成员谓词，`$and` 和 `$or` 等逻辑形式，以及 `$exists` 与 `$type` 等检查。

### 让 JSONB 匹配谓词

```sql
CREATE EXTENSION jsquery;
CREATE EXTENSION mdbqs;

SELECT '{"a": 2}'::jsonb
       <=> '{ a: { $lt: 3 } }'::mdbquery AS matches;

SELECT '{"quantity": 5, "price": 100}'::jsonb
       <=> '{ $and: [ { quantity: { $lt: 20 } }, { price: 100 } ] }'::mdbquery
       AS matches;
```

安装 SQL 为两种参数顺序都定义了该运算符，因此 `mdbquery` 也可以位于 JSONB 值左侧。

### 注意事项

- 源码会调用 `jsquery` 功能，回归测试也先创建该扩展，但 `mdbqs.control` 没有声明这一依赖。应显式安装 `jsquery`。
- 仓库和包名是 MDBQT/`mdbqt`，而可安装扩展是 0.1 版 `mdbqs`。
- 上游没有发布 README、许可证、发行历史或现代版本兼容矩阵。C 源码使用其 2017 年代码库对应的 PostgreSQL 内部 API；使用前必须针对确切服务器版本验证或移植。
