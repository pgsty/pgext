## 用法

来源：

- [Babelfish扩展BABEL_5_4_0 README](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/BABEL_5_4_0/README.md)
- [安装指南](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/BABEL_5_4_0/INSTALLING.md.tmpl)
- [`babelfishpg_tsql`控制文件](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/BABEL_5_4_0/contrib/babelfishpg_tsql/babelfishpg_tsql.control.in)
- [Babelfish限制](https://babelfishpg.org/docs/limitations/limitations-of-babelfish/)
- [处理T-SQL](https://babelfishpg.org/docs/usage/handling-tsql/)

`babelfishpg_tsql`实现了T-SQL语言和SQL Server兼容的系统目录行为，这是Babelfish数据库的一个组成部分，而不是一个可以独立添加到标准PostgreSQL中的兼容层：完整的堆栈需要带有Babelfish补丁的PostgreSQL引擎加上公共、TDS和T-SQL扩展。

### 核心工作流程

为预加载配置TDS协议扩展并重启Babelfish服务器：

```conf
shared_preload_libraries = 'babelfishpg_tds'
```

使用`CASCADE`创建TDS扩展，以便安装其依赖项，包括`babelfishpg_tsql`。在初始化前选择迁移模式。

```sql
CREATE EXTENSION IF NOT EXISTS babelfishpg_tds CASCADE;

ALTER SYSTEM SET babelfishpg_tsql.database_name = 'babelfish_db';
ALTER SYSTEM SET babelfishpg_tsql.migration_mode = 'multi-db';

CALL sys.initialize_babelfish('babelfish_user');
```

按照安装指南重新加载配置后，SQL Server客户端连接到通常位于端口1433的TDS监听器，并在Babelfish创建的逻辑数据库中发出T-SQL语句。

### 组件和对象索引

- `babelfishpg_tsql`提供T-SQL解析器、过程语言、系统对象、兼容函数以及T-SQL配置变量。
- `babelfishpg_tds`提供表格数据流监听器，也是常规安装入口点。
- `babelfishpg_common`提供共享的数据类型和函数。它和`uuid-ossp`是`babelfishpg_tsql`的声明依赖项。
- `babelfishpg_money`提供与堆栈相关的货币相关兼容对象。
- `sys.initialize_babelfish(login_name)`为Babelfish目录和服务初始化登录。
- `sys.sp_babelfish_configure`控制文档化的兼容性开关。
- `babelfishpg_tsql.database_name`标识托管Babelfish的物理PostgreSQL数据库。
- `babelfishpg_tsql.migration_mode`选择`single-db`或`multi-db`逻辑数据库映射。

### 运行边界

安装需要超级用户权限和与扩展版本匹配的Babelfish构建。不要单独安装`babelfishpg_tsql`并期望获得TDS连接性。迁移模式是一个配置决策，在数据库初始化后不应更改。

Babelfish实现了大量但不完整的SQL Server功能集。在迁移前，请根据官方限制验证应用程序语法、数据类型、系统目录假设、驱动程序和开关设置。PostgreSQL和T-SQL连接可以观察不同的命名和事务语义。

从5.5.0到5.4.0的目录更改是官方`BABEL_5_4_0`发布线的一个版本修正，而不是新功能或自动降级过程的证据。
