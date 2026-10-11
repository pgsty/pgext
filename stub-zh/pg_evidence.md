## 用法

来源：

- [README.md](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/README.md)
- [docs/api.md](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/docs/api.md)
- [docs/operations.md](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/docs/operations.md)
- [src/lib.rs](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/src/lib.rs)
- [Cargo.toml](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/Cargo.toml)
- [pg_evidence.control](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/pg_evidence.control)

`pg_evidence` 0.1.0 是开发预览版，用于存储文档版本、检索有界摘录，并将引用解析到准确的 UTF-8 字节区间。要求 PostgreSQL 18 与 UTF8 数据库。向量由客户端生成；只有启用嵌入的集合才需要 `vector`。

### 核心用法

```sql
CREATE EXTENSION pg_evidence;
SELECT evidence.init_collection('docs', '{}'::jsonb);
SET statement_timeout = '5s';
SELECT evidence.query('docs', '{
  "nodes":[{"id":"hits","op":"literal","text":"retrieval","limit":5}],
  "output":"hits","excerpt_bytes":512
}'::jsonb);
```

### 运行边界

`evidence.stage_version()` 暂存原文、摘要与字节区间，`evidence.publish_version()` 将版本设为当前版本。`evidence.query()` 支持字面量、正则、全文及可选向量检索，`evidence.resolve()` 校验保留的引用字节。`evidence.annotate()` 管理标签和关系，`evidence.purge()` 逻辑移除保留的原文。查询必须设置非零 `statement_timeout`：输出预算并不限制数据库内部工作量。安装需要超级用户；函数使用调用者权限，应用角色需要模式和语料表授权。受信任的语料写入者可通过直接 SQL 绕过 API 不变量。删除扩展会保留语料表；清理不会擦除 WAL、备份及其他副本。上游仍待选择许可证并完成原生 x86-64 验证，尚不能视为生产就绪。
