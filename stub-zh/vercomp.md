## 用法

来源：

- [已复核 commit 的 vercomp README](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/README.md)
- [已复核 commit 的 vercomp 1.0.0 安装 SQL](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/vercomp--1.0.0.sql)
- [已复核 commit 的 vercomp C 源码](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/vercomp.c)
- [上游推荐的后继 pg_semver](https://github.com/eendroroy/pg_semver)

1.0.0 版 `vercomp` 定义 `version` 数据类型、类型转换、比较与范围运算符，以及默认 B-tree 和哈希操作符类。除普通比较外，它还支持 `~`、`!~`、`^` 和 `!^` 版本范围运算符，并提供 `version_cmp`、`version_bump` 等辅助函数。

### 存储与比较版本

```sql
CREATE EXTENSION vercomp;

CREATE TABLE component_versions (
  component text PRIMARY KEY,
  release version NOT NULL
);

INSERT INTO component_versions VALUES
  ('alpha', '1.2.3'),
  ('beta', '2.0.0-rc1'),
  ('gamma', '2.0.0');

SELECT component, release
FROM component_versions
ORDER BY release;

SELECT version_cmp('1.2.3', '2.0.0');
```

### 注意事项

- 上游将该项目标为过时、归档了仓库，并推荐 `pg_semver` 作为后继。新部署应优先选择仍在维护的替代方案。
- README 含有一条旧说明，要求把 `$libdir/vercomp` 加入 `shared_preload_libraries` 并重启。但已复核的 control/C 源码没有声明 `_PG_init`、钩子、工作进程或共享内存初始化；catalog 证据也记录为无需加载。应将 README 步骤视为未解决的文档冲突，而非已经证明的运行时必需项。
- 上游环境测试面向 PostgreSQL 10，不能证明与当前版本兼容。
- 既有索引和列依赖该扩展的 `version` 类型与操作符类；替换前应规划显式数据转换。
