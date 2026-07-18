## 用法

来源：

- [上游固定版本 README](https://github.com/yihong0618/pg_opendal/blob/71f78ac7414c82425578060a5d1d233ce696a957/README.md)
- [固定版本 Rust 实现](https://github.com/yihong0618/pg_opendal/blob/71f78ac7414c82425578060a5d1d233ce696a957/src/lib.rs)
- [固定版本 Cargo 元数据](https://github.com/yihong0618/pg_opendal/blob/71f78ac7414c82425578060a5d1d233ce696a957/Cargo.toml)
- [固定版本扩展控制文件](https://github.com/yihong0618/pg_opendal/blob/71f78ac7414c82425578060a5d1d233ce696a957/pg_opendal.control)

pg_opendal 0.0.0 把部分 Apache OpenDAL 存储操作暴露为 SQL 函数。所审构建启用了本地文件系统、S3 与内存服务，提供读取、写入、存在检查、删除、元数据、建目录、列目录、复制、重命名和能力查询。

### 隔离的文件系统示例

先创建一个归服务器账户所有、可丢弃的目录，再只使用相对该根目录的路径：

```sql
CREATE EXTENSION pg_opendal;

SELECT pg_opendal_write(
    'fs',
    'demo.txt',
    'hello',
    '{"root":"/tmp/pg-opendal-sandbox"}'::jsonb
);

SELECT pg_opendal_read(
    'fs',
    'demo.txt',
    '{"root":"/tmp/pg-opendal-sandbox"}'::jsonb
);
```

服务配置必须是所有值均为字符串的 JSON 对象。读取函数返回 UTF-8 文本，而不是任意二进制。

### 文件系统、凭据与后端风险

每次调用都会新建 OpenDAL 操作器和 Tokio 运行时，然后阻塞 PostgreSQL 后端直到存储 I/O 完成。读取会把整个对象装入内存并要求有效 UTF-8；写入接收完整文本值；列目录还会对每个条目额外执行一次 stat。大对象、慢端点、重试和大目录会消耗后端内存与连接时间。

本地文件系统服务以 PostgreSQL 操作系统账户运行，并包含写入、删除、复制和重命名等破坏性操作。S3 配置把访问密钥和秘密密钥作为 SQL 参数传递，凭据可能暴露于查询文本、活动视图、日志、审计系统和错误报告。生成函数没有应用授权模型。应撤销 PUBLIC 执行权；只有在审查 search_path 后才通过固定服务的 SECURITY DEFINER 封装开放；尽量不要让凭据出现在调用者 SQL 中，并设置文件系统与网络沙箱。

上游明确标注项目仍在开发。这个 0.0.0 快照使用 pgrx 0.14.3，只声明 PostgreSQL 13 至 17，不含 18。任何非一次性用途前都要测试取消、超时、部分写入、路径穿越、符号链接、端点重定向、凭据轮换、后端崩溃和服务商一致性。
