## 用法

来源：

- [官方 README](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/README.md)
- [用法文档](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/doc/variant.md)
- [扩展 SQL](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/sql/variant.sql)
- [C 实现](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/src/variant.c)
- [扩展控制文件](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/variant.control)

`variant` 1.0.1 定义 `variant.variant` 容器，以 PostgreSQL 原生表示保存一个值，同时记录原始类型与类型修饰符。它可以用于异构设置或属性值，但存储行会耦合到底层类型 OID、二进制表示、I/O 函数与动态生成的类型转换。可移植的应用数据应优先使用显式类型列或 JSON。

### 注册并使用 Variant

上游不建议使用默认禁用的 variant，因为部分 PostgreSQL 路径不会保留类型修饰符。应注册一个具名修饰符并用于列：

```sql
CREATE EXTENSION variant;

SELECT variant.register('setting');

CREATE TABLE app_setting (
    setting_name text PRIMARY KEY,
    setting_value variant.variant(setting) NOT NULL
);

INSERT INTO app_setting VALUES
    ('retry_count', 3::integer),
    ('label', 'primary'::text),
    ('window', '((0,0),(1,1))'::box);

SELECT setting_name,
       variant.original_type(setting_value),
       variant.text_out(setting_value)
FROM app_setting;
```

普通类型转换是主要输入/输出接口。`variant.text_in(text, text)` 为指定的具名注册修饰符解析显式文本表示。`variant.original_type(variant.variant)` 标识已存基础类型，`variant.text_out(variant.variant)` 返回其文本形式。

### 类型转换、比较与索引

安装会调用 `variant.create_casts()`，为已有的非复合、非域、非伪类型生成函数与类型转换。安装新的受支持数据类型后应再次运行，并检查 `variant.missing_casts` 与 `variant.variant_casts`。动态创建全目录转换属于高权限 DDL；只能允许扩展管理员执行该函数。

扩展为 variant 提供比较运算符，还提供基于独立镜像相等运算符 `*=` 的哈希运算符类。对于计划使用的每种类型，都应测试语义相等、二进制镜像相等、不同存储类型间排序、排序规则、NaN 值与哈希索引行为，不能假定会保留底层类型的普通运算符类。

### 类型生命周期与可移植性

Variant 保存原生二进制数据，并非自描述可移植编码。上游警告，删除被 variant 行使用的类型会让这些值无法恢复；扩展不会为每个存储行创建依赖，也不能可靠地把注册修饰符限制到已批准类型。删除或升级任何扩展、域、枚举或用户定义类型前，必须盘点已存基础类型。

二进制格式与类型 OID 可能随 PostgreSQL 大版本、扩展升级、架构或恢复顺序改变。采用前应证明能逻辑转储/恢复到空集群，并能转回每一种原始类型。不要复制原始文件，也不要逻辑复制到类型身份不同的目录。复核的项目是陈旧的预览期原生代码，没有当前兼容矩阵；应固定准确制品、限制值大小、测试 TOAST 与畸形文本输入，并保留迁移到显式列或稳定序列化格式的路径。
