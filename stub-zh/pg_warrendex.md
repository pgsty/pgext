## 用法

来源：

- [官方 README.md](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/README.md)
- [官方 pg_warrendex.control](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/pg_warrendex.control)
- [官方 Cargo.toml](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/Cargo.toml)
- [官方 lib.rs](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/src/lib.rs)
- [官方 pg_warrendex--0.0.0--0.1.0.sql](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/sql/pg_warrendex--0.0.0--0.1.0.sql)

`pg_warrendex` 0.1.0 提供覆盖索引访问方法 `warrendex`。有序段保存键和附加列，在可见性条件允许时使符合条件的查询避免访问堆表。上游仅声明在 PostgreSQL 18 上测试，Cargo 特性本身不能证明 PostgreSQL 19 受支持。

### 基本流程

```sql
CREATE EXTENSION pg_warrendex;
CREATE TABLE index_demo (id bigint, group_id integer, label text);
CREATE INDEX index_demo_cover ON index_demo
    USING warrendex (group_id, id) INCLUDE (label);
VACUUM (ANALYZE) index_demo;
EXPLAIN (ANALYZE, BUFFERS)
SELECT id, label FROM index_demo
WHERE group_id = 7 AND id >= 100
ORDER BY id;
```

安装要求超级用户，扩展不受信任且不可迁移模式。源码使用 pgrx 0.19.2 和同级 Rust 支持库 warren-pg-speller；后者是构建依赖，不是另一个 CREATE EXTENSION 依赖。

### 规划与维护

动态库加载时安装规划器钩子。若要在会话规划语句之前启用它们，可执行 `LOAD 'pg_warrendex'`；可选的 `shared_preload_libraries` 配置会跨会话生效，并要求重启。普通索引操作与规划器钩子优化是两项不同能力。

键采用升序且空值置后，拒绝 `DESC` 和 `NULLS FIRST`。操作符类覆盖常用整数、文本、数值、时间、布尔、UUID 和二进制类型，不接受存储参数。应检查执行计划与堆读取计数，不能假定所有查询都会使用仅索引扫描。

插入操作可能在写事务中封存并合并有序段。VACUUM 标记失效条目，`REINDEX` 合并段并回收废弃页。扫描、写入和 VACUUM 会拒绝其他扩展版本生成的索引，因此版本变化后须先重建再使用。应使用具有代表性的可丢弃数据评估此早期版本。
