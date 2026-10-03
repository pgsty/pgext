## 用法

来源：

- [pgext/README.md](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/README.md)
- [pgext/Cargo.toml](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/Cargo.toml)
- [pgext/tqvector.control](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/tqvector.control)
- [pgext/sql/tqvector--0.1.0.sql](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/sql/tqvector--0.1.0.sql)
- [pgext/src/lib.rs](https://github.com/ahb-sjsu/turboquant-pro/blob/c434a994f352936ebc91f40d9cd30ca813fe2eac/pgext/src/lib.rs)

`tqvector` 0.1.0 是 TurboQuant Pro 的 Rust 组件，用于保存标量量化后的嵌入向量。其构建特性声明支持 PostgreSQL 14–17，无需 Python 运行时。control 声明 superuser=false，但安装脚本会创建需要高权限的原生 C 类型与函数。未声明共享预加载要求。

### 基本用法

```sql
CREATE EXTENSION tqvector;
SELECT tq_dim(tq_compress(ARRAY[0.1,0.2,0.3,0.4]::float4[], 3));
SELECT tq_decompress(tq_compress(ARRAY[0.1,0.2,0.3,0.4]::float4[], 3));
```

### 操作与限制

`tq_compress` 接受 2、3 或 4 位量化；`tq_decompress` 重建近似浮点数组。可通过 `tq_dim`、`tq_bits`、`tq_norm`、`tq_ratio` 和 `tq_size_bytes` 检查值。`tq_cosine_sim` 与 `tq_cosine_dist` 比较维度相同的向量；`<=>` 表示余弦距离，`<->` 表示欧氏距离。量化会损失信息，应使用自己的数据验证检索质量。

安装 SQL 定义的是距离运算符，没有索引访问方法，因此按距离排序本身不会创建 ANN 索引。传给 `tq_bulk_compress` 的表名和列名应视为可信管理输入。可选 GPU 构建特性与基本 SQL 用法独立。
