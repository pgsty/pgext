## 用法

来源：

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [extensions/pg_sheet/pg_sheet.control](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_sheet/pg_sheet.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/README.md)
- [extensions/pg_sheet/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_sheet/Cargo.toml)
- [extensions/pg_sheet/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_sheet/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/LICENSE)
- [extensions/pg_sheet/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_sheet/README.md)

`pg_sheet` 0.3.1 通过 `pgsheet` 为实体表增加可编辑覆盖层，保留源数据，并生成合并源值、覆盖值与公式的视图。

### 核心用法

```sql
CREATE EXTENSION pg_sheet;
CREATE TABLE sheet_entities (id uuid PRIMARY KEY, title text);
INSERT INTO sheet_entities VALUES ('11111111-1111-1111-1111-111111111111', 'Example');
SELECT pgsheet.create_sheet('review', 'public', 'sheet_entities');
SELECT pgsheet.add_column('review', 'notes', 'text');
SELECT pgsheet.set_value('review', '11111111-1111-1111-1111-111111111111', 'notes', 'Checked');
SELECT pgsheet.get_data('review', 100, 0);
```

### 运行边界

控制文件不限定超级用户安装，但仍需具备创建相应对象的权限。无需预加载。核心流程使用 `create_sheet`、`add_column`、`set_value` 与 `get_data`。列公式会转换为 SQL，单元格公式则保存后交给客户端 HyperFormula 计算。快照／差异、审计历史和单元格锁用于协作。过滤器、公式与类型定义会生成 SQL，应由可信人员维护。示例使用 UUID 实体标识。恢复覆盖层不会还原底层源表。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

### 当前版本与升级

0.3.1 修复可选角色 `app_user` 不存在时 `create_sheet` 失败的问题。仅在该角色存在时授权，应用权限仍需显式配置。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.4.1；上游二进制不代表 Pigsty 软件包可用性。
