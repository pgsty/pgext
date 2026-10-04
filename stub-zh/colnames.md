## 用法

来源：

- [版本文档](https://github.com/theory/colnames/blob/v1.7.1/doc/colnames.md)
- [SQL 函数定义](https://github.com/theory/colnames/blob/v1.7.1/sql/colnames.sql)
- [C 实现](https://github.com/theory/colnames/blob/v1.7.1/src/colnames.c)
- [扩展控制文件](https://github.com/theory/colnames/blob/v1.7.1/colnames.control)
- [发行说明](https://github.com/theory/colnames/blob/v1.7.1/Changes)

`colnames` 提供 `colnames(record)`，以 `name[]` 返回输入记录的字段名。通用触发器或动态 SQL 辅助函数需要了解复合值的结构、又不想把字段值转换为 JSON 时，可以使用它。发行版本 `1.7.1` 保留了扩展版本 `1.7.0`。

### 核心用法

```sql
CREATE EXTENSION colnames;

SELECT colnames(ROW(1, 'foo', 458.0));
SELECT colnames(t)
FROM (SELECT 1 AS id, 'Ada'::text AS name) AS t;

CREATE TYPE contact AS (id integer, name text);
SELECT colnames(NULL::contact);
```

三个调用分别返回 `{f1,f2,f3}`、`{id,name}` 和 `{id,name}`。匿名行使用生成的名称，具名行保留属性名。带类型的空值可由复合类型提供描述信息；不带类型的空值无法确定复合类型。

### 函数行为

`colnames(record)` 返回 `name[]`，标记为 `STABLE`，并且有意不声明为 `STRICT`。它保留列顺序和带引号标识符的拼写，忽略已删除的列，对空复合类型返回空数组。它读取行描述信息，无须读取表数据。

### 运行说明

扩展没有运行时扩展依赖，无须预加载或重启。调用函数时才加载 C 库。控制文件未将扩展标记为可信，因此创建扩展需要超级用户。扩展可重定位，支持安装到指定模式，或通过 `ALTER EXTENSION colnames SET SCHEMA` 移动。

发行版 `1.7.1` 的文档声明兼容 PostgreSQL 8.2–17。PGSTY 软件包面向 PostgreSQL 14–18；PostgreSQL 18 的验收依据是软件包与回归测试，不是该发行版的兼容性声明。

PostgreSQL 9.3 及以后也可以结合 `row_to_json` 和 `json_object_keys` 使用纯 SQL 方案。若把取得的名称用于动态 SQL，必须正确引用标识符，并对值保持参数化。
