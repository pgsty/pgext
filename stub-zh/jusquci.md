## 用法

来源：

- [上游 README](https://github.com/thjbdvlt/jusquci/blob/9238c91a290c8b0e1d036a53a148caf56ac1c2b6/README.md)
- [扩展 control 文件](https://github.com/thjbdvlt/jusquci/blob/9238c91a290c8b0e1d036a53a148caf56ac1c2b6/postgresql/jusquci.control)
- [文本搜索对象](https://github.com/thjbdvlt/jusquci/blob/9238c91a290c8b0e1d036a53a148caf56ac1c2b6/postgresql/jusquci--1.0.sql)
- [PostgreSQL 构建与停用词安装](https://github.com/thjbdvlt/jusquci/blob/9238c91a290c8b0e1d036a53a148caf56ac1c2b6/postgresql/makefile)

`jusquci` 提供面向当代法语的 PostgreSQL 文本搜索解析器。其分词器可识别省音、复合及包容性写法、连续标点、URL、表情符号及相关现代文本模式。安装扩展会创建 `jusquci` 解析器、`jusquci_stop` 词典和 `jusquci` 文本搜索配置。

### 解析法语文本

```sql
CREATE EXTENSION jusquci;

SELECT to_tsvector(
  'jusquci',
  'le quotidien,s''invente-t-il par mille manière de braconner???'
);

SELECT *
FROM ts_parse('jusquci', 'peut-être--là lecteur-rice-x-s');
```

必须把独立的 `french_jusquci.stop` 文件复制到 PostgreSQL 文本搜索数据目录。上游构建通过 `make install install_stop` 完成这一步；缺少该文件时，创建 `jusquci_stop` 可能失败。

版本 `1.0` 以 C 解析器实现，并要求由超级用户创建扩展。上游说明仅在 Debian Linux 与 PostgreSQL 16 上测试过，因此替换现有法语配置前，应验证 token 类型、停用词行为、搜索排序、编码以及目标版本兼容性。
