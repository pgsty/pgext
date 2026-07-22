## 用法

来源：

- [官方 README](https://github.com/claesjac/pg-json/blob/551e0067fe7731778362aa3eb265b810002ba9de/README.md)
- [扩展 SQL](https://github.com/claesjac/pg-json/blob/551e0067fe7731778362aa3eb265b810002ba9de/sql/jansson-json.sql)
- [回归示例](https://github.com/claesjac/pg-json/blob/551e0067fe7731778362aa3eb265b810002ba9de/test/sql/json.sql)

`jansson-json` 0.0.2 是一个历史性的 Jansson 后端 JSON 类型与操作符扩展，编写时间早于 PostgreSQL 内置 `json` 和 `jsonb` 类型。它的 SQL 会尝试创建一个名字正好为 `json` 的类型，因此与现代 PostgreSQL 冲突，应作为遗留代码研究，而不是作为当前 JSON 方案安装。

### 历史 API

在兼容且尚无内置 JSON 的服务器上，其原始流程如下：

```sql
CREATE EXTENSION "jansson-json";

SELECT '{"foo":{"answer":42}}'::json ~ 'foo.answer';
SELECT json_set_value('{"a":1}'::json, 'a', '2'::json);
SELECT '{"a":1}'::json || '{"b":2}'::json;
```

`json_get_value(json, text)` 和 `~` 操作符以文本形式提取路径。路径可遍历对象键和数组，例如 `foo.quax`、`foo[0].quax` 与 `[0]`。`json_set_value(json, text, json)` 替换值。`json_concat` 与 `||` 合并对象或数组；`json_equals`、`json_not_equals`、`=` 和 `!=` 比较解析后的 JSON 值，而不是原始文本格式。

### 兼容性边界

扩展依赖原生 Jansson 库，并为自定义 `json` 类型暴露 C 语言输入输出函数。PostgreSQL 从 9.2 开始内置同名类型，因此扩展安装 SQL 会在仍受支持的现代版本上发生冲突。它的 API 也不具备现代 `jsonb` 所预期的索引、包含、路径、验证和维护能力。

新应用应使用 PostgreSQL 内置 `jsonb`、提取操作符、下标、SQL/JSON 路径函数和 GIN 索引。不要轻率地重命名遗留类型或 SQL 对象：其已编译函数签名和操作符定义依赖原始类型标识与旧服务器内部接口。
