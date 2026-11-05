## 用法

来源：

- [上游来源 1](https://github.com/d-d-dev/pg_fuzzysearch/blob/3460d9769515fc03d554493a2aaba438611469a7/README.md)
- [上游来源 2](https://github.com/d-d-dev/pg_fuzzysearch/blob/3460d9769515fc03d554493a2aaba438611469a7/pg/pg_fuzzysearch.control)

`pg_fuzzysearch` 0.2.0 实现 fuzzysort 风格的子序列匹配与预处理模糊匹配目标。

### 核心流程

`fuzzy_prepare` 返回可在生成列中复用的 `fuzzy_target`。`fuzzy_match` 可指定最低得分，`fuzzy_score` 在无匹配时返回 NULL，空查询得分为零。规范化辅助函数使用与匹配相同的 Unicode 规范化方式。

```sql
CREATE EXTENSION pg_fuzzysearch;
SELECT fuzzy_match('pg', 'PostgreSQL');
SELECT fuzzy_score('pg', fuzzy_prepare('PostgreSQL'));
```

### 运行边界

仅支持 PostgreSQL 17 和 UTF-8 数据库。匹配基于子序列而非编辑距离。扩展提供扫描函数，没有索引访问方法；预处理目标会增加存储，需要按负载测量效果。已核对源码未明确声明本项目自身许可证，虽注明 fuzzysort 依赖采用 MIT 许可。
