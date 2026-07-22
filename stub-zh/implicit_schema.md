## 用法

来源：

- [官方 README 与警告](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/README.md)
- [0.1 版本安装 SQL](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema--0.1.sql)
- [事件触发器实现](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema.c)
- [扩展控制文件](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema.control)

`implicit_schema` 安装一个事件触发器：使用 schema 限定名创建表或视图时，若 schema 不存在便自动创建。它是面向会生成 schema 名称、却无法自行预置 schema 的软件所设计的狭窄兼容性补丁。

### 核心流程

```sql
CREATE EXTENSION implicit_schema;

CREATE TABLE fizz.fuzz (id integer);

SELECT nspname
FROM pg_namespace
WHERE nspname = 'fizz';
```

该扩展定义 `auto_create_schema()`，并在 `ddl_command_start` 上创建名为 `auto_create_schema` 的事件触发器。版本 `0.1` 只处理 `CREATE TABLE` 和 `CREATE VIEW`；其他对象类型不会获得这种行为。

### 运维与安全边界

上游明确称这个扩展是一个坏主意。它会情况隐蔽地把一次对象创建尝试变成额外的 `CREATE SCHEMA IF NOT EXISTS` 操作，从而在安装它的数据库内全局改变权限、审计、锁和错误预期。

已复核的 C 实现会把请求的 schema 名插入动态 SQL，却没有转义内嵌双引号。不要向不可信名称或用户暴露使用 schema 限定名创建对象的能力。应优先显式预置 schema；如果兼容性必须使用该扩展，应隔离部署，测试恶意标识符和并发 DDL，并审计每个可创建表或视图的角色。
