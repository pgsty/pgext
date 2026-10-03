## 用法

来源：

- [README.md](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/README.md)
- [pg_kiwi.control](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/pg_kiwi.control)
- [pg_kiwi--1.0.sql](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/pg_kiwi--1.0.sql)
- [pg_kiwi.c](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/pg_kiwi.c)
- [变更记录](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/CHANGELOG.md)
- [Kiwi 0.22.2 license](https://raw.githubusercontent.com/bab2min/Kiwi/v0.22.2/LICENSE)

`pg_kiwi` 1.0 将 Kiwi 韩文形态分析器接入 PostgreSQL 全文检索。上游源码说明面向 PostgreSQL 18 与 Kiwi 0.22.2，服务器还须能够读取 Kiwi 模型文件。

### 基本用法

```sql
CREATE EXTENSION pg_kiwi WITH SCHEMA public;
CREATE TABLE korean_docs (
 id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
 body text,
 terms tsvector GENERATED ALWAYS AS
   (to_tsvector('public.korean', body)) STORED
);
CREATE INDEX ON korean_docs USING gin (terms);
SELECT id FROM korean_docs
WHERE terms @@ to_tsquery('public.korean', '불국사');
```

### 配置与维护

扩展安装 `kiwi_parser` 全文解析器及 `korean` 配置，将词元类别映射到简单词典。使用生成的 `tsvector` 列并创建 GIN 索引，可避免每次查询都重新分析所有文档。

`pg_kiwi.model_path` 指定模型文件，支持重载；`pg_kiwi.pos_filter` 按会话控制词性过滤。修改分词方式或模型后，须重新评估已存向量并重建受影响的检索数据。创建扩展需要超级用户；该扩展未声明共享预加载要求。README 还演示了可选的 `pg_textsearch` BM25 集成，后者有独立的兼容性与预加载要求，基本 GIN 用法不依赖它。

control 版本仍为 1.0；变更记录中的 1.1.0 调整的是捆绑的 BM25 镜像，而非 Kiwi 解析器。扩展源码采用 MIT 许可证，动态链接的 Kiwi 0.22.2 库采用 LGPL-2.1-or-later 许可证。
