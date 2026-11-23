## 用法

来源：

- [官方 README](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/README.md)
- [扩展控制文件](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/query_schema.control)
- [扩展 SQL](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/sql/query_schema--1.1.sql)
- [升级 SQL](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/sql/query_schema--1.0--1.1.sql)
- [实现源码](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/src/query-schema.c)
- [支持的构建目标](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/Dockerfile)

上游发行版 `1.1.0` 中的 `query_schema` SQL 版本为 `1.1`，使用 PostgreSQL 的解析器与分析器，将查询结果描述为 Effect 模式源文本。上游将其限定为本地开发工具；现有发行制品面向 Alpine 上的 PostgreSQL 18，提供 amd64 和 arm64 架构。

### 工作流程

以超级用户身份安装到开发数据库。模块不要求预加载或重启服务：

```sql
CREATE EXTENSION query_schema;
CREATE TEMP TABLE example_accounts (id integer NOT NULL, label text);
SELECT query_schema('SELECT id, label FROM example_accounts');
SELECT * FROM query_shape('SELECT id, label FROM example_accounts WHERE id = $1');
```

`query_schema(text)` 返回基于 `Schema.Struct` 的文本表达式。`query_shape(text)` 返回 `path`、`pg_type`、`nullable`、`schema`、`note` 列，包含推断出的参数类型及嵌套 JSON、数组路径。参数行的 `nullable` 值为空。

替换模块文件后，已有安装使用 `ALTER EXTENSION query_schema UPDATE TO '1.1';` 升级。升级 SQL 只变更记录的扩展版本，不修改函数定义。

### JSON 与客户端解码

版本 `1.1` 改进了 `LATERAL` 连接以及 `UNION`、`INTERSECT`、`EXCEPT` 等集合运算的可空性分析，并区分顶层二进制客户端值与 JSON 内部值。

JSON 内的 `int8` 表示为 JSON 数字。生成的模式使用 `Schema.Int` 和 `SchemaGetter` 将其解码为大整数；JSON 解析后超出 JavaScript 安全整数范围的值会被拒绝。若要保留完整范围，应在构造 JSON 前将 `int8` 转为 `::text`，对应模式会使用 `Schema.BigIntFromString`。生成的模式可能同时需要以下两个导入：

```typescript
import { Schema, SchemaGetter } from "effect"
```

`@effect/sql-pg` 无法以二进制格式解码所有 PostgreSQL 类型。对于 `money`、范围、复合类型或多维数组等不支持的列，`query_schema` 会发出警告，`query_shape` 会添加注释。可将这些值选择为 `::text`，或用 `to_json` 包裹，再复核生成的客户端模式。

### 分析边界

每次传入一条语句。实现接受 `SELECT`、`INSERT`、`UPDATE`、`DELETE`、`MERGE`；数据修改语句描述其 `RETURNING` 字段。实用命令和多条语句会被拒绝。函数分析语句，并对查询执行重写，不执行结果计划。

分析使用当前数据库对象和名称解析规则。外连接、条件、聚合、CTE、枚举及 JSON 构造函数都会影响输出。无法确定的可空性保守表示为可空。顶层类型映射遵循 `@effect/sql-pg`，嵌套 JSON 则遵循 PostgreSQL 的 JSON 转换，两种表示可能不同。

这是开发辅助工具，不能证明任意输入安全，也不能保证推断模式完整验证运行时结果。应结合真实客户端解码器复核生成模式和注释。SQL 函数采用 PostgreSQL 的常规执行授权；共享开发数据库前应限制访问。已检查的源码未声明许可证，复用授权尚未核实。
