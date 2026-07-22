## 用法

来源：

- [官方 README](https://github.com/begriffs/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/README.md)
- [扩展 SQL](https://github.com/begriffs/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/pg_decode_infomask--0.1.sql)
- [C 实现](https://github.com/begriffs/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/pg_decode_infomask.c)
- [扩展控制文件](https://github.com/begriffs/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/pg_decode_infomask.control)

`pg_decode_infomask` 把 PostgreSQL 9.6 堆元组的 `t_infomask` 与 `t_infomask2` 整数解码为具名布尔字段。它有助于检查页面级 MVCC 提示位和锁位，但本身不会读取堆页面，也不会判定事务状态。

### 核心流程

另行安装 `pageinspect` 以取得元组头值，然后把整数传给解码器：

```sql
CREATE EXTENSION pageinspect;
CREATE EXTENSION pg_decode_infomask;

SELECT h.lp, h.t_xmin, h.t_xmax, x.*
FROM heap_page_items(get_raw_page('public.demo', 0)) AS h
CROSS JOIN LATERAL pg_get_xact_infomask(h.t_infomask) AS x;
```

示例使用了 `pageinspect`，但它没有被声明为扩展依赖。其原始页面函数通常需要较高权限；请遵循该扩展自己的安全指南。

### 函数

- `pg_get_xact_infomask_bits(integer)` 返回 `xmin` 的已提交、无效和冻结位，以及 `xmax` 的已提交或无效位。
- `pg_get_xact_infomask(integer)` 额外返回 `xmin_status text[]` 与 `xmax_status text[]` 摘要。
- `pg_get_lock_infomask_bits(integer)` 暴露元组锁及 multixact 相关位测试。
- `pg_get_infomask2_bits(integer)` 返回属性数，以及键已更新、HOT 已更新和仅堆元组标志。

这些函数只解释提供的位掩码。提示位不是对 `pg_xact` 的实时查询，解码器不能证明当前提交状态、可见性或数据一致性。

### 版本与安全边界

SQL 文件明确面向 PostgreSQL 9.6 的内部位定义。堆元组标志与服务器 C API 都依赖主版本，因此不要用该构建解释其他 PostgreSQL 主版本的值。必须先构建并验证针对目标版本的移植，才能依赖输出。

C 入口没有声明为 `STRICT`，读取第一个参数时也不检查 SQL 空值。绝不要传入 `NULL`。源码还使用旧式双参数 `CreateTemplateTupleDesc` API，自 2017 年以来没有更新，无法原样在现代 PostgreSQL 版本上构建。它只适合受控取证，并应对照匹配版本的服务器源码交叉检查结果，不应被当作生产完整性检查器。
