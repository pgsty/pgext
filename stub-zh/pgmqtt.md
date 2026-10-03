## 用法

来源：

- [pgmqtt 0.5.1 README](https://github.com/RayElg/pgmqtt/blob/0.5.1/README.md)
- [pgmqtt 0.5.1 interfaces](https://github.com/RayElg/pgmqtt/blob/0.5.1/docs/interfaces.md)
- [pgmqtt 0.5.1 configuration](https://github.com/RayElg/pgmqtt/blob/0.5.1/docs/configuration.md)
- [pgmqtt 0.5.1 limitations](https://github.com/RayElg/pgmqtt/blob/0.5.1/docs/limitations.md)
- [pgmqtt 0.5.1 release notes](https://github.com/RayElg/pgmqtt/releases/tag/0.5.1)

pgmqtt 将一个 MQTT 代理嵌入到 PostgreSQL 中。它可以通过逻辑解码发布 INSERT、UPDATE 和 DELETE 变化，并可以将传入的 MQTT 主题和 JSON 载荷映射到表写入操作。当数据库与 MQTT 的集成值得在 PostgreSQL 服务器进程中运行网络代理时，请使用此扩展。

### 预加载并创建扩展

设置逻辑 WAL 并预加载工作进程，然后重启 PostgreSQL：

    wal_level = logical
    shared_preload_libraries = 'pgmqtt'
    pgmqtt.database = 'postgres'

重启后连接到 pgmqtt.database 指定的数据库（默认为 postgres），在其中创建扩展。工作进程使用的表、映射和逻辑复制槽都属于该数据库，映射的业务表也应建在这里：

    CREATE EXTENSION pgmqtt;

监听地址、端口、认证和 TLS 设置由后台工作进程读取。文档中注明仅在启动时需要的设置需要重新启动工作进程或服务器，而不是仅仅使用 pg_reload_conf()。

### 发布表变化

创建一个传出映射：

    SELECT pgmqtt_add_outbound_mapping(
      'public',
      'orders',
      'orders/{{ op | lower }}',
      '{{ columns | tojson }}',
      1
    );

该映射将行变化发布到如 orders/insert 这样的主题。接口还接受 QoS 和模板类型（在支持时）。版本 0.5.1 每次最多处理 4096 条 CDC 变化。

检查或移除传出映射：

    SELECT * FROM pgmqtt_list_outbound_mappings();
    SELECT pgmqtt_remove_outbound_mapping('public', 'orders');

### 从 MQTT 写入行

将捕获的主题片段和 JSON 字段映射到目标表：

    SELECT pgmqtt_add_inbound_mapping(
      'sensor/{site_id}/temperature',
      'sensor_readings',
      '{"site_id":"{site_id}","value":"$.temperature"}'::jsonb,
      mapping_name => 'temp_readings'
    );

传入映射支持插入和文档中描述的 UPSERT 或 DELETE 模式，具有如 target_schema、conflict_columns、mapping_name 和 template_type 等选项。仅授予工作进程所需的表权限，并验证数据包类型和约束。

    SELECT * FROM pgmqtt_list_inbound_mappings();
    SELECT pgmqtt_remove_inbound_mapping('temp_readings');

### 管理与状态

    SELECT * FROM pgmqtt_status();
    SELECT pgmqtt_disconnect_client('device-42');
    SELECT pgmqtt_disconnect_role('mqtt_devices');
    SELECT pgmqtt_reload_acls('*');

pgmqtt_status 报告监听器、客户端、订阅、保留消息、CDC、传入写入和死信队列状态。管理调用由工作进程异步处理并排队。

### 配置索引

- pgmqtt.mqtt_enabled 和 pgmqtt.mqtt_port: TCP MQTT 监听器。
- pgmqtt.ws_enabled 和 pgmqtt.ws_port: WebSocket 监听器。
- pgmqtt.tick_interval_ms 和 pgmqtt.cdc_every_n_ticks: 工作进程的节拍频率。
- pgmqtt.max_client_buffer_bytes: 每个客户端的流控边界。
- pgmqtt.debug_log 和 pgmqtt.metrics_*: 调试和指标集成。
- pgmqtt TLS, JWT, 密码认证和 ACL 设置: 传输和客户端访问控制；社区版与企业版之间有所不同。在设置监听器期望之前，请验证版本。

### 协议与 CDC 边界

- 支持 MQTT 5.0 和 3.1.1 版本。实现了 QoS 0 和 1，请求的 QoS 2 被降级为 QoS 1。
- CDC 包括 INSERT、UPDATE 和 DELETE 操作，不包括 DDL 或 TRUNCATE。DELETE 数据包可能需要 REPLICA IDENTITY FULL。
- CDC 环有有限容量 8192，并在溢出时丢弃最旧的记录。QoS 0 主题缓冲区限制为 4096 并也丢弃最旧条目；每客户端待发队列超过 50,000 条消息时会断开该客户端。
- 社区版文档要求通过代理提供 TLS，而原生 TLS 和一些 JWT 特性是企业版边界。在设置监听器期望之前，请验证版本。

### v0.5.1 及操作

0.5.1 提供经过上游测试的 PostgreSQL 15–18 发布产物。部分服务器构建提供 output_plugin_libraries，会拒绝未列入名单的逻辑解码插件。如果工作进程提示不允许将 pgmqtt 用作输出插件，应按发布说明将它加入该构建的允许列表并重载配置；没有此配置项的服务器不应添加它。

可选的企业版多进程模式还要求启动时设置 pgmqtt.experimental_multiprocess，变更后须完整重启。此模式引入独立 CDC 工作进程及表存储的待发队列，但不意味着其他消息路径的溢出限制全部消失。所有 pgmqtt GUC 都需要超级用户权限。

将代理运行在 PostgreSQL 中扩展了数据库网络和资源边界。隔离监听接口、强制执行认证和主题 ACL、监控工作进程滞后和丢失缓冲区，并在生产使用前测试故障转移和重启行为。
