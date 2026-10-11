## 用法

来源：

- [README.md](https://github.com/John-Cameron-Burns/codesystem/blob/6c1a7bcea5e2e808d47c04bb68d96eac0a101cbe/README.md)
- [codesystem--0.1.0.sql](https://github.com/John-Cameron-Burns/codesystem/blob/6c1a7bcea5e2e808d47c04bb68d96eac0a101cbe/codesystem--0.1.0.sql)
- [sql/define.sql](https://github.com/John-Cameron-Burns/codesystem/blob/6c1a7bcea5e2e808d47c04bb68d96eac0a101cbe/sql/define.sql)
- [codesystem.control](https://github.com/John-Cameron-Burns/codesystem/blob/6c1a7bcea5e2e808d47c04bb68d96eac0a101cbe/codesystem.control)

`codesystem` 存储具名分类体系与具有类型约束的编码值。0.1.0 原型已实现定义、校验、层级导航与修订追踪；设计中的派生标准与映射关系尚未实现。

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
```

### 运行边界

上游在 PostgreSQL 17 上测试此修订。原生扩展需要管理员安装。`define_code_system` 的输入查询必须提供编码、层级、父编码和标签四列，定义过程会拒绝不合法的层级、重复编码和断裂的层级关系。字面量必须带体系前缀，外部文本可用 `code_from` 转换。`code_level`、`code_parent`、`code_ancestor`、`code_children` 和 `code_label` 用于层级查询。不同修订中相同的文本表示不同值，注册表的结构性修改受到保护。恢复自定义格式备份时需要上游的顺序辅助脚本，先恢复注册表，再恢复带编码类型的表数据；普通恢复顺序可能失败。
