## 用法

来源：

- [官方 README](https://github.com/metabrainz/postgresql-musicbrainz-collate/blob/6c350bcbfd3418db4c90826eeea22c24146b2d05/README.musicbrainz_collate.md)
- [扩展 SQL](https://github.com/metabrainz/postgresql-musicbrainz-collate/blob/6c350bcbfd3418db4c90826eeea22c24146b2d05/sql/musicbrainz_collate.sql)
- [ICU 排序键实现](https://github.com/metabrainz/postgresql-musicbrainz-collate/blob/6c350bcbfd3418db4c90826eeea22c24146b2d05/musicbrainz_collate.c)

`musicbrainz_collate` 把 UTF-8 文本转换为以 `bytea` 表示的 ICU Unicode 排序算法排序键。当需要复现这个已归档扩展所实现的 MusicBrainz 风格、区域无关且识别数字的顺序时，可在 `ORDER BY` 或表达式索引中使用该键。

### 核心流程

```sql
CREATE EXTENSION musicbrainz_collate;

SELECT name
FROM artist
ORDER BY musicbrainz_collate(name);

CREATE INDEX artist_musicbrainz_name_idx
    ON artist (musicbrainz_collate(name));

SELECT name
FROM artist
ORDER BY musicbrainz_collate(name)
LIMIT 50;
```

唯一的公开函数是 `musicbrainz_collate(text) RETURNS bytea`。它是严格函数并标记为不可变，因此 PostgreSQL 可在表达式索引中使用它。实现打开 ICU 的根或默认排序器并启用数字排序，使数字串按数值而不是逐字符排序。

### 编码与索引维护

函数假设输入字节为 UTF-8，在采用其他服务器编码的数据库中没有实用价值。使用前应确认 `SHOW server_encoding`。

即使 SQL 函数声明为不可变，ICU 版本变化仍可能改变生成的排序键。升级 ICU 库后，要对所有基于该函数的表达式索引执行 `REINDEX INDEX artist_musicbrainz_name_idx`，并重新验证应用可见顺序。副本和故障切换节点也必须使用兼容的 ICU 构建。

### 兼容边界

历史 README 要求 PostgreSQL 8.3 或更高版本以及 ICU 3.8 或更高版本，但仓库已经归档，也没有现代 PostgreSQL 与 ICU 的兼容矩阵。采用遗留行为前，应在目标 PostgreSQL 版本上与内置 ICU 排序规则比较结果。
