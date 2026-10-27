## 用法

来源：

- [Extension versions](https://docs.databricks.com/aws/en/oltp/projects/extensions)
- [Enable Lakebase Search](https://docs.databricks.com/aws/en/oltp/projects/lakebase-search)
- [Extension guide](https://docs.databricks.com/aws/en/oltp/projects/lakebase-tokenizer)

`lakebase_tokenizer` 为托管 Lakebase Postgres 增加整词处理和中日韩形态分析词典。供应商为 PostgreSQL 16–18 列出的版本为 0.2.0。

### 分词与搜索

先在项目设置中启用 Lakebase Search。该操作不可撤销，会重启项目的全部计算节点并断开连接。扩展可重定位，可安装到数据库角色有权管理的模式中：

```sql
CREATE SCHEMA IF NOT EXISTS tokenizer_ext;
CREATE EXTENSION IF NOT EXISTS lakebase_tokenizer WITH SCHEMA tokenizer_ext;
SELECT to_tsvector('tokenizer_ext.tokenizer_chinese', '现代分布式文件系统架构');
```

`tokenizer_wholeword` 支持规范化、停用词、同义词和英文词干化。`tokenizer_morph` 提供中日韩分析；内置配置包括 `tokenizer_chinese`、`tokenizer_japanese`、`tokenizer_korean`。输出向量可用于 PostgreSQL GIN 或 `lakebase_text` BM25 索引；`lakebase_text` 为可选功能。文档与查询应使用一致配置，混合模式中文配置使用 `plainto_tsquery`。

### 更新

`ALTER EXTENSION lakebase_tokenizer UPDATE` 更新扩展，不会重新生成已存储的向量或重建索引。文档明确说明 0.1.x→0.2.0 保持分词输出和既有配置，无需执行这两项操作。后续版本改变输出时，应遵循对应版本的维护要求。供应商文档没有确认实现语言、许可证和预加载设置。
