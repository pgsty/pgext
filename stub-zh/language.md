## 用法

来源：

- [0.0.4 版本 SQL 对象](https://github.com/adjust/pg-language/blob/36e553fec17c126be3fed68b28c8c13f7f88299a/language--0.0.4.sql)
- [生成语言映射](https://github.com/adjust/pg-language/blob/36e553fec17c126be3fed68b28c8c13f7f88299a/languages)
- [PGXN README](https://pgxn.org/dist/language/0.0.4/README.html)

`language` 定义了紧凑的、单字节的应用专用 `language` 类型。它接受的值由上游 `languages` 映射生成，与 PostgreSQL 过程语言目录项无关。

```sql
CREATE EXTENSION language;
SELECT supported_languages();
CREATE TABLE profile (
  profile_id bigint PRIMARY KEY,
  preferred_language language NOT NULL
);
```

扩展提供文本/二进制输入输出、比较操作符，以及默认 B-tree 与 hash 操作符类，因此值可以排序、索引和比较。应通过 `supported_languages()` 发现安装构建实际携带的映射，不要假定它覆盖某套 ISO 代码。

内部表示是生成的一字节代码。修改或重新生成映射可能重新解释已存储字节，因此二进制兼容性与语义都依赖精确构建。所有节点应固定同一源码版本，升级前验证转储恢复与复制行为；未审计映射前，不要把该类型当作通用标准注册表。
