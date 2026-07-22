## 用法

来源：

- [官方 README](https://github.com/ontaev/kalmyk-hunspell/blob/92f10f374186899e658ef06fd823adae089efaee/README.md)
- [官方扩展控制文件](https://github.com/ontaev/kalmyk-hunspell/blob/92f10f374186899e658ef06fd823adae089efaee/hunspell_xal.control)

`hunspell_xal` 为 PostgreSQL 全文搜索安装卡尔梅克语 Hunspell 词典。其词表基于 B. D. Muniev 主编的 1977 年《卡尔梅克语—俄语词典》，当前审阅的上游版本只支持名词词缀。

### 核心流程

安装扩展文件并创建 `hunspell_xal` 后，先检查软件包安装的对象，再把它的 Hunspell 词典接入卡尔梅克语文本搜索配置。典型的 PostgreSQL 配置流程如下：

```sql
CREATE EXTENSION hunspell_xal;

SELECT dictname
FROM pg_ts_dict
WHERE dictname LIKE '%xal%';

CREATE TEXT SEARCH CONFIGURATION xal (COPY = simple);
ALTER TEXT SEARCH CONFIGURATION xal
  ALTER MAPPING FOR word, asciiword
  WITH hunspell_xal, simple;
```

建立索引前，应使用有代表性的卡尔梅克语文本验证分词和规范化结果：

```sql
SELECT to_tsvector('xal', 'кальмг келн');
SELECT to_tsquery('xal', 'кальмг');
```

### 语言与运维边界

- 上游 README 明确说明此版本只支持名词词缀，动词词缀仍是后续工作；不要假定它能完整处理卡尔梅克语形态变化。
- 词典质量取决于随附的 `.aff`、`.dic` 数据以及文本搜索配置映射的词元类型。应由领域审阅者测试预期词形、专有名词、拼写变体和误匹配。
- 更新词典文件不会改写现有的已存储 `tsvector` 值。搜索语义变化后，应重建生成向量及相关索引。
- 应确认词典数据来源及许可证符合预期的再分发与部署场景。
