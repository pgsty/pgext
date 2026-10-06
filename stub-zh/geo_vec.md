## 用法

来源：

- [README.md](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/README.md)
- [docs/API.md](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/docs/API.md)
- [Cargo.toml](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/Cargo.toml)
- [sql/geo_vec--0.1.0.sql](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/sql/geo_vec--0.1.0.sql)
- [geo_vec.control](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/geo_vec.control)

`geo_vec` 将向量近邻搜索与地理包围盒过滤组合到一个索引中。这个源码预览从 pgvectorscale 的 DiskANN 实现派生，使用独立的扩展名称。

### 核心用法

```sql
CREATE EXTENSION vector;
CREATE EXTENSION postgis;
CREATE EXTENSION geo_vec;
CREATE TABLE places (id bigint PRIMARY KEY, embedding vector(3), geom geometry(Point,4326));
CREATE INDEX ON places USING geo_vec (embedding vector_cosine_ops, geom);
SELECT id FROM places
WHERE geom && ST_MakeEnvelope(-122.5,37.7,-122.4,37.8,4326)
ORDER BY embedding <=> '[0.1,0.2,0.3]'::vector LIMIT 20;
```

### 运行边界

依赖 `vector` 与 `postgis`，安装需要超级用户。README 声明支持 PostgreSQL 17 及以后版本，所核对的清单提供了 17 与 18 特性；更旧的特性开关不能单独证明兼容性。上游未要求预加载或重启。

该访问方法支持余弦、L2 与内积向量算子类、几何包围盒和可选的小整数数组标签。`geo_vec.query_search_list_size` 与 `geo_vec.query_rescore` 在查询开销和召回率之间取舍；`geo_vec.spatial_brute_force_threshold` 让较小的空间候选集使用穷举扫描，较大的区域使用近似图搜索。应在目标工作负载上衡量召回率，并保持几何坐标系一致。

构建索引消耗内存和存储，原生构建要求相应的 AVX2/FMA 或 NEON 支持。目录记录源码可用性，不代表已验证的 Pigsty 软件包矩阵。
