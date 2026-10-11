## 用法

来源：

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_search/alohadb_search--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/alohadb_search--1.0.sql)
- [contrib/alohadb_search/search_bm25.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/search_bm25.c)
- [contrib/alohadb_search/search_suggest.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/search_suggest.c)
- [contrib/alohadb_search/meson.build](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/meson.build)
- [contrib/alohadb_search/alohadb_search.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/alohadb_search.control)

`alohadb_search` 1.0 是 AlohaDB 的文本相关性、模糊与语音匹配、搜索建议及地理距离模块。审阅源码属于基于 PostgreSQL 18 的 AlohaDB，尚未确认可直接用于原版 PostgreSQL。

### 核心用法

```sql
CREATE EXTENSION alohadb_search;
SELECT search_bm25(to_tsvector('english', 'database search'),
                   to_tsquery('english', 'database'));
SELECT search_fuzzy_match('alohadb', 'alohdb', 2);
SELECT search_haversine(21.3, -157.8, 21.31, -157.81);
```

### 运行边界

模块安装到 `public`，不可重定位，安装需要特权。源码未定义预加载钩子。`search_bm25()` 接受 `tsvector`、`tsquery` 及可选的 `k1`、`b`，此修订使用固定的语料规模和平均长度估计，并非新的索引访问方法或语料统计引擎。`search_edit_distance()` 与 `search_fuzzy_match()` 比较字符串，`search_phonetic()` 生成语音编码。`search_autocomplete()` 查询指定表和列；`search_expand_synonyms()` 读取 `alohadb_search_synonyms`，`search_analyze()` 提供文本分析。`search_haversine()` 返回米制距离，`search_nearby()` 按半径查询表。动态查询辅助函数会将表参数直接拼入 SQL，因此只应传入可信的表表达式，生产使用前须审查权限与实现。
