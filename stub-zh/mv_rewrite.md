## 用法

来源：

- [固定提交的上游 README](https://github.com/d-e-n-t-y/pg_fdw_mv_rewrite/blob/0b5ed382c3c956309a7378f74c8619478bb3b3b0/README.md)
- [0.6.2 版安装 SQL](https://github.com/d-e-n-t-y/pg_fdw_mv_rewrite/blob/0b5ed382c3c956309a7378f74c8619478bb3b3b0/mv_rewrite--0.6.2.sql)
- [固定提交的 planner hook 实现](https://github.com/d-e-n-t-y/pg_fdw_mv_rewrite/blob/0b5ed382c3c956309a7378f74c8619478bb3b3b0/mv_rewrite.c)
- [固定提交的扩展入口](https://github.com/d-e-n-t-y/pg_fdw_mv_rewrite/blob/0b5ed382c3c956309a7378f74c8619478bb3b3b0/extension.c)

mv_rewrite 0.6.2 版是实验性 planner 扩展，会把已登记 materialized view 作为候选执行路径。如果找到语义可用且估算成本更低的视图，EXPLAIN 会显示 Custom Scan (MVRewriteScan)，查询会透明读取该 materialized view。

### 登记 materialized view

```sql
CREATE EXTENSION mv_rewrite;

CREATE TABLE rewrite_sales (region text, amount numeric);
INSERT INTO rewrite_sales VALUES
    ('east', 10),
    ('east', 20),
    ('west', 30);

CREATE MATERIALIZED VIEW rewrite_sales_mv AS
SELECT region, count(amount) AS n
FROM rewrite_sales
GROUP BY region;

SELECT mv_rewrite.enable_rewrite('rewrite_sales_mv');

EXPLAIN (VERBOSE)
SELECT region, count(amount)
FROM rewrite_sales
GROUP BY region;
```

安装脚本会把当前数据库的 session_preload_libraries 修改为带版本号的模块名，并在安装会话内加载它。新会话会继承该数据库设置。安装前必须审查这个持久的数据库级副作用。

### 正确性与支持边界

上游明确说明代码尚未达到生产可用。它只考虑部分 GROUP BY aggregate、DISTINCT、ordering、简单 SELECT、join、WHERE 和 HAVING 形状；recursive query 与 lateral join 没有得到正确支持。planner matching 成本较高，因此扩展提供 minimum-cost threshold、enabled-table list、progress logging 和全局 enable switch。

登记操作不会刷新 materialized view。即使匹配逻辑正确，过期视图仍会让重写查询返回过期结果。必须建立 refresh 和 dependency 流程，并在关闭重写时对比 plan 与结果。源码依赖旧版 PostgreSQL planner 内部接口，自 2019 年后没有更新。应隔离测试每类查询、扩展交互和 PostgreSQL 主版本，并保留可靠移除数据库 preload 设置的方法。
