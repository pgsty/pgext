## 用法

来源：

- [官方控制文件](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag/rag.control)
- [官方 pgrag README](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/README.md)
- [官方包清单](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag/Cargo.toml)

`rag` 是已弃用 pgrag 项目的基础扩展。它提供本地文档转文本和按字符分块，以及多个远程嵌入、重排和聊天 API 的 HTTPS 客户端。项目已经停止开发和维护，因此只应在有明确迁移与安全方案时使用。

### 核心流程

先安装它声明的 `vector` 依赖，然后可在不安装模型扩展的情况下使用本地转换函数：

```sql
CREATE EXTENSION vector;
CREATE EXTENSION rag;

SELECT rag.markdown_from_html(
    '<html><body><h1>Title</h1><p>A short paragraph.</p></body></html>'
);

SELECT rag.chunks_by_character_count(
    'The quick brown fox jumps over the lazy dog', 20, 4
);
```

### 主要 API 分组

- 本地转换：`rag.markdown_from_html(text)`、`rag.text_from_pdf(bytea)` 和 `rag.text_from_docx(bytea)`。PDF 提取不支持 OCR 或复杂布局。
- 本地分块：`rag.chunks_by_character_count(text, integer, integer)` 返回相互重叠的文本块。
- OpenAI：`rag.openai_set_api_key(text)`、`rag.openai_get_api_key()`、嵌入辅助函数和 `rag.openai_chat_completion(json)`。
- Anthropic 与 Fireworks：对应的密钥设置/读取函数，以及 `rag.anthropic_messages(text, json)`、Fireworks 嵌入辅助函数和 `rag.fireworks_chat_completion(json)`。
- Voyage AI：密钥设置/读取、嵌入辅助函数，以及标量/数组形式的重排分数和距离函数。

### 运维说明

所核验版本为 0.0.0，不可重定位并依赖 `vector`；控制文件还设置 `superuser = true`。它未声明预加载要求。远程辅助函数会在数据库会话中发起外部 HTTPS 请求，因此延迟、供应商故障、配额和响应处理都会成为查询问题。密钥读取函数会返回已保存凭据；应限制函数权限，绝不能向不受信任角色开放。上游仓库已归档并弃用，必须钉住依赖。
