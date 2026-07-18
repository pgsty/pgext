## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/isdaniel/rust_pg_extensions/blob/f25df80c4ce06bf81bb7e3a58fc5087ea85c64ac/README.md)
- [工具实现](https://github.com/isdaniel/rust_pg_extensions/blob/f25df80c4ce06bf81bb7e3a58fc5087ea85c64ac/src/utility_lib.rs)
- [网络实现](https://github.com/isdaniel/rust_pg_extensions/blob/f25df80c4ce06bf81bb7e3a58fc5087ea85c64ac/src/networking_lib.rs)

`all_in_one_lib` 是一个 pgrx 实验，把哈希辅助函数、AES-GCM 加解密、服务器网络发现、逻辑解码输出插件和示例外部数据包装器合并在一个仅超级用户可安装的扩展中。

```sql
CREATE EXTENSION all_in_one_lib;
SELECT compute_hash('payload', 'sha256');
SELECT get_server_hostname(), get_server_ip();
```

不要用 `data_encrypt` 保护真实机密。已复核源码对每次 AES-256-GCM 操作使用同一个硬编码 nonce；同一密钥下复用 nonce 会破坏 GCM 安全性。它还从 SQL 接收原始密钥，可能经查询、日志或监控泄漏。已有密文需要单独设计迁移，不能继续沿用。

网络函数会暴露主机标识与非回环地址。解码器只实现少量元组类型且含未完成路径，FDW 也只是演示。应从 `PUBLIC` 撤销所有函数，只在逐项审计后授权，且不要在生产安装这个混合实验库。各项需求应选择专用、持续维护的扩展和外部密钥管理。
