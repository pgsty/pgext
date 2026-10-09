## 用法

来源：

- [README.md](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/README.md)
- [pgv_extra/sql/vector_extra--0.1.0.sql](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/pgv_extra/sql/vector_extra--0.1.0.sql)
- [pgv_extra/meson.build](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/pgv_extra/meson.build)
- [pgv_extra/vector_extra.control](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/pgv_extra/vector_extra.control)

`vector_extra` 是 Tensor Dancer 源码分发中的矩阵向量辅助扩展。版本化 SQL 定义了 `pgv_mulmv(matrix bytea, vec vector)`，返回 `vector`。

### 核心用法

```sql
CREATE EXTENSION vector;
CREATE EXTENSION vector_extra;
SELECT oid::regprocedure FROM pg_proc WHERE proname = 'pgv_mulmv';
```

### 运行边界

必须先安装 `vector`：SQL 使用了它的数据类型，但控制文件没有声明依赖。矩阵采用上游专用二进制表示，不能传入任意字节数组，应由匹配版本的项目工具生成。构建链接 BLAS、LAPACK 和 GGML。这是早期源码原型，没有核实的 PostgreSQL 主版本支持范围或稳定的矩阵交换契约。控制文件未标记为受信任，也未要求预加载。
