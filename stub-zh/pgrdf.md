## 用法

来源：

- [Official source](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.41)
- [Official source](https://github.com/styk-tv/pgRDF/blob/v0.6.41/pgrdf.control)

- [0.6.40 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.40)
- [pgrdf.control](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/pgrdf.control)
- [README.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/README.md)
- [Cargo.toml](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/Cargo.toml)
- [guide/01-install.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/guide/01-install.md)
- [guide/tour.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/guide/tour.md)
- [guide/05-graphs.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/guide/05-graphs.md)
- [guide/06-validation-recipes.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/guide/06-validation-recipes.md)
- [0.6.39 to 0.6.40 upgrade](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/sql/pgrdf--0.6.39--0.6.40.sql)

`pgrdf` 0.6.41 在 PostgreSQL 内保存 RDF 图，支持 Turtle/TriG/N-Quads 导入、SPARQL 查询与更新、SHACL 校验和 RDFS/OWL 推理。先预加载 `pgrdf` 并重启，再由超级用户创建扩展。

### 核心用法

```ini
shared_preload_libraries = 'pgrdf'
```

```sql
CREATE EXTENSION pgrdf;
SELECT pgrdf.add_graph('http://example.org/people');
SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> . ex:alice ex:name "Alice" .',
  pgrdf.graph_id('http://example.org/people'));
SELECT * FROM pgrdf.sparql('SELECT ?s ?p ?o WHERE { ?s ?p ?o } LIMIT 10');
SELECT * FROM pgrdf.surface();
```

### 运行边界

固定的 `pgrdf` 模式提供 `add_graph`、`graph_id`、`parse_turtle`、`load_turtle`、`sparql`、`materialize`、`validate`、`stats` 和 `surface`。文件加载器读取服务端路径，应谨慎授权。`can_clear_graphs()` 报告清空图所需权限；非所有者写图仍需文档规定的底层表权限，并遵循图锁。`shacl_capability()` 显示支持的校验范围，路径深度／截断设置限制遍历。源码构建选项覆盖 PostgreSQL 14–18，已发布的 Linux 二进制面向 PG18。0.6.39 使用不复用的序列分配图 ID，图身份由 IRI 确定，编号间隙属于正常情况。安装匹配的库后，在各数据库执行 `ALTER EXTENSION pgrdf UPDATE`；新库遇到旧 SQL 会以 SQLSTATE 55000 拒绝操作。不要使用无法正常构建的 0.6.35–0.6.38 PGXN 源码归档。备份须共同保留并验证图与字典数据。

### 0.6.40 的所有者修复

如果存储表曾转交给其他角色，0.6.39 的升级可能让 `_pgrdf_graph_id_seq` 仍归升级执行者所有，从而无法创建新图。0.6.40 将该序列所有者对齐到 `_pgrdf_quads` 的所有者；只授予序列使用权限并不足够。`pgrdf.ownership_drift()` 报告存储对象所有者的差异及建议修复命令，更改角色前应逐项检查。安装匹配二进制后，在各数据库执行 SQL 升级；0.6.34–0.6.38 可直接升级到 0.6.40。

### 版本 0.6.41：图集合与备份

新图分区对主语、谓词、宾语三元组实施唯一性约束，重复导入不会新增重复项，显式断言优先于推导结果。原地升级不会重写已有重复三元组的图，只会发出警告。修改现有数据前，应检查 `pgrdf.duplicate_report()` 并审查 `pgrdf.dedup_graph()` 操作。

本版本修复逻辑转储遗漏图注册信息、默认图三元组及字典分配器的问题。升级后应重新备份并验证恢复。已存储的 U+0000 字面量现在读取时会报错，可用 `pgrdf.nul_term_report()` 定位后修复。图所有者、所有者约束及图操作记录增加了管理控制；应检查权限，不要假定此前向 PUBLIC 开放的维护函数仍可调用。

安装匹配二进制，重启以加载预加载动态库，再执行 `ALTER EXTENSION pgrdf UPDATE;`。PostgreSQL 18 至少需要 18.7。旧 SQL 与动态库不匹配保护及既有序列所有权检查仍然适用。
