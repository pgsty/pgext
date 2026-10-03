## 用法

来源：

- [Official README.md](https://github.com/Mateus504R35/pgfaiss/blob/8cf777c2f38379958d0be447037ae37affbe78a3/README.md)
- [Official faiss_pg.control](https://github.com/Mateus504R35/pgfaiss/blob/8cf777c2f38379958d0be447037ae37affbe78a3/faiss_pg.control)
- [Official faiss_pg--0.2.sql](https://github.com/Mateus504R35/pgfaiss/blob/8cf777c2f38379958d0be447037ae37affbe78a3/faiss_pg--0.2.sql)
- [Official faiss_pg.cpp](https://github.com/Mateus504R35/pgfaiss/blob/8cf777c2f38379958d0be447037ae37affbe78a3/faiss_pg.cpp)

`faiss_pg` 0.2 由 pgfaiss 项目提供，从 PostgreSQL 数组创建 Faiss 文件并通过 SQL 搜索。它是研究型扩展，没有实现 PostgreSQL 索引访问方法或规划器集成的索引扫描。

### 创建与搜索

```sql
CREATE EXTENSION faiss_pg;
CREATE TABLE items (id bigint PRIMARY KEY, embedding real[] NOT NULL);
INSERT INTO items VALUES (1, ARRAY[0.1,0.2,0.3]::real[]),
                         (2, ARRAY[0.8,0.7,0.6]::real[]);
SELECT faiss_build_index('items_flat_l2', 'public.items'::regclass,
  'id'::name, 'embedding'::name, 'l2', 'flat', false,
  '{"fetchBatchSize":10000}');
SELECT * FROM faiss_search('items_flat_l2', ARRAY[0.1,0.2,0.3]::real[], 2);
```

### 输入与结果

来源键必须为 integer 或 bigint；嵌入向量为长度一致的一维 `real[]`，键、数组及数组元素都不能为空。Flat 执行精确搜索，HNSW 与 IVF-Flat 执行近似搜索。L2 结果为平方距离，越小越好；内积与归一化余弦返回相似度，越大越好。`faiss_search_batch()` 接受展平后的查询向量，返回从零开始的查询编号、ID 与距离。搜索参数包括 `efSearch` 与 `nprobe`。

### 文件、更新与访问

`faiss_indexes` 保存文件路径与来源元数据。默认服务端目录是 /var/lib/postgresql/faiss_indexes，PostgreSQL 操作系统用户需要写权限。来源表写入后索引被标记为失效，重建前搜索会报错。构建时会对来源表取得 SHARE 锁。文件按后端进程缓存；`faiss_clear_cache()` 与 `faiss_clear_all_cache()` 可释放缓存。外部文件需要独立备份与生命周期管理。由管理员安装，并将函数、元数据表和文件系统权限限制给可信操作人员；此原型没有单独的文件访问权限边界。

### 运行要求

源码使用 Faiss CPU 库 1.13.1、C++17、OpenMP 与 BLAS/LAPACK。失效标记辅助函数由 PL/pgSQL 实现。README 提供 Linux 上 PostgreSQL 16 的示例，但未定义受支持的大版本范围，也未要求共享预加载。

SQL 保留了默认的 `PUBLIC` 函数执行权限，C 文件操作没有显式的超级用户或文件角色 ACL 检查。应将调用方选择的 `indexDir` 与元数据中的文件路径视为可信输入；普通表权限不能构成完整的文件系统访问边界。
