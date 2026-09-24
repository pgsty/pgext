## 用法

来源：

- [Control file](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/pg_weave.control)
- [Makefile](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/Makefile)
- [README](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/README.md)
- [Installation SQL](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/sql/pg_weave--0.1.0.sql)
- [Vector storage boundary](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/sql/pg_weave--0.7.0--0.8.0.sql)
- [Latest upgrade SQL](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/sql/pg_weave--0.9.0--0.10.0.sql)
- [Regression examples](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/sql/weave.sql)
- [Module initialization and settings](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/src/am/customscan.c)
- [Licensing and provenance](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/doc/LICENSING.md)
- [Migration status](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/doc/MIGRATION.md)
- [PostgreSQL extension trust](https://www.postgresql.org/docs/18/extend-extensions.html)

`pg_weave` 通过 `weave` 索引访问方法提供 BM25 词法检索。本次核对的控制版本为 0.10.0。扩展仍处于实验阶段：词法通道可查询，向量通道仅支持存储，宣传中的模糊检索与融合搜索工作流尚未完成。

### 启用与核心工作流

上游面向 PostgreSQL 17 及以上版本，并报告在 17 和 18 上进行了测试。安装匹配的扩展文件后，启用扩展并为分析后的文档建立索引：

```sql
CREATE EXTENSION pg_weave;

CREATE TABLE docs (id serial PRIMARY KEY, d wdoc);
INSERT INTO docs (d) VALUES
  (to_wdoc('english', 'postgres streaming replication')),
  (to_wdoc('english', 'a slow green turtle'));

CREATE INDEX docs_weave ON docs USING weave (d);

SELECT id FROM docs WHERE d @@@ to_wquery('english', 'postgres') ORDER BY id;
SELECT weave_count('docs_weave', to_wquery('english', 'postgres'));
SELECT ndocs, avgdl, nterms FROM weave_index_stats('docs_weave');
```

默认操作符类 `wdoc_lex_ops` 为 `wdoc` 值建立索引。`to_wdoc` 与 `to_wquery` 接受文本搜索配置；文档与查询应使用同一配置。`@@@` 操作符匹配查询，`<=>` 支持按相关性排序。

无需显式预加载或重启。控制文件声明 `trusted = true` 和 `relocatable = true`；服务器已安装扩展文件时，具有数据库 `CREATE` 权限的用户即可安装。这是上游的可信扩展声明，并不表示项目已经适合生产使用。

### 检查与维护

- `weave_count` 通过索引统计匹配数。
- `weave_index_stats` 返回文档数、平均文档长度与词汇表大小。
- `weave_check` 检查索引结构。
- `weave_merge` 将待处理数据合并到分段；`weave_vacuum` 执行索引压缩整理。

`pg_weave.wand_initial_k` 默认为 32，控制排序扫描的初始候选宽度。`pg_weave.build_collapse_max_mb` 默认为 4096 MiB；更大的索引构建会保留有界层级，而不会强制合并为单个分段。两项配置均允许用户设置。用于具体负载前，应检查有代表性的查询计划与维护开销。

### 版本与功能边界

安装过程从 0.1.0 基线开始，经过九个升级脚本到达 0.10.0。仓库内元数据仍写作 0.7.0，README 状态仍写作 0.3.0；本次安装版本依据控制文件与完整 SQL 升级链确定。

`wvec` 类型已经存在，但 `wvec_weave_ops` 仅登记存储能力，没有查询操作符。README 中的 `score()`、`fuse()` 函数及其拟议的多通道操作符类并不存在于已安装 SQL 中。请使用上文已核对的词法工作流。从其他搜索扩展自动迁移的路径也被上游标为尚未实现。

项目声明使用 PostgreSQL 许可证，并记录了随附组件各自的 MIT 与 BSD 许可证。规范上游仓库位于 Codeberg，GitHub 为同步镜像。源码可用和声明为可信扩展，并不代表已有发行包或生产支持。
