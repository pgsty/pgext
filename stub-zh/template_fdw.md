## 用法

来源：

- [固定提交的上游 README](https://github.com/okbob/template_fdw/blob/621f30e37f6ad530c1ed47d8764383368a7a3831/README.md)
- [固定提交的扩展 control 文件](https://github.com/okbob/template_fdw/blob/621f30e37f6ad530c1ed47d8764383368a7a3831/template_fdw.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/template_fdw/)

`template_fdw` 创建仅含 schema、并有意拒绝读写的外部表。它们可作为临时表的持久模板，让 `plpgsql_check` 等工具在静态分析期间解析表结构，同时防止误用持久模板关系。

```sql
CREATE EXTENSION template_fdw;
CREATE SERVER template_server FOREIGN DATA WRAPPER template_fdw;

CREATE FOREIGN TABLE public.template_rows (a integer, b integer)
  SERVER template_server;

-- Both statements against the foreign table are expected to fail.
SELECT * FROM public.template_rows;
INSERT INTO public.template_rows VALUES (10, 20);

-- A temporary table can safely copy and shadow the template definition.
CREATE TEMP TABLE template_rows
  (LIKE public.template_rows INCLUDING ALL);
INSERT INTO template_rows VALUES (10, 20);
SELECT * FROM template_rows;
```

该 FDW 是 development 与 validation guard，不是 remote storage。外部表上的任何成功 `SELECT` 或 DML 都违背设计意图，应让 regression test 失败。control 版本 `0.0` 与 upstream/PGXN release `0.0.2` 不同；应通过 source commit 识别 build，并在每次 PostgreSQL upgrade 后测试 error behavior。

应把模板放在专用 schema，只授予 validation role 所需权限，并显式限定 persistent name，避免 search-path 变化意外指向真实表。
