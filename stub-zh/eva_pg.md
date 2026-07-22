## 用法

来源：

- [扩展 control 文件](https://github.com/eva-ics/eva4/blob/main/eva_pg/eva_pg/eva_pg.control)
- [pgrx 实现](https://github.com/eva-ics/eva4/blob/main/eva_pg/eva_pg/src/lib.rs)
- [扩展包元数据](https://github.com/eva-ics/eva4/blob/main/eva_pg/eva_pg/Cargo.toml)

`eva_pg` 0.0.1 是 EVA ICS 单体仓库中的 PostgreSQL 子项目。它只暴露一个 pgrx 函数 `oid_match`，用于按 EVA 掩码语法匹配 EVA 对象标识符；它不会把平台的其他服务 API 引入 PostgreSQL。

### 核心流程

由超级用户安装扩展，然后用 `kind:path/segments` 形式的 EVA 标识符与掩码进行比较：

```sql
CREATE EXTENSION eva_pg;

SELECT oid_match('sensor:plant/line1/temp', 'sensor:plant/+/temp');
SELECT oid_match('unit:plant/line1/pump', '+:plant/#');
```

完整掩码 `*` 或 `#` 会立即匹配。其他情况下，标识符和掩码都必须包含分隔类别与路径的冒号，否则函数会报错。

### 匹配规则

- 掩码类别 `+` 或 `?` 匹配任意一个类别；其他类别必须精确相等。
- 路径段 `+` 或 `?` 匹配一个路径段。
- 路径段 `*` 或 `#` 会立即接受剩余路径，因此其后的掩码段会被忽略。
- 不使用剩余路径通配符时，标识符与掩码必须拥有相同数量的路径段。

### 运维说明

control 文件要求超级用户安装，且扩展不可重定位。当前 Rust 包声明了 PostgreSQL 13 到 18 的 pgrx 特性，默认构建特性是 PostgreSQL 16；应针对目标大版本准确构建。该子项目没有独立的 SQL 回归套件，因此在把它用于授权或过滤逻辑前，应测试非法标识符、空路径段、通配符位置以及应用自己的掩码。
