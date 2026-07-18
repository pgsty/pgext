## 用法

来源：

- [固定提交的上游 README 与示例](https://github.com/salva/plswipl/blob/769a10e0ab37ca0bfdd1da84c2632af9fc9fb365/README.md)
- [固定提交的扩展控制文件](https://github.com/salva/plswipl/blob/769a10e0ab37ca0bfdd1da84c2632af9fc9fb365/plswipl.control)

`plswipl` 把 SWI-Prolog 作为非可信 PostgreSQL 过程语言嵌入。函数体指定服务器端 Prolog 源文件，该语言调用文件中的谓词并产生 PostgreSQL 结果。

假设已复核的服务器端文件定义了谓词 `break_chars/2`，上游示例采用以下形式：

```sql
CREATE EXTENSION plswipl;

CREATE FUNCTION break_chars(text)
RETURNS SETOF text
AS $pl$"/srv/plswipl/example.prolog"$pl$
LANGUAGE plswipl;

SELECT * FROM break_chars('abc');
```

控制文件把扩展对象固定在 `pg_catalog`，声明扩展不可重定位、仅限超级用户安装，并创建非可信语言。指定文件会在数据库服务器安全上下文中读取；绝不能让不可信用户选择或修改该路径及其内容。

上游明确把版本 `0.1` 标为“勉强可用”的开发中版本。README 把数据转换、数组/JSON、`SPI`、数据库内模块、可信执行、模式/会话用户行为和异常处理列为未完成事项。没有完成原生代码审计、后端崩溃测试、严格文件系统隔离，以及未完成的验证器、触发器和错误路径修复前，不要部署到生产服务器。
