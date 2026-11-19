## 用法

来源：

- [官方资料](https://github.com/edisnord/albanian-stemmer/blob/1.2.1/README.md)
- [官方资料](https://github.com/edisnord/albanian-stemmer/blob/1.2.1/postgres/albanian_stemmer.control)
- [官方资料](https://github.com/edisnord/albanian-stemmer/blob/1.2.1/LICENSE)
- [官方资料](https://github.com/edisnord/albanian-stemmer/blob/1.2.1/postgres/albanian_stemmer--1.0.sql)

albanian_stemmer 提供阿尔巴尼亚语 Snowball 词典与全文检索配置。分发版本 1.2.1 保留扩展 SQL 版本 1.0。上游支持 PostgreSQL 14–18，创建扩展前需安装匹配的原生库。控制文件标记 trusted 且可迁移，无需预加载。

### 全文检索

```sql
CREATE EXTENSION albanian_stemmer;
SELECT to_tsvector('albanian', 'Sëmundjet e mollës dhe të pemëve');
SELECT plainto_tsquery('albanian', 'mollët');
SELECT ts_lexize('albanian_stem', 'mollëve');
```

albanian 配置复制 simple 解析，并将词语 token 类映射到 albanian_stem；数字、URL 等其他类别保留 simple 行为。词典折叠 ë/ç、移除停用词并剥离屈折变化。自定义词典可让 albanian_stem_template 使用 StopWords = false。

### 检索数据维护

生成文档向量与查询词时应使用同一配置。词干提取是有损规范化，不是完整词形还原器，也不保证语义相同。规则变化可能改变已存词项，采用新规则后应重新计算物化向量并重建受影响的索引。仅凭 SQL 版本 1.0 无法识别分发包中的词干算法修订。
