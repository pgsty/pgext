## 用法

来源：

- [固定提交的上游 README](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/README)
- [固定提交的 1.0 版安装 SQL](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/pg_fuck_block--1.0.sql)
- [固定提交的 C 实现](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/pg_fuck_block.c)
- [固定提交的扩展 control 文件](https://github.com/sangli00/pg_get_page_tuple/blob/a4330098dfc5535e8f5e416658a17261cafff71e/pg_fuck_block.control)

`pg_fuck_block` `1.0` 暴露底层函数 `pg_get_page_tuple(regclass, bigint)`，用于检查一个 heap block 中的 tuple slot。它是基于 PostgreSQL 存储内部结构的旧式诊断实验，不是具备快照一致性的表读取接口。

### 核心流程

该函数返回匿名 `record`；分解输出时必须提供与目标关系一致的列定义列表：

```sql
CREATE EXTENSION pg_fuck_block;

CREATE TABLE page_demo (
  id bigint,
  note text,
  created_at timestamptz
);
INSERT INTO page_demo VALUES (1, 'first row', clock_timestamp());

SELECT *
FROM pg_get_page_tuple('page_demo'::regclass, 0)
  AS t(id bigint, note text, created_at timestamptz);
```

第二个参数是从零开始的 heap block 编号。block 超出关系当前大小时会抛出 `block number out of range`。

### 重要对象

- `pg_get_page_tuple(regclass, bigint) RETURNS SETOF record`：扫描 heap page 中的普通 line pointer，并返回实现没有判定为 `HEAPTUPLE_DEAD` 的 tuple。

### 运维说明

该函数的 C 代码没有检查表 `SELECT` 权限，而扩展函数默认允许 `PUBLIC` 执行。应撤销不可信角色的执行权限。结果不受普通 MVCC 快照约束：可能暴露 recently dead、in-progress 或其他不可见 tuple，同时跳过非普通 line pointer。

已复核的 2016 年实现依赖已经移除或变化的后端 API，在多次调用迭代完成前就释放关系锁与 buffer，并且没有用 `heap_getattr()` 的结果填充输出 null bitmap。因此，并发修改、包含 `NULL` 的行、TOAST 数据或现代 PostgreSQL 构建可能导致错误数据、内存安全故障甚至后端崩溃。只能在与已测试服务器源码完全一致的隔离副本上使用；生产诊断应优先采用仍在维护的 `pageinspect` 等工具。
