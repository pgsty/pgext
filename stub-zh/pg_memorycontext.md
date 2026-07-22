## 用法

来源：

- [固定提交的上游 README](https://github.com/jackgo73/pg_memorycontext/blob/ac9dc1d208f712dcdee2bb5dffada347148b4b80/README.md)
- [1.0 版安装 SQL](https://github.com/jackgo73/pg_memorycontext/blob/ac9dc1d208f712dcdee2bb5dffada347148b4b80/pg_memorycontext--1.0.sql)
- [固定提交的私有内存实现](https://github.com/jackgo73/pg_memorycontext/blob/ac9dc1d208f712dcdee2bb5dffada347148b4b80/pg_memorycontext.c)
- [正式 PGXN 发行](https://pgxn.org/dist/pg_memorycontext/)

pg_memorycontext 1.0 版按 context name 聚合当前 backend 的 memory context tree。其视图返回同名 context 数量与已分配 block 总字节数，并按总大小和数量排序。

### 配置与查询

上游要求预加载该库并重启，然后创建扩展：

```conf
shared_preload_libraries = 'pg_memorycontext'
```

```sql
CREATE EXTENSION pg_memorycontext;

SELECT memorycontextname, count, totalsize
FROM pg_memorycontext
ORDER BY totalsize DESC, count DESC;
```

结果只描述执行查询的当前 backend。它是一个时间点诊断，不能测量其他 session，也不提供历史保留。

尽管结果会动态变化，安装 SQL 仍把 `pg_memorycontext()` 声明为 `IMMUTABLE`；不能把这个声明当作缓存或规划器安全保证。此外，C 代码使用更宽的 `long` 累加大小，但 SQL 把 `totalsize` 暴露为 32 位 `integer`，因此总量超过 2 GiB 时可能发生转换错误，而不是返回结果。

### 不安全的兼容边界

源码复制私有 AllocSetContext 布局，把遍历到的每个 context 都转换为这种结构，并从 TopMemoryContext 沿私有 first-child/next-child 指针行走。PostgreSQL 存在多种 memory context 实现，其布局也会随服务器版本变化。上游笼统的 PostgreSQL 9.2+ 声明不能当作现代兼容保证；构建不匹配时可能误读内存或让 backend 崩溃。

安装 SQL 虽然撤销底层函数权限，却把诊断视图 SELECT 授予 PUBLIC。普通用户仍能通过视图调用并观察 backend 内部分配名称和大小。应撤销视图 grant，只开放给诊断角色。必须针对确切 PostgreSQL server header 重新构建，在可承受崩溃的测试环境中验证代表性负载；不要仅凭 README 声明就在生产中预加载这份 2018 年遗留代码。
