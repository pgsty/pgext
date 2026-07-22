## 用法

来源：

- [扩展控制文件](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/pgrollup.control)
- [版本 1.0 安装 SQL](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/pgrollup--1.0.sql)
- [上游项目状态](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/README.md)

`pgrollup` 版本 `1.0` 是一个实验性框架，它解析伪 DDL，并创建表、触发器、函数和视图，以增量维护聚合结果。

### 核心流程

创建扩展前，需要在 PostgreSQL 的 Python 环境中安装不受信任语言 `plpython3u` 和上游 Python 包。先用 `dry_run` 预览生成的 DDL。

```sql
CREATE EXTENSION pgrollup;

SELECT pgrollup_parse($rollup$
CREATE INCREMENTAL MATERIALIZED VIEW sales_rollup AS (
  SELECT region, count(*), sum(amount)
  FROM sales
  GROUP BY region
);
$rollup$, dry_run => true);
```

审核输出 SQL 后，再以 `dry_run => false` 运行。默认维护模式是 `trigger`；`manual` 使用 `do_rollup`，而 `cron` 还需要 `pg_cron`。可用 `rollup_mode` 切换模式，用 `assert_rollup` 将汇总与重新计算结果比较，并用 `drop_rollup` 进行受管清理。

### 支持的聚合与注意事项

内置代数定义覆盖 count、sum、min、max、布尔聚合、平均值、方差与标准差。只有相应扩展已经存在时，才会启用可选的 HLL、t-digest、TopN 与 DataSketches 定义。

该项目会创建动态 SQL、事件触发器、数据触发器与不受信任的 Python 函数。手动与定时维护依赖合适的单列序列或汇总键。上游 README 大部分仍是未完成工作清单，并记录了解析器与查询限制，包括不支持外连接和子查询。应把生成对象视为实验性功能：审核每条命令，测试插入、更新、删除与模式变更，定期运行 `assert_rollup`，并保留独立计算的真实基准。
