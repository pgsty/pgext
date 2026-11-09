## 用法

来源：

- [Official README](https://github.com/sweatybridge/pg_webrtc/blob/80ceefdf29b077bbbece269842d84108b3f9d6a7/README.md)
- [Extension implementation](https://github.com/sweatybridge/pg_webrtc/blob/80ceefdf29b077bbbece269842d84108b3f9d6a7/src/lib.rs)
- [Extension control file](https://github.com/sweatybridge/pg_webrtc/blob/80ceefdf29b077bbbece269842d84108b3f9d6a7/pg_webrtc.control)
- [Package metadata](https://github.com/sweatybridge/pg_webrtc/blob/80ceefdf29b077bbbece269842d84108b3f9d6a7/Cargo.toml)

`pg_webrtc` 0.1.0 是面向 PostgreSQL 18 的早期预加载扩展，将 pgrx 0.19.2 与 `webrtc-rs` 连接起来，并建立后台工作进程边界。当前唯一 SQL 面是构建信息：

```sql
CREATE EXTENSION pg_webrtc;
SELECT webrtc.version();
SELECT webrtc.engine();
```

只有在服务器启动阶段加载库时，扩展才会注册该工作进程。使用时需将其加入 `shared_preload_libraries` 并重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_webrtc'
```

### 当前边界

工作进程目前持有 Tokio runtime 并保持空闲；尚未提供 peer connection、offer/answer、ICE candidate 或 data-channel SQL 命令。扩展有意设为仅超级用户、不可重定位，并且未来可能打开网络套接字。应将其视为运行时/进程边界原型，而不是可用的 WebRTC API。上游仓库未声明许可证，再分发前必须解决该问题。
