## 用法

来源：

- [已归档的上游 README 与停用公告](https://github.com/timescale/pg_prometheus/blob/b5a113ef014597e9adcb1021d40bba005a8d741a/README.md)
- [固定提交的扩展 SQL](https://github.com/timescale/pg_prometheus/blob/b5a113ef014597e9adcb1021d40bba005a8d741a/sql/prometheus.sql)
- [固定提交的样本解析器](https://github.com/timescale/pg_prometheus/blob/b5a113ef014597e9adcb1021d40bba005a8d741a/src/parse.c)
- [固定提交的扩展控制文件](https://github.com/timescale/pg_prometheus/blob/b5a113ef014597e9adcb1021d40bba005a8d741a/pg_prometheus.control)

pg_prometheus 0.2.2 版为 Prometheus 展示格式样本定义 prom_sample 类型，并创建原始或规范化的 PostgreSQL 存储及查询视图。它为旧版 Prometheus PostgreSQL 远程存储适配器设计，也可选择使用 TimescaleDB 超表。

### 创建并查询存储

上游要求把该库加入 shared_preload_libraries，并在安装前重启服务器：

```conf
shared_preload_libraries = 'pg_prometheus'
```

```sql
CREATE EXTENSION pg_prometheus;

SELECT create_prometheus_table(
    'metrics',
    normalized_tables => true,
    use_timescaledb => false
);

INSERT INTO metrics (sample)
VALUES ('cpu_usage{service="nginx",host="machine1"} 34.6 1494595898000');

SELECT time, value, labels
FROM metrics
WHERE name = 'cpu_usage'
ORDER BY time DESC;
```

规范化存储把标签集与时间戳/值行分开；原始存储保存完整的 prom_sample。辅助函数会在调用者模式中创建多个表、索引、视图和触发器。COPY 集成应写入生成的复制表，而不是视图。

### 已归档的集成

Timescale 先把项目置于停用维护模式，之后归档了仓库。审阅源码最后更新于 2020 年，Docker 示例面向 PostgreSQL 11，适配器与 TimescaleDB 集成也反映了当时的接口。README 指向后继项目，但这并不会让当前扩展重新变成受维护软件。

应把 pg_prometheus 视为遗留数据格式和迁移来源。需要用不可信的展示格式输入验证 C 解析器，检查所有生成对象名称与权限，明确保留策略和分区管理，并测量标签基数增长。没有完整的安全、兼容与迁移审查，不要部署旧适配器或已归档容器。
