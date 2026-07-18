## 用法

来源：

- [1.0 版本 PL/pgSQL 实现](https://github.com/itarient/pg_calcpi/blob/72012875fb47239b69b0347cde5de272b477ef52/pg_calcpi--1.0.sql)
- [扩展 control 文件](https://github.com/itarient/pg_calcpi/blob/72012875fb47239b69b0347cde5de272b477ef52/pg_calcpi.control)
- [上游测试查询](https://github.com/itarient/pg_calcpi/blob/72012875fb47239b69b0347cde5de272b477ef52/test/find_subseq8_before_subseq8.sql)

`pg_calcpi` 使用 PL/pgSQL 实现圆周率数字的 spigot 算法。`calcpi(n)` 返回复合类型 `digit` 的集合，每行包含从一开始的序号与一位十进制数字。

```sql
CREATE EXTENSION pg_calcpi;
SELECT seq, val
FROM calcpi(20)
ORDER BY seq;
```

每次调用都会创建并反复更新会话本地工作表，因此运行时间会随请求位数迅速增长。这是教学代码，不是数值计算库，也不是共享数据库服务器上的高效负载。

函数会先执行未限定 schema 的 `DROP TABLE IF EXISTS calcpi_internal_data`，然后才创建临时表。如果调用者能够解析并拥有同名持久表，该调用可能将其删除。只能在 `search_path` 受控的隔离 schema 或数据库中使用，也不能向不受信角色授予执行权。
