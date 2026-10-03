## 用法

来源：

- [extensions/pg_mqtt/pg_mqtt.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_mqtt/pg_mqtt.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_mqtt/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_mqtt/Cargo.toml)
- [extensions/pg_mqtt/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_mqtt/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_mqtt/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_mqtt/README.md)
- [extensions/pg_mqtt/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_mqtt/pgbrew.toml)
- [extensions/pg_mqtt/src/protocol/codec.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_mqtt/src/protocol/codec.rs)
- [extensions/pg_mqtt/src/protocol/handlers.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_mqtt/src/protocol/handlers.rs)

`pg_mqtt` 0.3.0 使用 `pgmqtt` 表保存消息并实现 MQTT 代理。所核对解码器只接受协议级别 5，拒绝其他版本；扩展 README 中旧的 3.1.1 声明已经过时。

### 核心用法

```conf
shared_preload_libraries = 'pg_mqtt'
pg_mqtt.database = 'postgres'
pg_mqtt.port = 1883
```

```sql
CREATE EXTENSION pg_mqtt;
SELECT pgmqtt.mqtt_publish('sensors/room1/temperature', '22.5');
SELECT * FROM pgmqtt.mqtt_messages('sensors/%', 10);
```

### 运行边界

控制文件要求超级用户安装。需要预加载并重启，在配置的工作数据库中创建扩展。`pg_mqtt.database`、`pg_mqtt.port` 和 `pg_mqtt.worker_count` 配置工作进程。源码默认在所有接口的 1883 端口监听，应隔离这一实验性服务。`mqtt_publish`、`mqtt_messages`、`mqtt_sessions` 及模式绑定函数提供数据接口；`mqtt_messages` 使用 SQL LIKE 模式；`mqtt_status` 只报告配置，不是监听端口健康检查。不能仅根据项目描述推断完整的 MQTT 安全或交付保证。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
