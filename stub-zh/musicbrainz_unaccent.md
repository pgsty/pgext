## 用法

来源：

- [官方扩展控制文件](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/musicbrainz_unaccent.control)
- [官方上游文档](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/README.musicbrainz_unaccent.md)

`musicbrainz_unaccent` — musicbrainz_unaccent：为 PostgreSQL 提供可用于索引的 Unicode 去重音函数和文本搜索字典。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "musicbrainz_unaccent";
SELECT extversion
FROM pg_extension
WHERE extname = 'musicbrainz_unaccent';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
