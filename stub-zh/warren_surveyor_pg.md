## 用法

来源：

- [warren-surveyor-pg/README.md](https://github.com/soulstompp/warren-surveyor/blob/59f4d0f475cd6a8ed6811be945b9738c35932eca/warren-surveyor-pg/README.md)
- [warren-surveyor-pg/Cargo.toml](https://github.com/soulstompp/warren-surveyor/blob/59f4d0f475cd6a8ed6811be945b9738c35932eca/warren-surveyor-pg/Cargo.toml)
- [warren-surveyor-pg/src/lib.rs](https://github.com/soulstompp/warren-surveyor/blob/59f4d0f475cd6a8ed6811be945b9738c35932eca/warren-surveyor-pg/src/lib.rs)
- [warren-surveyor-pg/warren_surveyor_pg.control](https://github.com/soulstompp/warren-surveyor/blob/59f4d0f475cd6a8ed6811be945b9738c35932eca/warren-surveyor-pg/warren_surveyor_pg.control)

`warren_surveyor_pg` 0.1.0 提供 `surveyor` 索引访问方法。其标记索引不存储用户数据，也不会出现在查询扫描中；规划钩子读取现有索引以改进估算，同时启用随附的查询等价改写钩子。

### 核心用法

```sql
CREATE EXTENSION warren_surveyor_pg;
LOAD 'warren_surveyor_pg';
CREATE TABLE survey_demo (id bigint PRIMARY KEY, value text);
CREATE INDEX ON survey_demo USING surveyor (id);
ANALYZE survey_demo;
EXPLAIN (ANALYZE, BUFFERS) SELECT * FROM survey_demo WHERE id = 42;
```

### 运行边界

在表的唯一键上创建标记索引，保留实际的 B-tree/GIN/GiST 索引，并分析表。参与的会话须在规划前载入库；示例用 `LOAD` 启用当前会话。会话预加载可用于后续连接，无须全局重启。安装要求超级用户。

README 面向 PostgreSQL 18；Cargo 另有 PostgreSQL 19 构建特性，但这不代表经过验证的兼容承诺。测量增加规划阶段的 I/O 和 CPU 开销，不保证改善所有工作负载。应同时检查估算与实际行数，以及规划和执行阶段的缓冲区读写。`warren_surveyor_pg.planning_read_limit` 限制读取量，达到边界后使用原估算；递归查询可能在规划期间按所述预算执行。

扩展只改写其认为等价的查询结构。全面启用前应对代表性负载比较结果集与执行计划。这是初始源码版本，与 `pg_warrendex` 是不同扩展。
