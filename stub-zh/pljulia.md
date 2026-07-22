## 用法

来源：

- [官方 README](https://gitlab.com/pljulia/pljulia/-/blob/343b46ef81bfa97f76e6e0b9acbe06ed0dae5e24/README.md)
- [0.8 版扩展 SQL](https://gitlab.com/pljulia/pljulia/-/blob/343b46ef81bfa97f76e6e0b9acbe06ed0dae5e24/pljulia--0.8.sql)
- [官方回归 SQL](https://gitlab.com/pljulia/pljulia/-/tree/343b46ef81bfa97f76e6e0b9acbe06ed0dae5e24/sql)

`pljulia` 把 Julia 嵌入为非受信 PostgreSQL 过程语言。0.8 版支持普通函数、匿名代码块、常见标量/数组/复合值转换、集合返回函数、会话本地共享数据与 SPI 访问。代码在数据库后端内部运行，因此只有可信管理员才能定义 Julia 函数。

### 核心流程

针对确切 PostgreSQL 与 Julia 安装构建后，创建语言并定义函数：

```sql
CREATE EXTENSION pljulia;

CREATE FUNCTION julia_filter_nulls(x integer[])
RETURNS integer[]
AS $$
return filter!(el -> el != nothing, x)
$$ LANGUAGE pljulia;

SELECT julia_filter_nulls(ARRAY[1, 2, NULL, 4]);
```

SQL 空值映射为 Julia `nothing`。文档给出了布尔、整数、浮点、numeric、text、数组与复合值转换；复合值输入表示为字典，返回时可使用字典或位置元组。

### 集合、SPI 与会话状态

- `return_next` 从集合返回函数发出行。
- `spi_exec(query, limit)` 返回有界结果数组或受影响行数。
- `spi_exec(query)` 返回由 `spi_fetchrow` 消费的游标；未完整消费时应以 `spi_cursor_close` 关闭。
- `spi_prepare` 与 `spi_exec_prepared` 支持参数化计划。
- `GD` 是 Julia 字典，只在当前 PostgreSQL 会话生命周期内由 PL/Julia 调用共享。

```sql
CREATE FUNCTION julia_count_rows()
RETURNS bigint
AS $$
rows = spi_exec("SELECT count(*) AS n FROM app.events", 1)
return Int64(rows[1]["n"])
$$ LANGUAGE pljulia;
```

应优先使用预备参数，不要插值拼接 SQL 文本。SPI 工作属于调用者 PostgreSQL 事务。

### 实验性边界

README 将项目标为开发中。它记录了基本 `DO` 支持并描述触发器/事件触发器，但常规触发器章节仍要求使用单独分支，而当前 main 树已经包含触发器回归 SQL。应把触发器视为实验特性，验证确切签出修订，不能假定文档与源码提供相同特性集合。

控制文件把语言安装到 `pg_catalog`，不可重定位，也不要求预加载。语言创建时没有 `TRUSTED`，因此具有非受信过程语言的安全边界。Julia 库、内存分配、原生代码、崩溃与阻塞调用都与后端进程共享。

0.8 版要求 Julia 及其头文件/库与构建、运行主机匹配。应固定 Julia、PostgreSQL、PL/Julia、包与原生依赖版本，并测试后端生命周期、并行、内存增长、异常、编码、升级与转储恢复。不要向不受信角色开放函数创建或可变 Julia 包加载能力。
