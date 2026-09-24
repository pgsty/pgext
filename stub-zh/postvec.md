## 用法

来源：

- [README](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/README.md)
- [Control file / 控制文件](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/postvec.control)
- [postvec/Cargo.toml](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/Cargo.toml)
- [postvec/src/lib.rs](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/src/lib.rs)
- [postvec/LICENSE](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/LICENSE)

`postvec` 异步维护嵌入向量列，并融合全文搜索与向量搜索结果。本文对应固定提交的 0.2.0 开发源码，已发布的 0.1.0-1 制品属于另一版本边界。

### 启用

需要 PostgreSQL 16–18 和 `vector` 0.8 或更新版本。可配置已下载模型的内嵌 CPU 推理后端，或可访问的远程 UniVec 后端。将 `postvec` 加入 `shared_preload_libraries`，配置 `postvec.database` 后重启；远程模式还使用 `postvec.grpc_endpoints` 与 `postvec.http_endpoints`。

```sql
CREATE EXTENSION postvec CASCADE;
SELECT postvec.refresh_models();
```

### 核心工作流

```sql
CREATE TABLE docs (id bigserial PRIMARY KEY, body text);
SELECT postvec.enable('public.docs', 'body',
  model => 'sentence-transformers-all-minilm-l6-v2', create_fts_index => true);
INSERT INTO docs(body) VALUES ('PostgreSQL extension development');
SELECT * FROM postvec.status();
SELECT d.body, s.rrf_score
FROM postvec.search('public.docs', 'body', 'database extensions') s
JOIN docs d ON d.id = s.pk_value::bigint ORDER BY s.rrf_score DESC;
```

### 对象与边界

`postvec.enable` 添加影子向量列、触发器和回填任务；需等待队列完成后才能获得完整搜索结果。`postvec.create_vector_index` 创建向量索引，自动和立即创建路径会阻塞写入，繁忙表宜另行并发建索引。`postvec.adopt` 接管已有向量，`postvec.disable` 取消管理。模型转换需要兼容的转换器，不能视为普遍无损。

管理操作需要表所有者或超级用户权限。工作进程以引导超级用户运行并绕过 RLS，因此启用表前需检查上游的行安全限制和授权要求。远程推理会把原文发送到数据库主机之外。扩展采用 PostgreSQL 许可证，独立服务端和下载模型各有其许可条款。
