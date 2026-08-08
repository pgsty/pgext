## 用法

来源：

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [PL/Ruby 语言参考](https://github.com/commandprompt/plruby/blob/v2.5.0/doc/plruby.md)
- [PL/Ruby 实用手册](https://github.com/commandprompt/plruby/blob/v2.5.0/doc/cookbook.md)
- [PL/Ruby v2.5.0 控制文件](https://github.com/commandprompt/plruby/blob/v2.5.0/plruby.control)
- [PL/Ruby 变更日志](https://github.com/commandprompt/plruby/blob/v2.5.0/CHANGELOG.md)

`plruby` 是由 Command Prompt 维护的过程语言扩展，可将 Ruby 3 嵌入 PostgreSQL。软件包发行版 2.5.0 安装的 SQL 扩展版本为 `2.5`。它支持标量函数和集合返回函数、触发器、事件触发器、过程、匿名 `DO` 块、SPI 查询、游标以及预备计划。

### 创建函数

```sql
CREATE EXTENSION plruby;

CREATE FUNCTION ruby_add(integer, integer)
RETURNS integer
LANGUAGE plruby
AS $$
  args[0] + args[1]
$$;

SELECT ruby_add(2, 3);
```

参数通过 `args` 暴露；Ruby 的最后一个表达式会成为 SQL 返回值。语言参考中记录了 PostgreSQL 标量、数组、复合类型和 record 的转换规则。

### 集合返回函数

使用 `return_next` 从集合返回函数中发出行：

```sql
CREATE FUNCTION ruby_series(integer)
RETURNS SETOF integer
LANGUAGE plruby
AS $$
  1.upto(args[0]) { |n| return_next(n) }
$$;

SELECT * FROM ruby_series(3);
```

### SPI 与数据库操作

PL/Ruby 提供 PostgreSQL 的服务器编程接口，用于执行 SQL、使用预备计划和游标。请通过参数传递 SQL 值，不要将其插入命令文本；当会话不再需要长期存活的游标或预备状态时，应将其释放。

在 PostgreSQL 允许 `COMMIT` 或 `ROLLBACK` 的场景中，过程可以使用文档所述的事务控制接口。函数和触发器仍受 PostgreSQL 常规事务限制约束。

### 触发器与会话状态

触发器函数通过 `$_TD` 接收触发器元数据，并返回 PL/Ruby 文档规定的行操作。它还支持事件触发器、匿名 `DO` 块、后端本地会话数据和共享数据。这些功能在数据库后端内部运行，因此异常、阻塞调用或内存泄漏都会直接影响该后端。

### 版本 2.5.0

- `bytea` 现在映射为原始且能安全包含 NUL 的 Ruby `String`，编码为 `ASCII-8BIT`，而不再映射为 PostgreSQL 十六进制文本。这是一项破坏性转换变更：请审查解析或构造 `\x...` 字符串的函数，并显式构建字节，例如使用 `Array#pack`。
- `$_SD` 新增函数级状态，在同一会话的多次调用之间持续存在，并在函数重新编译时重置。`$_SHARED` 仍是在 PL/Ruby 函数之间共享的会话级状态。
- `spi_colnames`、`spi_coltypes` 和 `spi_coltypmods` 可提供结果列元数据，`ltree_plruby` 则新增可选启用的 `ltree` 转换。
- 安装 2.5.0 共享库和 SQL 文件后，请在每个已经安装该扩展的数据库中运行 `ALTER EXTENSION plruby UPDATE`。

### 安全与要求

- `plruby` 是不受信任的语言。Ruby 3 没有安全的进程内沙箱，因此只有超级用户可以创建 PL/Ruby 函数，代码会以 PostgreSQL 服务器进程的操作系统权限执行。
- 请将所有 PL/Ruby 源码当作特权服务器代码进行审查。绝不能允许租户或普通应用角色提交任意 Ruby 代码。
- 上游 v2.5.0 支持 PostgreSQL 11-18 和 Ruby 3.x。当前 Pigsty 软件包面向 PostgreSQL 14-18。
- 不要求设置 `shared_preload_libraries`。服务器端库被替换后，现有会话必须重新连接，才能确保新的运行时已经生效。
- `jsonb_plruby`、`hstore_plruby` 和 `ltree_plruby` 是配套转换扩展。函数必须显式声明 `TRANSFORM FOR TYPE ...`，才能接收原生 Ruby 结构，而不是走常规 datum 包装器/转换路径。
