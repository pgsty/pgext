## 用法

来源：

- [扩展控制文件](https://github.com/tvondra/palloc_bench/blob/7f5c9caa357aa44c49b7af6167389aeaa54b71d3/palloc_bench.control)
- [扩展 SQL](https://github.com/tvondra/palloc_bench/blob/7f5c9caa357aa44c49b7af6167389aeaa54b71d3/palloc_bench--1.0.sql)
- [基准实现](https://github.com/tvondra/palloc_bench/blob/7f5c9caa357aa44c49b7af6167389aeaa54b71d3/palloc_bench.c)

`palloc_bench` 是一个用于 PostgreSQL 内存分配的微型开发者基准。它创建嵌套分配上下文，在最深层反复调用 `palloc` 与 `pfree`，然后通过警告报告墙钟耗时；它不会返回结构化基准结果。

### 核心流程

只应在隔离测试数据库中使用保守参数，并避免同时运行其他负载。

```sql
CREATE EXTENSION palloc_bench;

REVOKE ALL ON FUNCTION palloc_bench(integer, integer, integer) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION palloc_bench(integer, integer, integer) TO benchmark_admin;

SET client_min_messages = warning;
SELECT palloc_bench(8, 100000, 64);
```

三个参数依次是嵌套内存上下文数量、分配与释放迭代次数，以及分配字节数。函数返回 `void`；结果需要从 `WARNING: duration = ... ms` 消息读取。上下文构造不在计时循环内，每次迭代同时包含分配与释放。

### 安全与解读

实现没有验证负数或过大输入。大量上下文会消耗内存和栈，大量迭代会独占后端，负数或极大的分配尺寸可能转换成无效或巨大的分配请求。函数没有权限检查且默认向公众授予执行权，因此应从 `PUBLIC` 撤销，避免形成简单的资源耗尽入口。

计时器使用 `gettimeofday` 且不是单调时钟，结果还受 PostgreSQL 构建、分配器、CPU、调度器、缓存与并发负载影响。应在受控条件下重复运行，不能把一次警告当作生产性能保证。固定源码来自 2014 年并使用旧内存上下文 API，使用前必须针对确切 PostgreSQL 源码版本编译测试。不需要预加载。
