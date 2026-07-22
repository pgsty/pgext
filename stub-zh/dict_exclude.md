## 用法

来源：

- [官方 README](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/README.md)
- [扩展 SQL](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/dict_exclude--1.0.sql)
- [词典实现](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/dict_exclude.c)

`dict_exclude` 是一个原型全文搜索词典，用管理员提供的正则表达式丢弃匹配的词元。它适合在普通词干词典之前过滤有规律的 OCR 噪声或应用特定的词元族。

### 核心流程

在 PostgreSQL 的 `tsearch_data` 目录下安装名为 `exclude.rules` 的服务端规则文件。文件每行包含一个 PostgreSQL 正则表达式，例如：

```text
^abc$
^def[0-9]+$
```

创建扩展、复制现有配置，并把 `dict_exclude` 放在映射首位，使不匹配的词元继续交给 `english_stem`。

```sql
CREATE EXTENSION dict_exclude;

CREATE TEXT SEARCH CONFIGURATION ocr_filtered
    (COPY = pg_catalog.english);

ALTER TEXT SEARCH CONFIGURATION ocr_filtered
    ALTER MAPPING FOR asciihword, asciiword
    WITH dict_exclude, english_stem;

SELECT to_tsvector(
    'ocr_filtered',
    'fat abc cat def123 visible'
);
```

### 对象与规则行为

- `dict_exclude_template` 是全文搜索模板。
- `dict_exclude` 是由该模板创建的词典。
- `dict_exclude_init` 加载并编译规则；`dict_exclude_lexize` 对匹配项返回空词位集合，并把不匹配项传给下一个词典。
- 文件基本名固定为 `exclude`，因此资源名只能是 `exclude.rules`；词典没有选择其他规则集的选项。

### 运维注意事项

规则文件从数据库服务器安装树读取，必须由管理员部署。文件缺失或规则非法都会导致词典初始化失败。编译后的规则会缓存在后端词典状态中，因此应在新会话验证变更，并根据服务器版本按需重建或重新加载受影响的全文搜索对象。

每个词元都会依次匹配正则列表。过宽或代价过高的表达式可能删除合法词语或消耗大量 CPU，生成的 `tsvector` 也会改变搜索语义。上游明确把该扩展标为仅供开发环境使用的原型，并只记录了 PostgreSQL 9.4 时代的用法；采用前应测试规则、性能、编码和目标版本兼容性。
