## 用法

来源：

- [官方 README](https://github.com/Brick-Abode/pldotnet/blob/7edb597e296f69609bb80054202c731861e2d9d8/README.md)
- [官方安装指南](https://github.com/Brick-Abode/pldotnet/blob/7edb597e296f69609bb80054202c731861e2d9d8/INSTALL.md)
- [官方 0.9 版本扩展 SQL](https://github.com/Brick-Abode/pldotnet/blob/7edb597e296f69609bb80054202c731861e2d9d8/pldotnet--0.9.sql)

`pldotnet` 将 .NET 嵌入 PostgreSQL，并安装不受信任的过程语言 `plcsharp` 与 `plfsharp`。函数、过程、`DO` 块、触发器、集合返回函数、记录以及 SPI 数据库访问都可用 C# 或 F# 实现。

### 前置条件与核心流程

已复核的 0.9 指南要求 PostgreSQL 10 或更高版本、.NET 6 或更高版本，以及匹配的原生软件包和库。创建扩展前，应安装针对精确 PostgreSQL 大版本构建的软件包。

```sql
CREATE EXTENSION pldotnet;

CREATE FUNCTION add_one(value integer)
RETURNS integer
AS $$
    return value + 1;
$$ LANGUAGE plcsharp;

SELECT add_one(41);

CREATE FUNCTION double_value(value bigint)
RETURNS bigint
AS $$
    value * 2L
$$ LANGUAGE plfsharp;

SELECT double_value(21);
```

扩展 SQL 为两种语言创建调用处理器、内联处理器与验证器。SQL 参数和返回值会转换为兼容 .NET/Npgsql 的类型。项目还支持从外部程序集加载代码，以及由托管代码调用 SPI；具体功能签名与类型映射应参考官方测试。

### 安全与运行时边界

`plcsharp` 与 `plfsharp` 是不受信任的语言。托管代码在 PostgreSQL 后端进程内执行，可能消耗进程内存、阻塞会话、通过 .NET API 访问服务器文件系统或网络，或通过原生互操作使后端崩溃。只能允许超级用户创建这类函数，也不能让普通数据库角色替换受信函数体。

.NET 运行时与托管依赖会成为数据库服务器 ABI 和补丁维护面的一部分。必须使用精确的 PostgreSQL/.NET/软件包组合验证启动、函数编译、SPI 行为、可空值和类型转换、并行使用、升级与崩溃恢复。目录将 0.9 标为预览版；不得把上游基准描述视为运行保证。
