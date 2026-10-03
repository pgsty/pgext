## 用法

来源：

- [pg_volvec.control](https://github.com/yuuch/pg_volvec/blob/a2fc6dfcc4fc09595f2894de18ff5447c80af941/pg_volvec.control)
- [README.md](https://github.com/yuuch/pg_volvec/blob/a2fc6dfcc4fc09595f2894de18ff5447c80af941/README.md)
- [pg_volvec--1.0.sql](https://github.com/yuuch/pg_volvec/blob/a2fc6dfcc4fc09595f2894de18ff5447c80af941/pg_volvec--1.0.sql)
- [src/bridge/pg_volvec.c](https://github.com/yuuch/pg_volvec/blob/a2fc6dfcc4fc09595f2894de18ff5447c80af941/src/bridge/pg_volvec.c)

`pg_volvec` 通过 PostgreSQL 执行器钩子，将支持的计划子树交给列式引擎执行，同时保留 PostgreSQL 规划器。它是实验性执行器，使用 LLVM 支持表达式与元组解码。

### 核心用法

```sql
CREATE EXTENSION pg_volvec;
LOAD 'pg_volvec';
SET pg_volvec.enabled = on;
EXPLAIN SELECT sum(v) FROM (VALUES (1), (2), (3)) AS sample(v);
```

### 运行边界

README 要求具有指定上游修订且启用 LLVM JIT 的 PostgreSQL 17+，这不意味着支持所有主／次版本构建。需以超级用户安装和加载。安装 SQL 只在当前后端加载共享库，其他会话需自行加载或配置合适的预加载方式。`pg_volvec.enabled`、`pg_volvec.trace_hooks` 与 `pg_volvec.jit_deform` 控制实现。未支持的计划回退到 PostgreSQL，应使用代表性查询验证正确性和性能。历史脚本中的其他扩展名称不是当前启用方式。
