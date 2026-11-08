## 用法

来源：

- [course_embedding_search.control](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docker/db/extension/course_embedding_search.control)
- [course_embedding_search--1.0.sql](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docker/db/extension/course_embedding_search--1.0.sql)
- [course_embedding_search.cpp](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docker/db/extension/course_embedding_search.cpp)
- [search-queries.md](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docs/backend/search-and-discovery/search-queries.md)
- [20261028_0022_course_embedding_search_blocks.py](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/migrations/versions/20261028_0022_course_embedding_search_blocks.py)
- [Dockerfile](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docker/db/Dockerfile)

`course_embedding_search` 是 course-explorer 应用内的原生扫描扩展。SQL 1.0 版本面向项目的 PostgreSQL 18 和 pgvector 环境，扫描应用打包的向量块，返回各组最大的余弦分数。

### 启用与检查

安装匹配的库和 SQL 文件后，由超级用户启用两个扩展：

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS course_embedding_search;
SELECT oid::regprocedure
FROM pg_proc WHERE proname = 'course_embedding_search_scan';
```

应用迁移先安装扩展，再创建搜索表。`course_embedding_search_scan(regclass, vector, float8, int)` 接受向量块关系、已经减去中心向量的查询向量、查询向量的平方范数和分组数量，返回 `group_no` 和 `score`；应用 SQL 随后关联分组身份并按分数阈值过滤。

### 数据与运维边界

此函数不能通用地扫描任意向量列。它要求应用规定的准确列序和打包格式：块编号、行数、整数分组数组、平方范数数组及 float32 字节载荷。服务制品生成器用与查询一致的嵌入配置和中心向量准备这些数据块。维度、范数、分组 ID 或数组及载荷大小不合法时会报错。

函数具有严格、易变和并行不安全属性。它按给定分组数量分配最大分数数组，并使用当前快照扫描。原生关系扫描不执行普通 SQL 的表权限或行级安全检查。应将函数执行权限限定给可信应用角色，只接受经过应用验证的关系和输入。

扩展可迁移，首次安装无需预加载或重启。它不创建应用表、不获取嵌入向量，也不生成服务制品。恢复时需保留匹配的扩展、迁移后的表结构和制品生成配置。
