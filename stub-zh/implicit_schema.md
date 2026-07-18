## 用法

来源：

- [上游行为说明与警告](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/README.md)
- [事件触发器实现](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema.c)
- [扩展控制文件](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema.control)

`implicit_schema` 安装一个事件触发器：使用 schema 限定名创建表或视图时，若 schema 不存在便自动创建。它原本是为会生成带时间戳 schema 名称的应用提供的兼容性补丁。

```sql
CREATE EXTENSION implicit_schema;
CREATE TABLE fizz.fuzz (id integer);
SELECT nspname FROM pg_namespace WHERE nspname = 'fizz';
```

上游把这种行为称为坏主意。它会悄悄把一次对象创建变成额外 DDL，从而改变权限、审计、锁和错误语义。已复核 C 代码直接用请求名称构造 `CREATE SCHEMA`，没有安全转义标识符中的引号。不要把它授权给不可信用户；应优先显式预建 schema，仅在完成恶意标识符与权限测试后使用。
