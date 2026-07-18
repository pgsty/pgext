## 用法

来源：

- [固定提交的上游 README](https://github.com/jatobadb/pg_memorydump/blob/558ad20f2b2826607ec56b881bf3db49a79fc167/README.md)
- [1.0 版安装 SQL](https://github.com/jatobadb/pg_memorydump/blob/558ad20f2b2826607ec56b881bf3db49a79fc167/pg_contextdump--1.0.sql)
- [固定提交的内存上下文实现](https://github.com/jatobadb/pg_memorydump/blob/558ad20f2b2826607ec56b881bf3db49a79fc167/pg_contextdump.c)
- [固定提交的扩展控制文件](https://github.com/jatobadb/pg_memorydump/blob/558ad20f2b2826607ec56b881bf3db49a79fc167/pg_contextdump.control)

pg_contextdump 1.0 版暴露当前后端的 PostgreSQL 内存上下文树。其视图报告上下文名称/类型、临时生成的父/子 ID、分配参数、内存块数量、总字节数、空闲字节数和内存块大小直方图。

### 配置与查询

上游要求预加载该库并重启 PostgreSQL，之后再创建扩展：

```conf
shared_preload_libraries = 'pg_contextdump'
```

```sql
CREATE EXTENSION pg_contextdump;

SELECT contextname, contexttype, parent_id, totalsize, freespace
FROM pg_contextdump
ORDER BY parent_id, id;
```

输出只描述执行查询的当前后端，而不是集群中的全部会话。size 的单位为字节，ID 只为本次调用生成，并不是稳定的内存上下文标识符。

### 兼容性与信息暴露

该实现复制 PostgreSQL 私有内存上下文的结构布局，并通过类型转换遍历活动上下文树。上游只声明兼容 PostgreSQL 11.5，并称更新版本需要测试。这些内部结构会跨主版本甚至小版本变化；构建不匹配时可能返回无效数据或导致后端崩溃。必须针对服务器的确切源码版本重新构建并验证。

安装 SQL 虽然撤销了底层函数对 PUBLIC 的直接执行权限，却把视图的 SELECT 权限授予 PUBLIC。普通用户仍能通过视图看到后端内部名称和分配行为。应撤销该视图权限，只授予运维诊断角色。采集过程会在被测会话内分配内存并遍历树，因此也不应高频运行。
