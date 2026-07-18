## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/CyberDem0n/bg_mon/blob/d215e94f88a9088bdde0ce54f74d7d8bb2deb8e0/README.md)
- [后台工作进程与 HTTP 实现](https://github.com/CyberDem0n/bg_mon/blob/d215e94f88a9088bdde0ce54f74d7d8bb2deb8e0/bg_mon.c)
- [MIT 许可证](https://github.com/CyberDem0n/bg_mon/blob/d215e94f88a9088bdde0ce54f74d7d8bb2deb8e0/LICENSE)

`bg_mon` 是无 SQL 对象的后台工作进程模块：无需创建 SQL 扩展。将动态库加入 `shared_preload_libraries`，在 `postgresql.conf` 中配置，然后重启 PostgreSQL。

```conf
shared_preload_libraries = 'bg_mon'
bg_mon.listen_address = '127.0.0.1'
bg_mon.port = 8080
bg_mon.naptime = 1
bg_mon.history_buckets = 20
```

工作进程汇总后端活动、进程、内存、系统、网络、磁盘和分区统计，通过 `/` 提供当前 JSON、通过 `/<bucket>` 提供历史桶，并在 `/ui` 提供简单浏览器界面。构建时必须有 libevent，Brotli 历史压缩则可选。

嵌入式 HTTP 服务器没有文档化的认证或 TLS。应保留默认回环地址；确需远程访问时，在前方部署带认证的代理。把 `bg_mon.listen_address` 绑定为 `0.0.0.0` 会向网络暴露数据库活动与主机遥测。应在代表性负载下验证采集开销、端点权限和历史记录占用内存。
