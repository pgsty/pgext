## 用法

来源：

- [官方 README](https://github.com/za-arthur/dict_translate/blob/c95381026d760bdfae6a8cb2bab94c7ca9c06575/README.md)
- [官方扩展 SQL](https://github.com/za-arthur/dict_translate/blob/c95381026d760bdfae6a8cb2bab94c7ca9c06575/dict_translate--1.0.sql)

`dict_translate` 1.0 提供 `translate` 全文搜索词典模板。它先用另一个词典规范化输入 token，再将其展开为服务端 `.trn` 文件中配置的翻译词组，从而让译词作为等价 lexeme 参与检索。

### 核心流程

在所有可能执行该配置的数据库服务器上创建翻译文件，例如 `$SHAREDIR/tsearch_data/test_trn.trn`：

```text
forest wald forst holz
home haus heim
```

安装扩展，并围绕该文件创建词典与搜索配置：

```sql
CREATE EXTENSION dict_translate;

CREATE TEXT SEARCH DICTIONARY test_trn (
  TEMPLATE = translate,
  DictFile = test_trn,
  InputDict = pg_catalog.english_stem
);

CREATE TEXT SEARCH CONFIGURATION test_cfg (COPY = simple);

ALTER TEXT SEARCH CONFIGURATION test_cfg
  ALTER MAPPING FOR asciiword, asciihword, hword_asciipart,
                    word, hword, hword_part
  WITH test_trn, english_stem;

SELECT to_tsvector('test_cfg', 'homes forest haus');
SELECT to_tsvector('test_cfg', 'homes forest haus')
       @@ to_tsquery('test_cfg', 'home');
```

`DictFile` 使用不带 `.trn` 后缀的翻译文件名。每行以源 lexeme 开头，后接需要一同输出的 lexeme。`InputDict` 指定在查找翻译前用于规范化输入的词典。

### 重要对象

- `translate` 是扩展安装的全文搜索模板。
- `dtrn_init(internal)` 根据 `DictFile` 和 `InputDict` 选项初始化词典。
- `dtrn_lexize(internal, internal, internal, internal)` 为全文搜索子系统执行规范化与翻译展开。

### 运维说明

`.trn` 文件属于文件系统配置，不是数据库数据；应在主库、备库和替换节点之间一致部署并纳入版本管理。修改文件不会重写既有 `tsvector` 值，因此翻译规则变化后必须刷新已存向量或索引。应使用可信的规范化词典，并同时测试索引配置与查询配置中的词干行为。
