## 用法

来源：

- [README.md](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/README.md)
- [REFERENCE.md](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/REFERENCE.md)
- [pg_hydradb.control](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/pg_hydradb.control)
- [sql/pg_hydradb--0.1.0.sql](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/sql/pg_hydradb--0.1.0.sql)
- [Cargo.toml](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/Cargo.toml)
- [License](https://github.com/hydra-db/hydradb/blob/v0.4.0/LICENSE)

`pg_hydradb` 0.1.0 随 HydraDB 0.4.0 发行，在普通 PostgreSQL 顶点表与边表上注册并遍历图。触发器在同一数据库内维护邻接数据。支持 PostgreSQL 16–18，不提供独立 HydraDB 服务的全部查询接口。

### 注册与遍历

由超级用户将不可迁移的扩展安装到固定的 `hydradb` 模式。在新数据库中，注册使用非空整数、文本或 UUID 键的表，键须有主键或唯一约束，边端点须有指向顶点的外键：

```sql
CREATE EXTENSION pg_hydradb;
CREATE TABLE people (id bigint PRIMARY KEY);
CREATE TABLE friendships (
    id bigint PRIMARY KEY,
    person_a bigint NOT NULL REFERENCES people(id),
    person_b bigint NOT NULL REFERENCES people(id)
);
INSERT INTO people VALUES (1), (2), (3);
INSERT INTO friendships VALUES (1, 1, 2), (2, 2, 3);
SELECT hydradb.create_graph('social',
    '[{"label":"person","table":"public.people","key":"id"}]',
    '[{"label":"knows","table":"public.friendships","key":"id", "source":"person_a","source_label":"person", "target":"person_b","target_label":"person"}]');
SELECT hydradb.reachable_vertices('social', '{"person":[1]}', '{"max_depth":3}');
SELECT hydradb.paths_between('social', '{"person":1}', '{"person":3}', '{}');
```

`hydradb.reachable_edges` 返回可达边，`hydradb.paths_from` 返回从种子出发的路径。选项包括 `max_depth`、`direction`、`edge_kind`、`edge_kinds` 和 `paths_per_pair`，结果使用注册的标签与键。查询函数检查源表访问权限；受 RLS 约束的调用者会被拒绝，而不是获得过滤后的图。应使用表所有者或本就有权绕过 RLS 的角色。

### 维护与恢复

在 READ COMMITTED 下运行 `hydradb.create_graph`、`hydradb.refresh_graph` 和 `hydradb.drop_graph`。不支持视图、分区表和继承。涉及图所用对象的破坏性 DDL 会被拒绝；列重写、触发器或外键缺失，以及超过 2,000 字节的文本键可能将图标记为需要刷新。修正原因后运行 `hydradb.refresh_graph`。备份须同时包含源数据及扩展的邻接与配置数据。

文档规定的升级路径是先删除每个图，执行不带 CASCADE 的 `DROP EXTENSION pg_hydradb`，安装匹配文件后重建扩展并重新注册图。REPEATABLE READ 写事务可能在刷新后收到 SQLSTATE 40001，需要重试。

### 可选预加载与遥测

查询无需预加载。将 `pg_hydradb` 加入现有共享预加载列表并重启，会启用共享活动计数器与每小时运行的使用报告 worker。报告默认开启，包含由集群标识派生的稳定安装 ID、版本、取整后的计数及使用和错误直方图；上游约定不发送查询文本、标识符、键或结果。`hydradb.usage_report()` 可查看当前数据库的报告。

设置 `pg_hydradb.usage_report = off` 并重载即可关闭报告，重新开启需要重启。`pg_hydradb.usage_report_url` 和 `pg_hydradb.usage_report_interval` 控制端点和间隔。不预加载时不会发送报告，也没有活动计数。源码采用 Elastic-2.0 许可证，部署时应核对其限制。
