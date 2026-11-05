## 用法

来源：

- [上游来源 1](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/README.md)
- [上游来源 2](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/pg_lion.control)

`pg_lion` 0.3.0 为 PostgreSQL 16 及以上版本提供 Roaring 位图倒排索引，用于过滤、计数与分组聚合。

### 核心流程

`lion` 访问方法支持标量相等与范围条件、数组、JSONB 和全文检索向量。索引选项 `summaries` 与 `slices` 通过增加存储和写入工作加速范围或位切片聚合。`lion_bm25_score` 可对全文匹配排序，查询仍使用 PostgreSQL 原有运算符。

```sql
CREATE EXTENSION pg_lion;
CREATE TABLE lion_events (country text, amount integer);
INSERT INTO lion_events VALUES ('JP', 10), ('JP', 20), ('NZ', 30);
CREATE INDEX ON lion_events USING lion (country);
VACUUM ANALYZE lion_events;
SELECT count(*) FROM lion_events WHERE country = 'JP';
```

### 运行边界

安装需要超级用户。预加载是可选项：将 `pg_lion` 加入 `shared_preload_libraries` 并重启，可启用自定义 WAL 格式。以该模式创建的索引在重建前，写入和 VACUUM 都依赖预加载。堆页面可见性及清理频率会影响仅索引计数，不能假定所有查询均不访问堆。
