## 用法

来源：

- [官方v0.1.0 README](https://github.com/huangjimmy/pg_cjk_parser/blob/v0.1.0/Readme.md)
- [v0.1.0 发行说明](https://github.com/huangjimmy/pg_cjk_parser/releases/tag/v0.1.0)
- [v0.1.0 控制文件](https://github.com/huangjimmy/pg_cjk_parser/blob/v0.1.0/pg_cjk_parser.control)
- [v0.1.0 SQL 函数](https://github.com/huangjimmy/pg_cjk_parser/blob/v0.1.0/pg_cjk_parser--0.1.0.sql)

`pg_cjk_parser` 是一个基于内置解析器的 PostgreSQL 全文搜索解析器。在 UTF-8 数据库中，它对非 CJK 文本保持默认行为，同时为中文、日文和韩文文本生成重叠的 2-gram 令牌。该扩展安装了解析支持函数；您需要创建使用这些函数的文本搜索解析器和配置。

### 核心工作流程

```sql
CREATE EXTENSION pg_cjk_parser;

CREATE TEXT SEARCH PARSER public.pg_cjk_parser (
    START = prsd2_cjk_start,
    GETTOKEN = prsd2_cjk_nexttoken,
    END = prsd2_cjk_end,
    LEXTYPES = prsd2_cjk_lextype,
    HEADLINE = prsd2_cjk_headline
);

CREATE TEXT SEARCH CONFIGURATION public.config_2_gram_cjk (
    PARSER = public.pg_cjk_parser
);

SELECT alias, description, token
FROM ts_debug(
    'public.config_2_gram_cjk',
    'PostgreSQL 全文検索和中文检索'
);
```

显式在生成的 `tsvector` 列和查询中使用此配置，或将其设置为会话默认值：

```sql
SET default_text_search_config = 'public.config_2_gram_cjk';

SELECT to_tsvector('public.config_2_gram_cjk', '日本語全文検索');
```

### 重要对象

- `prsd2_cjk_start`, `prsd2_cjk_nexttoken`, `prsd2_cjk_end`, `prsd2_cjk_lextype`, 和 `prsd2_cjk_headline`：用于 `CREATE TEXT SEARCH PARSER` 的支持函数。
- `cjk_zht2zhs(text)`：将映射的传统中文字符转换为简体中文，而其他字符保持不变。
- 解析器标记类型 `cjk`：生成重叠的 CJK 大二元组；CJK 标点符号作为单个单元格发出。

```sql
SELECT cjk_zht2zhs('漢語');
-- 汉语
```

### 版本说明和注意事项

- v0.1.0 版本修复了 `0.1.0` 在混合宽度 UTF-8 字符中的错误扫描，并正确处理四字节 CJK 代码点。
- 上游支持从 PostgreSQL 11 到 18 的版本。
- 数据库必须使用 UTF-8 编码以实现 CJK 大二元组行为。使用其他编码时，解析器的行为类似于 PostgreSQL 默认解析器。
- 创建文本搜索解析器需要高级权限。您需要分别决定映射、词典、停用词和排名；示例配置仅定义了解析器。

`cjk_zht2zhs`
