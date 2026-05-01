## 用法

来源：[README](https://github.com/danolivo/pg_pathcheck/blob/main/README.md)、[0.9.1 PG17/18 release](https://github.com/danolivo/pg_pathcheck/releases/tag/v0.9.1_pg17-18)、[PGXN metadata](https://api.pgxn.org/dist/pg_pathcheck.json)、[source](https://api.pgxn.org/src/pg_pathcheck/pg_pathcheck-0.9.1/pg_pathcheck.c)

`pg_pathcheck` 是 PostgreSQL 规划器诊断模块，用于验证可达的规划器 `Path` 树，并报告疑似已释放、损坏，或被错误关系复用的指针。它只能作为预加载共享库使用：注册 planner hook 和自定义 GUC，但不创建 SQL 函数、操作符、视图或表。

### 加载模块

针对要测试的 PostgreSQL 源码线构建 `pg_pathcheck`。上游 README 记录了 PostgreSQL 17/18 与 PostgreSQL master/19devel 的独立长期分支；版本 `0.9.1` 发布于 PG17/18 分支，PGXN 元数据也描述了 0.9.1 源码包。

将模块加入 `shared_preload_libraries` 并重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_pathcheck'
```

不需要执行 `CREATE EXTENSION pg_pathcheck`。预加载后，运行普通 SQL、`EXPLAIN`、回归测试或 PostgreSQL 测试套件即可；`pg_pathcheck` 会在查询规划时检查 planner path。

单个临时集群示例：

```bash
initdb -D pgdata
echo "shared_preload_libraries = 'pg_pathcheck'" >> pgdata/postgresql.conf
pg_ctl -D pgdata -l pgdata/logfile start

psql -c 'EXPLAIN SELECT ...'
```

对于 `make check-world` 生成的 PostgreSQL 测试集群，使用 `TEMP_CONFIG`：

```bash
cat > /tmp/pg_pathcheck.conf <<'EOF'
shared_preload_libraries = 'pg_pathcheck'
EOF

TEMP_CONFIG=/tmp/pg_pathcheck.conf make check-world
```

### 配置

`pg_pathcheck.elevel` 控制检测到异常 `Path` 时使用的严重级别：

```sql
SET pg_pathcheck.elevel = 'log';
SET pg_pathcheck.elevel = 'warning';  -- default
SET pg_pathcheck.elevel = 'error';
SET pg_pathcheck.elevel = 'panic';
```

使用 `warning` 或 `log` 可在测试继续运行时扩大覆盖；使用 `error` 可在第一条异常查询处停止；仅在需要后端崩溃和 core dump 进行事后调试时使用 `panic`。

`pg_pathcheck.stage_checks` 启用额外的逐阶段检查：

```sql
SET pg_pathcheck.stage_checks = off;  -- default
SET pg_pathcheck.stage_checks = on;
```

启用后，模块会在 base-rel、join-rel 和 upper-rel hook 边界检查 pathlist，从而把发现的问题关联到更窄的规划阶段。常规运行建议保持关闭，因为额外遍历会增加规划开销，尤其是 join 较多的查询。

### 检查内容

模块会遍历 planner roots、relation pathlists、partial pathlists、cheapest path slots、parameterized paths、subquery subroots、subplan subroots，以及嵌套的 `Path` 字段，例如 join children、append children、sort subpaths、bitmap paths 和类似复合 path 成员。

它主要报告两类问题：

- Invalid `NodeTag`：指针不再像 PostgreSQL `Path` 节点。
- Parent mismatch：base 或 join relation 上看似有效的 `Path` 指向另一个 `RelOptInfo`，可能表示陈旧 path 指针存活后发生了同大小内存复用。

典型报告会包含异常 tag 或 mismatch、所在槽位（如 `pathlist`、`partial_pathlist`、`cheapest_total_path` 或嵌套 path 字段）、可解析出的 relation 名称、pathlist 细节，以及可通过 PostgreSQL debug 状态取得的查询字符串。

### 注意事项

`pg_pathcheck` 是面向 PostgreSQL 规划器开发和扩展测试的调试辅助工具，不是面向用户的 SQL 扩展。PostgreSQL cassert/debug 构建能提供更好的信号，因为已释放内存更容易识别。上游 README 提到 PG17/18 分支与 master 分支覆盖范围不同：PG17/18 会在 `create_plan` 和 `setrefs.c` 等后续规划阶段之前检查，而 master 可使用更新的 planner shutdown hook。
