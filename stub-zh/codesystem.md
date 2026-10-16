## 用法

来源：

- [README.md](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/README.md)
- [docs/USER_GUIDE.md](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/docs/USER_GUIDE.md)
- [codesystem--0.2.0.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/codesystem--0.2.0.sql)
- [codesystem--0.1.0--0.2.0.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/codesystem--0.1.0--0.2.0.sql)
- [codesystem.control](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/codesystem.control)
- [.github/workflows/test.yml](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/.github/workflows/test.yml)
- [tools/gss_update.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/tools/gss_update.sql)
- [tools/install_isic_isco.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/tools/install_isic_isco.sql)
- [NOTICE](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/NOTICE)

`codesystem` 存储具名分类体系与具有类型约束的 `code` 值。0.2.0 支持列表层级、修订、派生标准、映射关系以及按模式定义的标识符。模式只能验证语法，不能判断标识符是否真的已签发。

### 核心用法

```sql
CREATE EXTENSION codesystem;
CREATE TABLE classification_source (code text, level text, parent text, label text);
INSERT INTO classification_source VALUES ('A', 'section', '', 'Agriculture');
SELECT define_code_system(
  'toy',
  '[{"name":"section","width":1,"charset":"[A-Z]","link":"none","expected":1}]',
  'SELECT * FROM classification_source', 'Toy classification', 'Example');
SELECT name, current_revision, ncodes FROM code_systems;
SELECT code_label('toy:A'::code);
SELECT define_code_system_pattern('my_id', 'AA-NNN');
SELECT code_from('AB-123', 'my_id');
```

### 运行边界

上游 CI 以 PostgreSQL 13–17 为目标。此原生扩展需要管理员安装；只有 SQL 扩展版本变化时才执行 `ALTER EXTENSION codesystem UPDATE`；分发包 0.2.1 仍使用 SQL/control 版本 0.2.0，已有 0.2.0 数据库无需升级。

`define_code_system` 校验层级，以及输入查询提供的编码、层级、父编码和标签四列。SQL 字面量必须带体系前缀，外部文本可用 `code_from` 转换。`code_level`、`code_parent`、`code_ancestor`、`code_children` 和 `code_label` 用于层级查询。不同修订中相同的文本表示不同值；编码仍有效时，可用 `code_current` 转换到当前修订。

`define_code_system` 还可通过 `refines` 和 `shared_levels` 定义派生分类，再用 `code_rollup` 向基础体系归并。`define_crosswalk` 注册对应关系，`code_map` 可能返回多个匹配，`code_map_unique` 则要求恰好匹配一个。`define_code_system_pattern` 定义有界模板或正则表达式，`code_part` 和 `code_parts` 返回命名片段。`code_from_padded` 可处理丢失的前导零，`code_first_with_prefix` 与 `code_last_with_prefix` 提供可利用索引的前缀范围。一个体系不能在列表形式与模式形式之间切换。

注册表的结构性修改受到保护，允许的编码集合变化时应定义新修订。写注册表需要相应表权限，读取编码需要注册表的 SELECT 权限。恢复自定义格式备份时，应使用上游 `tools/restore.sh` 辅助脚本，先恢复注册表，再恢复带编码类型的表数据，普通恢复顺序可能失败。

### 可选分类数据工具

分发包 0.2.1 提供 GSS 编码模板及实体类型数据。另行加载 `tools/gss_update.sql` 后，`gss_update_entity_types()` 可从 ONS 更新实体类型；加载 `tools/install_isic_isco.sql` 后，`codesystem_install_isic_rev4()` 与 `codesystem_install_isco_08()` 分别从 UN 和 ILO 下载分类。这些辅助工具需要预先创建 `plpython3u`、由超级用户安装，并允许数据库服务器访问对应来源；核心扩展本身不依赖它们。

先用 `dry_run` 检查下载内容；写入分类注册表时应连接主库。ISIC/ISCO 数据不随包重新分发，其使用仍遵循发布机构条款。上游 `NOTICE` 记录其他随包数据的署名和许可。
