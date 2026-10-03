## 用法

来源：

- [1.0.1 README](https://github.com/okbob/pgmeminfo/blob/VERSION_1_0_1/README.md)
- [Installation SQL](https://github.com/okbob/pgmeminfo/blob/VERSION_1_0_1/pgmeminfo--1.0.sql)
- [Control file](https://github.com/okbob/pgmeminfo/blob/VERSION_1_0_1/pgmeminfo.control)

`pgmeminfo` 报告当前 PostgreSQL 后端的内存分配器统计及内存上下文层级，不汇总整个集群的内存。

### 检查内存

由超级用户安装扩展，然后检查当前连接：

```sql
CREATE EXTENSION pgmeminfo;
SELECT * FROM pgmeminfo();
SELECT * FROM pgmeminfo_contexts();
SELECT * FROM pgmeminfo_contexts(deep => 1);
SELECT * FROM pgmeminfo_contexts(deep => -1, accum_mode => 'off');
```

`pgmeminfo()` 返回 `arena`、`uordblks`、`fordblks`、`keepcost` 等分配器计数。这些指标与内存分配器相关，不是进程 RSS，也不能用于估算集群空闲内存。

`pgmeminfo_contexts(deep, accum_mode)` 返回上下文名称、父节点、层级和字节计数。默认累积模式为 `all`；`off` 不累加后代上下文，`deep => -1` 取消层级深度限制。

### 运行与版本

不需要预加载或重启。上游发布版本 1.0.1 保持 SQL 扩展版本 1.0；不要把安装包版本 1.0.1 用作 SQL 升级目标。后端执行查询时，上下文名称和统计值可能变化。如果不希望普通用户查看内部上下文名称，应检查函数访问权限。
