## 用法

来源：

- [extensions/pg_uom/pg_uom.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/pg_uom.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_uom/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/Cargo.toml)
- [extensions/pg_uom/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_uom/src/seed.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/src/seed.rs)
- [extensions/pg_uom/src/convert.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/src/convert.rs)
- [extensions/pg_uom/src/unit.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_uom/src/unit.rs)

`pg_uom` 0.3.0 在 `pguom` 中保存单位类别与定义，支持带量纲检查的标量及复合单位转换。

### 核心用法

```sql
CREATE EXTENSION pg_uom;
SELECT pguom.seed_si();
SELECT pguom.convert(1.0, 'kg', 'lb');
SELECT pguom.convert(0.0, 'degC', 'degF');
```

### 运行边界

控制文件不限定超级用户安装，但仍需具备创建相应对象的权限。无需预加载。`seed_si` 初始化单位，`create_category`、`create_unit`、`list_units` 和 `compatible_units` 管理目录。简单温度转换支持偏移量，复合温度单位会被拒绝；不同量纲不能互转。`pg_uom.enabled` 可关闭转换并返回原始输入，应用应控制该设置，不能把调用成功等同于已完成数值变换。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
