## 用法

来源：

- [固定提交的上游 README](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/README.md)
- [固定提交的 PostgreSQL 配置参考](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/doc/src/config-pg.md)
- [固定提交的函数指南](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/doc/src/use-plrust.md)
- [固定提交的扩展控制文件](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/plrust/plrust.control)
- [固定提交的 Rust 包元数据](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/plrust/Cargo.toml)

plrust 是一个过程语言处理器，会使用 Rust 和 pgrx 把每个 LANGUAGE plrust 函数编译成本机机器码。它通过 Rust API 暴露 PostgreSQL 数据类型、SPI 查询、集合返回函数和触发器。

### 必需配置与示例

审阅版本要求预加载，并提供可写的编译工作目录：

```conf
shared_preload_libraries = 'plrust'
plrust.work_dir = '/var/lib/postgresql/plrust-work'
```

```sql
CREATE EXTENSION plrust;

CREATE FUNCTION add_two_numbers(a numeric, b numeric)
RETURNS numeric
STRICT
LANGUAGE plrust
AS $$
    Ok(Some(a + b))
$$;

SELECT add_two_numbers(2, 2);
```

CREATE FUNCTION 会调用 Rust 工具链，可能耗时较长。控制文件中的扩展对象版本是 1.1，与本目录一致；固定提交的 Rust 包元数据版本则是 1.2.8。不要混淆 SQL 对象版本与 Rust 包版本。

### 信任与构建链边界

可信 PL/Rust 构建只面向受支持的 x86_64/aarch64 Linux 配置，并依赖受限的 postgrestd 环境和强制代码检查。其他构建属于非可信构建，不能提供相同的隔离。可信与非可信编译函数使用不同目标，不能透明混用。

编译器、链接器、工作目录、生成的共享对象、pgrx 绑定和 Rust 依赖都会成为数据库攻击面。若未设置 plrust.allowed_dependencies，审阅文档说明所有 Rust 包都可使用。能创建函数的用户可以消耗 CPU、磁盘、内存和依赖下载资源，也可能引入供应链代码。应使用显式且固定版本的依赖允许列表、不可变工具链、受限可写目录、无多余网络访问和严格的语言 CREATE 权限。不要弱化必需的代码检查；上游明确说明这会破坏可信运行的预期。
