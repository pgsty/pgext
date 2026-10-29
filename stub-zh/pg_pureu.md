## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/README.md)
- [pg_pure.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pure.control)
- [pg_pureu.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pureu.control)
- [pg_pure--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pure--0.1.sql)
- [pg_pureu--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pureu--0.1.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/CHANGELOG.md)

`pg_pureu` 在 PostgreSQL 中执行非受信任的 Pure 标量函数，并提供 SPI 数据库访问能力。它属于 `pg_pure` 0.1.1 分发，SQL 版本为 0.1。只有超级用户才能用该语言创建函数，因为 Pure 可以通过 `extern` 调用任意原生函数。

### 核心流程

```sql
CREATE EXTENSION pg_pureu;
CREATE FUNCTION pure_factorial(n int) RETURNS numeric
LANGUAGE pg_pureu AS $$ foldl (*) 1L (1..n) $$;
SELECT pure_factorial(30);
```

扩展不可迁移，需要 PostgreSQL 17–19、适合嵌入的 libpure 分支及其前导库。`numeric` 在 Pure 中使用精确大整数或有理数；小数结果返回时按 PostgreSQL 的 numeric 显示精度舍入。机器整数在最终的范围检查之前就可能发生运算回绕，因此中间值可能增大时应使用大整数。

### 对象与数据库访问

语言安装调用、内联和验证处理函数。支持的 SQL 标量类型包括 smallint/integer/bigint、浮点、numeric、布尔与文本。SQL NULL 表示为 `sql::null`，文本返回值必须是 Pure 字符串。

`sql::query` 和 `sql::exec` 访问数据库，`sql::notice` 与 `sql::warning` 输出消息。数据库错误会成为带有 SQLSTATE 和消息的 `sql::error`，未捕获的异常会转为 SQL 错误。0.1 尚未实现参数化 SPI、集合返回函数、触发器、数组和复合类型；README 中有些示例明确属于路线图。不要把不可信值直接拼进 SQL 字符串。

### 配置与故障边界

配套语言同样使用 `pg_pure` 动态库。可选的 `shared_preload_libraries` 加载会在主进程中预热解释器，需要重启；否则各后端独立初始化。`pg_pure.using` 默认加载 `math`。已实现查询取消和栈溢出钩子，但通过 `eval` 建立的全局定义不具备事务性，回滚后仍然存在。

不要在同一后端混用 libpure 与 PostgreSQL JIT 的不同 LLVM 主版本：上游记录了后端退出时崩溃、影响整个集群的故障。其 0.1 受支持配置使用不带 LLVM JIT 的 PostgreSQL。受信任的 `pg_pure` 配套语言仍在等待沙箱，执行时会拒绝运行；安装它不会让当前语言适合不可信的函数作者。
