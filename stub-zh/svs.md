## 用法

来源：

- [官方文档](https://github.com/intel/SVS-Extension-for-PostgreSQL/blob/d504d1c49536fd8f9ff08cff0c0432d84978c8b2/README.md)
- [扩展控制文件](https://github.com/intel/SVS-Extension-for-PostgreSQL/blob/d504d1c49536fd8f9ff08cff0c0432d84978c8b2/svs.control)
- [官方仓库](https://github.com/intel/SVS-Extension-for-PostgreSQL)

`svs` 为 pgvector 提供带 LeanVec 与 LVQ 压缩的 SVS Vamana 近似近邻索引。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `svs`：

```sql
CREATE EXTENSION svs CASCADE;
```

经审查的控制文件或官方流程要求 `vector`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE EXTENSION vector;
CREATE EXTENSION svs;

CREATE TABLE items (id serial PRIMARY KEY, embedding vector(768));

-- Create a Vamana index
CREATE INDEX ON items USING vamana (embedding vector_l2_ops);

-- Nearest-neighbor search
SELECT id FROM items ORDER BY embedding <-> '[0.1, 0.2, ...]' LIMIT 10;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `vamana` | ACCESS METHOD | 用于表或索引定义的访问方法。 |
| `vector_l2_ops` | OPERATOR CLASS | 供索引使用的操作符类。 |
| `halfvec_cosine_ops` | OPERATOR CLASS | 供索引使用的操作符类。 |
| `halfvec_ip_ops` | OPERATOR CLASS | 供索引使用的操作符类。 |
| `halfvec_l2_ops` | OPERATOR CLASS | 供索引使用的操作符类。 |
| `pg_stat_vamana_worker` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_stat_vamana_worker_slot` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `svs_restart_worker` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
