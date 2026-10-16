## 用法

来源：

- [README.md](https://github.com/Dn-Media-Group-Public/age_gds/blob/4b9210f3511c53fd8f967f9995efacbda18d15a2/README.md)
- [docs/user-guide.md](https://github.com/Dn-Media-Group-Public/age_gds/blob/4b9210f3511c53fd8f967f9995efacbda18d15a2/docs/user-guide.md)
- [sql/age_gds--1.0.sql](https://github.com/Dn-Media-Group-Public/age_gds/blob/4b9210f3511c53fd8f967f9995efacbda18d15a2/sql/age_gds--1.0.sql)
- [age_gds.control](https://github.com/Dn-Media-Group-Public/age_gds/blob/4b9210f3511c53fd8f967f9995efacbda18d15a2/age_gds.control)

`age_gds` 为 Apache AGE 提供 NetworKit 图分析能力。安装的 SQL 扩展版本为 1.0，而独立的实现诊断函数 `age_gds.version()` 在文档中返回 0.1.0。可以从既有 AGE 图创建投影，运行算法，并按需写回结果。

### 核心用法

```sql
CREATE EXTENSION age;
LOAD 'age';
SET search_path = ag_catalog, "$user", public;
CREATE EXTENSION age_gds;
SELECT age_gds.graph_project(
  name => 'social_projection', age_graph => 'social',
  node_query => $$MATCH (n:Person) RETURN id(n) AS id$$,
  edge_query => $$MATCH (a:Person)-[:KNOWS]->(b:Person)
                 RETURN id(a) AS src, id(b) AS dst, 1.0 AS weight$$);
SELECT * FROM age_gds.pageRank_stream(graph_name => 'social_projection')
ORDER BY score DESC;
```

### 运行边界

先安装 Apache AGE，并在会话中载入 `age`。安装需要超级用户。源码要求 PostgreSQL 14 或更高版本、适配的 AGE、OpenMP 和随附的 NetworKit；实际兼容性仍取决于 AGE 与服务器的组合。文档没有要求全局预加载 `age_gds`。

`graph_project` 使用节点和边的 Cypher 查询物化具名投影。`pageRank_stream`、社区发现和最短路径函数返回结果行；对应写入函数会修改 AGE 图。`graph_list`、`graph_exists`、`graph_reindex` 和 `graph_drop` 用于管理投影。普通快照不会自动反映源数据变更，需要重建，或使用文档所述的实时索引模式。

投影文件默认写入服务器数据目录，与 AGE 表分开保存，备份恢复时应单独考虑。`age_gds.session_cache_mb` 限制每个后端的缓存，`age_gds.data_dir` 控制存储目录且需要超级用户权限。大型投影与算法会消耗内存和 CPU。
