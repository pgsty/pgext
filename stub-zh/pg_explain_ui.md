

## 用法

该函数会对查询执行 EXPLAIN，并将结果发送到 [`explain.dalibo.com`](https://explain.dalibo.com/) 进行可视化展示。

> 注意：该操作会向互联网服务发送 HTTP POST 请求！

```bash
SELECT explain_ui($$query$$);
```

该扩展的设计理念可以参考这篇博客文章：[Writing a Postgres Extension With Pgrx for Visual Query Plans](https://davidgomes.com/writing-postgres-extension-with-pgrx-query-plans/)


### 示例

首先生成 pgbench 示例表：

```bash
pgbench -is10
```

然后使用以下语句对查询进行可视化解释：

```sql
CREATE EXTENSION explain_ui;
SELECT explain_ui($$SELECT * FROM pgbench_accounts;$$);
```

执行后会返回一个可视化执行计划的 URL：

```
postgres@u22:5432/postgres=# SELECT explain_ui($$SELECT * FROM pgbench_accounts;$$);
                    explain_ui
--------------------------------------------------
 https://explain.dalibo.com/plan/05377227a29f0418
(1 row)

Time: 2284.667 ms (00:02.285)
```

点击或在浏览器中打开该 [URL](https://explain.dalibo.com/plan/05377227a29f0418) 即可查看可视化执行计划。