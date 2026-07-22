## 用法

来源：

- [官方 README](https://github.com/postgrespro/monq/blob/0c245caeeee1057144d91c57c2d7d229310afe7e/README.md)
- [0.1 版 SQL API](https://github.com/postgrespro/monq/blob/0c245caeeee1057144d91c57c2d7d229310afe7e/monq--0.1.sql)
- [扩展控制文件](https://github.com/postgrespro/monq/blob/0c245caeeee1057144d91c57c2d7d229310afe7e/monq.control)

`monq` 将 MongoDB 查询语法的一个子集转换为 JsQuery 表达式，并针对 PostgreSQL `jsonb` 求值。它提供 `mquery` 类型与可交换的 `<=>` 匹配操作符；这是兼容辅助工具，并非 MongoDB 服务端或完整的 MongoDB 查询语义实现。

### 核心流程

先安装必需的 `jsquery` 扩展，把查询文本转换为 `mquery`，再与 JSON 文档进行匹配。

```sql
CREATE EXTENSION jsquery;
CREATE EXTENSION monq;

SELECT '{"a":["ssl","security","pattern"]}'::jsonb <=>
       '{a: {$all: ["ssl", "security"]}}'::mquery;
```

SQL 中同时声明了 `jsonb <=> mquery` 与 `mquery <=> jsonb`。支持的类别包括比较操作符、逻辑操作符、`$exists`、`$type`、`$text`、`$all`、`$elemMatch` 与 `$size`。

### 兼容性限制

README 明确不支持 `$mod`、`$regex`、`$where`、位操作、地理空间、注释及投影操作符。结果受 JsQuery 行为限制，因此在建立针对性测试语料前，不要假定它与 MongoDB 的类型转换、字段缺失、数组、排序规则或空值语义一致。SQL API 没有为 `<=>` 定义索引操作符类；应在真实数据上确认执行计划，而不能假定谓词会使用索引。

项目自述仍处于开发中，并面向 PostgreSQL 9.4 时代的 API。应将 `monq` 与 `jsquery` 一起固定版本，验证不可信查询文本的解析行为，并在迁移前逐个将应用查询与预期 MongoDB 结果比较。
