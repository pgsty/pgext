## 用法

来源：

- [官方扩展 SQL](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/mdbqs--0.1.sql)
- [官方解析器实现](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/mdbqs_gram.y)
- [官方回归示例](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/sql/mdbqs_test.sql)

`mdbqs` 0.1 将 MongoDB 查询语法的一个子集解析为独立 PostgreSQL 扩展 `jsquery` 的 API，并对 JSONB 值求值。它不会连接 MongoDB。只有旧应用需要这一准确查询方言且能严格控制查询字符串时才应使用。

### 安装并匹配 JSONB

控制文件没有声明运行时依赖，所以必须先安装 `jsquery`：

```sql
CREATE EXTENSION jsquery;
CREATE EXTENSION mdbqs;

SELECT '{"a": 2}'::jsonb
       <=> '{ a: { $gte: 2 } }'::mdbquery;

SELECT '{"tags": ["ssl", "security"]}'::jsonb
       <=> '{ tags: { $all: ["ssl", "security"] } }'::mdbquery;
```

`<=>` 操作符为 `jsonb` 与 `mdbquery` 的两种参数顺序都提供定义，并返回布尔值。

### 支持范围

语法包括标量相等以及 `$eq`、`$ne`、`$lt`、`$lte`、`$gt`、`$gte`；数组的 `$in`、`$nin` 与 `$all`；`$size`、`$exists`、`$type`、`$not` 与 `$mod`；以及使用 `$and`、`$or`、`$nor` 的表达式树。解析树会转换成 `jsquery`，因此支持的值类型与匹配行为最终取决于已安装的 `jsquery` 版本。

### 安全与兼容性

这是与旧 PostgreSQL 及 `jsquery` 内部函数名耦合的解析器。多个 MongoDB 构造和类型会被拒绝，它不是完整的 MongoDB 兼容层。已复核解析器的错误处理器会在语法无效时调用 C 进程退出函数，可能终止当前后端连接，而不是返回普通解析错误。绝不能接收不可信用户的任意 `mdbquery` 文本；应在隔离进程中验证，并只绑定已知模板。扩展没有定义索引操作符类，因此应检查查询计划；除非显式提供其他兼容索引表达式，否则应预期逐行求值。需要固定 PostgreSQL 与 `jsquery`，并测试畸形输入、嵌套深度、数值转换、Unicode、NULL 或缺失字段以及后端失败恢复。
