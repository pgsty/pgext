## 用法

来源：

- [上游 README 与概念验证状态](https://github.com/Hendler/flame/blob/763f9c9ba588187569fa35dad4f9090921fbd24c/README.md)
- [扩展 control 文件](https://github.com/Hendler/flame/blob/763f9c9ba588187569fa35dad4f9090921fbd24c/flame.control)
- [Cargo 软件包元数据](https://github.com/Hendler/flame/blob/763f9c9ba588187569fa35dad4f9090921fbd24c/Cargo.toml)
- [模型初始化](https://github.com/Hendler/flame/blob/763f9c9ba588187569fa35dad4f9090921fbd24c/src/fast_embeddings.rs)

`flame` `0.0.1` 版是用 Rust/pgrx 编写的文本嵌入概念验证函数。`fast_embed(text)` 使用 fastembed 的 `AllMiniLML6V2` 模型，并以 PostgreSQL real 数组返回嵌入。

### 示例

```sql
CREATE EXTENSION flame;
SELECT array_length(fast_embed('hello world'), 1);
SELECT fast_embed('PostgreSQL text embedding');
```

模型在每个后端进程中初始化一次，可能需要以 PostgreSQL 操作系统账户下载/缓存模型工件。这会带来冷启动延迟、网络供应链暴露、每连接内存开销和不可预测的运维失败。README 把版本称为概念验证并列出缺失测试，其预期维度也与实际选择的模型不一致；初始化失败可能 panic，嵌入失败则返回空数组。应固定并预置模型文件、限制执行、约束输入/并发，并在保存向量前验证输出长度和版本。
