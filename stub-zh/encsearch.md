## 用法

来源：

- [README.md](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/README.md)
- [src/lib.rs](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/src/lib.rs)
- [Cargo.toml](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/Cargo.toml)
- [META.json](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/META.json)
- [encsearch.control](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/encsearch.control)

`encsearch` 将加密文本与带密钥的搜索词元分别存入密文字段和 `int8[]` 字段。实验性的 0.1.0 版本可借助 GIN 索引进行精确、单词及模式搜索。上游已测试 PostgreSQL 17，Cargo 中声明的其他版本目标尚未测试。

### 核心用法

```ini
session_preload_libraries = 'encsearch'
```

```sql
CREATE EXTENSION encsearch;
CREATE TABLE encrypted_notes (id bigint GENERATED ALWAYS AS IDENTITY, body bytea, idx int8[]);
CREATE INDEX encrypted_notes_tokens ON encrypted_notes USING gin (idx);
INSERT INTO encrypted_notes (body, idx)
VALUES (enc_encrypt('Patient has diabetes'), enc_index('Patient has diabetes'));
SELECT enc_decrypt(body) FROM encrypted_notes
WHERE enc_like(body, idx, '%diab%');
```

### 运行边界

此扩展不受信任，需要管理员安装。应先配置 `session_preload_libraries`，再建立客户端连接并设置 `encsearch.key`，否则提前设置的密钥会成为可见的占位参数。执行数据示例前，应向会话或显式接收密钥的函数重载提供安全保存的密钥，长度至少为 32 个字符。`enc_gen_key()` 可生成 64 个字符的密钥，丢失密钥后将无法恢复明文。

`enc_encrypt` 和 `enc_decrypt` 处理文本与密文；`enc_index` 生成需存储的词元；`enc_like`、`enc_ilike` 和 `enc_eq` 先筛选候选项，再解密核对。`enc_has_words` 仅检查词元包含关系，不解密；`enc_q_like`、`enc_q_exact` 和 `enc_q_words` 返回原始查询词元。`enc_algorithm` 和 `enc_algorithms` 用于查看加密算法。

`encsearch.algorithm` 默认为 `aes-256-gcm`，确定性的 `aes-256-siv` 会暴露重复明文。建立和查询词元时，`encsearch.token_hash` 必须一致：默认使用 `hmac-sha256`，也可选 `blake3`。更换词元哈希算法需要重新生成词元；更换密钥需要先用旧密钥解密，再用新密钥重建密文和词元。

盲索引会暴露相同值、单词、字符片段及其出现频率。密钥会到达数据库服务器，也可能进入语句或参数日志。这项实验性功能是在隐私与搜索性能之间取舍，无法防范已持有密钥的服务器。
