## 用法

来源：

- [官方上游文档](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/README.musicbrainz_unaccent.md)
- [官方扩展 SQL](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/sql/musicbrainz_unaccent.sql)
- [官方实现](https://github.com/metabrainz/postgresql-musicbrainz-unaccent/blob/b727896da12f03f404ed96a4d0cc8c967681fbaf/musicbrainz_unaccent.c)

`musicbrainz_unaccent` 通过 C 函数和文本搜索词典删除一组固定的 Unicode 重音符号。与 PostgreSQL 通用的 `unaccent` 函数不同，它的函数被声明为不可变，因此可用于表达式索引。它最初用于 MusicBrainz 规范化，其仓库现已归档。

### 核心流程

安装扩展，然后使用标量函数或文本搜索词典：

```sql
CREATE EXTENSION musicbrainz_unaccent;

SELECT musicbrainz_unaccent('ľščťžýáí');
SELECT ts_lexize('musicbrainz_unaccentdict', 'ľščťžýáí');

CREATE INDEX artist_unaccent_idx
  ON artist (musicbrainz_unaccent(name));
```

标量示例返回 `lsctzyai`。希望使用表达式索引的查询必须应用完全相同的不可变表达式。

### 重要对象

- `musicbrainz_unaccent(text)` 返回规范化文本，并声明为 `IMMUTABLE` 与 `STRICT`。
- `musicbrainz_unaccentdict_template` 是文本搜索词典模板。
- `musicbrainz_unaccentdict` 是安装后的词典；它会先将词元转换为小写，再应用编译进程序的映射。

### 兼容性与注意事项

无论数据库编码为何，实现都假定输入为 UTF-8，因此不得用于非 UTF-8 数据库。映射被编译到扩展二进制文件中，而不是从可编辑的词典规则文件加载，所以修改规范化规则需要重新构建。小写转换还取决于 PostgreSQL 的区域设置行为。上游项目已归档并依赖较旧的 PostgreSQL 内部接口，采用前应验证 ABI 兼容性、区域敏感字符、索引重建行为以及应用所需的每种语言。应固定已复核的源码与映射，避免规范化键意外变化。
