## 用法

来源：

- [0.6.39 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.39)
- [pgrdf.control](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/pgrdf.control)
- [README.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/README.md)
- [Cargo.toml](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/Cargo.toml)
- [guide/01-install.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/01-install.md)
- [guide/tour.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/tour.md)
- [guide/05-graphs.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/05-graphs.md)
- [guide/06-validation-recipes.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/06-validation-recipes.md)

`pgrdf` 0.6.39 在 PostgreSQL 内保存 RDF 图，支持 Turtle/TriG/N-Quads 导入、SPARQL 查询与更新、SHACL 校验和 RDFS/OWL 推理。先预加载 `pgrdf` 并重启，再由超级用户创建扩展。

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
