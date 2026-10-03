## 用法

来源：

- [extensions/pg_fsm/pg_fsm.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/pg_fsm.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_fsm/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/Cargo.toml)
- [extensions/pg_fsm/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_fsm/src/machine.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/src/machine.rs)
- [extensions/pg_fsm/src/binding.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/src/binding.rs)
- [extensions/pg_fsm/src/transition.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_fsm/src/transition.rs)

`pg_fsm` 0.3.0 在 `pgfsm` 中定义有限状态机，支持转移条件、动作、表绑定与历史记录。

### 核心用法

```sql
CREATE EXTENSION pg_fsm;
SELECT pgfsm.create_machine('order_flow', 'draft', 'Order workflow');
SELECT pgfsm.add_state('order_flow', 'submitted');
SELECT pgfsm.add_transition('order_flow', 'draft', 'submitted', 'submit');
SELECT * FROM pgfsm.list_machines();
```

### 运行边界

控制文件不限定超级用户安装，但仍需具备创建相应对象的权限。无需预加载。先定义状态与转移，再绑定应用状态列或调用转移 API。条件和动作执行数据库表达式，应限制其配置权限。`pg_fsm.enabled` 为假时关闭约束，`pg_fsm.disable_default_notify` 关闭默认通知。0.3.0 增加了使用 `if_not_exists` 的幂等初始化。此项目不同于规范名称为 `fsm` 的扩展。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
