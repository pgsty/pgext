## 用法

来源：

- [PGXN 2.6.0 README](https://pgxn.org/dist/plphp/2.6.0/README.html)
- [PL/php 语言参考](https://pgxn.org/dist/plphp/2.6.0/doc/plphp.html)
- [plphp 控制文件](https://api.pgxn.org/src/plphp/plphp-2.6.0/plphp.control)
- [PL/php 2.6.0 变更日志](https://api.pgxn.org/src/plphp/plphp-2.6.0/CHANGELOG.md)

`plphp` 把 PHP 作为非可信 PostgreSQL 过程语言嵌入，用于函数、过程、触发器、事件触发器与匿名块。只有在确实需要 PHP API、并能接受代码以 PostgreSQL 服务器操作系统权限运行时才应使用；它不是沙箱语言。

### 核心流程

```sql
CREATE EXTENSION plphp;

CREATE FUNCTION php_square(integer)
RETURNS integer
LANGUAGE plphp
AS $$
    return $args[0] * $args[0];
$$;

SELECT php_square(12);
```

PL/php 2.6 支持使用非 ZTS embed SAPI 的 PHP 8.1 至 8.4，以及 PostgreSQL 11 至 18。由于该语言明确属于非可信语言，安装与创建都需要超级用户。

### 主要接口

- `spi_exec(...)`、`spi_prepare(...)`、`spi_exec_prepared(...)`、`spi_query(...)` 与游标辅助函数通过 SPI 访问 PostgreSQL。
- `return_next()` 从集合返回函数输出行。
- `spi_commit()` 与 `spi_rollback()` 在过程中提供事务控制。
- `subtransaction(...)` 创建可捕获的子事务边界。
- `$_TD` 暴露触发器与事件触发器上下文。
- `$_SHARED` 是会话全局存储，`$_SD` 是每个函数私有的会话存储。
- `plphp.on_init`、`plphp.start_proc` 与 `plphp_modules` 用于初始化会话解释器。

可选的 `jsonb_plphp`、`hstore_plphp` 与 `bytea_plphp` 扩展提供原生转换。函数必须声明 `TRANSFORM FOR TYPE ...` 才会使用转换；安装配套扩展不会自动改变所有 PL/php 函数。

### 安全与运维

PL/php 函数可以读写文件、打开网络连接、调用 PHP 能力，并影响 PostgreSQL 操作系统账户可访问的一切。函数所有权与创建权只能授予等同服务器管理员信任级别的角色。若不得不使用高权限包装器，应按常规加固 `SECURITY DEFINER` 与 `search_path`。

SPI 查询文本、动态标识符、会话初始化代码、加载模块与持久解释器状态都需要审查。长时间运行的 PHP 代码会阻塞 PostgreSQL 后端，`$_SHARED` 或 `$_SD` 中保留的内存会持续到会话结束。对不可信值应使用预备语句，限制资源用量，测试取消与错误路径，并在应用代码保留过多状态时回收会话。

2.6 新增每函数 `$_SD` 与二进制安全的 `bytea_plphp` 转换。PostgreSQL 记录的控制版本是 `2.6`，源码发行包版本是 `2.6.0`。
