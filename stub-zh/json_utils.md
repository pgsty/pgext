## 用法

来源：

- [目录源码版本的官方 README](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/README.md)
- [目录源码版本的扩展控制文件](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/json_utils.control)
- [0.1 版扩展 SQL](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/json_utils--0.1.sql)
- [PGXN 发行页面](https://pgxn.org/dist/json_utils/)

json_utils 是面向 PostgreSQL 文本 JSON 类型的早期纯 SQL 扩展。它增加三个比较或包含风格的运算符和一个哈希运算符类，但其 README 已警告 0.1 版存在缺陷，而且实现比较的是序列化文本而不是 JSON 结构。

### 核心流程

安装后，这些运算符作用于 JSON 值和字符串。下列示例展示的是其字面行为，而不是语义化的 JSON 包含关系。

```sql
CREATE EXTENSION json_utils;

SELECT '{"a":1}'::json = '{"a":1}'::json;
SELECT '{"a":1,"b":2}'::json ? 'a'::varchar;
SELECT '{"a":1,"b":2}'::json @> '"b":2'::varchar;
```

### 已安装对象

- 相等函数把两个文档转换为字符串，再比较这些字符串。
- 键测试会查找由带引号键名和紧随其后的冒号组成的精确序列化模式。
- 包含风格测试执行不受约束的子串搜索。
- 哈希支持函数根据序列化文档的 MD5 文本导出整数，扩展还为 JSON 注册默认哈希运算符类。

### 安全性与兼容性

这些操作属于词法比较，因此空白、键顺序、格式和子串都可能在逻辑 JSON 值没有变化时改变结果，或产生误导性结果。包含风格运算符可能匹配字符串内部或嵌套内容，不能当作结构化包含。目录 SQL 把所有辅助函数都声明为易变函数。

当前 PostgreSQL 已提供结构化 JSONB 相等、包含、存在运算符和索引支持。新模式应优先使用这些内置 JSONB 能力。维护旧安装时，应使用恶意格式、重复键和嵌套键测试每个谓词，并在创建扩展前检查目标数据库中的运算符冲突。

该扩展可重定位；根据控制文件，安装不要求超级用户，也没有声明预加载或重启要求。
