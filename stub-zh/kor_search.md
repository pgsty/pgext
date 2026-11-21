## 用法

来源：

- [README.md](https://github.com/junhkang/kor_search/blob/v2.0.0/localization/en/README.md)
- [CHANGELOG.md](https://github.com/junhkang/kor_search/blob/v2.0.0/CHANGELOG.md)
- [kor_search.control](https://github.com/junhkang/kor_search/blob/v2.0.0/kor_search.control)
- [kor_search--2.0.0.control](https://github.com/junhkang/kor_search/blob/v2.0.0/kor_search--2.0.0.control)
- [kor_search--1.1.0--2.0.0.sql](https://github.com/junhkang/kor_search/blob/v2.0.0/sql/kor_search--1.1.0--2.0.0.sql)

`kor_search` 2.0.0 使用 Rust 与 Lindera/ko-dic 分析韩英混合文本，生成 PostgreSQL tsvector，并通过库内词典扩展查询别名。上游在 Linux、UTF-8 数据库及 PostgreSQL 16–18 上测试，运行时不调用外部服务。

### 索引检索

安装匹配的本地库文件后，以超级用户创建扩展：

```sql
CREATE EXTENSION kor_search;
CREATE TABLE documents (body text NOT NULL);
INSERT INTO documents VALUES ('삼성에서 PostgreSQL을 사용한다');
CREATE INDEX documents_search ON documents USING gin (kor_vector_v1(body));
SELECT body FROM documents
WHERE kor_vector_v1(body) @@ kor_search_tsquery('PostgreSQL', 'default');
SELECT * FROM kor_analyze_v1('삼성에서 PostgreSQL을 사용한다');
```

`kor_vector_v1()` 生成带位置信息的 tsvector，可用普通 GIN 建索引。`kor_analyze_v1()` 展示分词结果。`kor_search_tsquery(query, scope)` 优先扩展最长词典短语，在别名间使用 OR、词组间使用 AND；分析后没有词元的输入返回 NULL。存储生成列可减少昂贵的短语复核。

### 别名词典

向 `kor_search_aliases(scope, concept, term, kind)` 添加领域词，或加载可选的示例词典。每个作用域内，规范化词项只能属于一个概念。修改别名无需重建文档索引。scope 只选择词典，不限制文档或用户权限；检索用户需要模式 USAGE 与词典 SELECT 权限。

分析器不提供任意子串检索、自动翻译或英语词干提取。限制为每篇文档 512 KiB，每次查询 4 KiB 或 64 个分析词元，每个别名八个词元，每个概念 32 项扩展。

### 升级与权限

新装 2.0 不要求 `pg_trgm`，未使用的旧三元组索引已在 1.1 删除。四个旧检索函数保留，新应用宜使用索引接口。1.1 迁移修复了初版的词典 SQL 插值与固定缓冲区处理。

安装新文件，在维护窗口重连已有会话，然后执行：

```sql
ALTER EXTENSION kor_search UPDATE TO '2.0.0';
```

经过测试的升级路径会途经 1.1.0。应检查并明确将旧领域词复制到新词典，旧种子行的编辑和删除须另行备份。2.0 允许安装时指定模式，但之后不能重定位；版本专用控制文件覆盖根控制文件。无需预加载；禁止安装本地模块的托管服务无法使用此构建。
