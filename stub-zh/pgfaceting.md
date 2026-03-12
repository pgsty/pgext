

## 用法

> [pgfaceting: 使用 Roaring 位图倒排索引实现快速分面搜索](https://github.com/cybertec-postgresql/pgfaceting)

`pgfaceting` 扩展通过使用 Roaring 位图构建的倒排索引实现快速分面计数。依赖 `pg_roaringbitmap` 扩展。

```sql
CREATE EXTENSION pgfaceting;
```

### 分面类型

- **`plain_facet(column)`**：直接使用列值作为分面
- **`datetrunc_facet(column, precision)`**：按日期截断（例如按月/按年分组）
- **`bucket_facet(column, buckets)`**：将连续变量划分到预定义的区间

### 核心函数

```sql
-- 为表创建分面基础设施
SELECT pgfaceting.add_faceting_to_table(
    'products',
    'id',
    ARRAY[
        plain_facet('color'),
        plain_facet('size'),
        bucket_facet('price', ARRAY[0, 10, 50, 100, 500])
    ]
);

-- 运行维护以合并增量变更
SELECT pgfaceting.run_maintenance();

-- 合并指定表的增量数据
SELECT pgfaceting.merge_deltas('products');

-- 获取前 N 个分面值
SELECT pgfaceting.top_values('products', 10);

-- 按分面过滤条件统计结果数量
SELECT pgfaceting.count_results('products', filters);
```

### 架构

该扩展为每个被索引的表维护两张辅助表：一张存储 Roaring 位图的主分面表，将分面值映射到行 ID；另一张增量表用于在维护运行之间记录增量变更。

目前仅支持 32 位整数 ID 列。
