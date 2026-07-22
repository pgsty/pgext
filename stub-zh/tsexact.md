## 用法

来源：

- [Official extension control file](https://github.com/postgrespro/tsexact/blob/b08a6ce7ed40b5b62cccdb12d7e138c653d2efd0/tsexact.control)
- [Official upstream documentation](https://github.com/postgrespro/tsexact/blob/b08a6ce7ed40b5b62cccdb12d7e138c653d2efd0/README.md)

`tsexact` 1.0 为 PostgreSQL 全文检索向量增加精确片段辅助函数。它原本用于在 PostgreSQL 9.5 及更早版本模拟短语搜索；上游建议 PostgreSQL 9.6 及以后版本使用内置短语搜索操作符。

### 核心流程

```sql
CREATE EXTENSION tsexact;

SELECT ts_exact_match(
  'cat:3 fat:2 sad:5'::tsvector,
  'cat:2 fat:1 sad:4'::tsvector
);

SELECT ts_squeeze('cat:2,9 fat:1,6 sad:4'::tsvector);
```

`ts_exact_match(tsvector, tsvector)` 检查片段的相对位置是否以某个偏移出现在文档中，并忽略权重。它的三参数重载接受针对文档的权重掩码。`ts_squeeze(tsvector)` 消除位置编号间隙，`setweight(tsquery, text)` 给查询词元设置权重，`poslen(tsvector)` 返回总位置跨度。

### 索引模式与注意事项

全文索引不能直接支持 `ts_exact_match`。先用可索引的 `@@` 条件筛选候选，再以 `ts_exact_match` 做精确复核。

```sql
SELECT *
FROM documents
WHERE search_vector @@ plainto_tsquery('fat rat')
  AND ts_exact_match(
        search_vector,
        ts_squeeze(to_tsvector('fat rat'))
      );
```

分词、词典、停用词、权重与位置间隙都会影响结果。在当前 PostgreSQL 上应优先使用原生短语查询；只有确实需要旧向量语义且有回归测试覆盖时才保留 `tsexact`。
