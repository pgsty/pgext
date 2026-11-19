## 用法

来源：

- [官方资料](https://github.com/ltslab/orafit/blob/v1.0.3/README.md)
- [官方资料](https://github.com/ltslab/orafit/releases/download/v1.0.3/orafit-jdbc-1.0.3-runtime.tar.gz)
- [官方资料](https://github.com/ltslab/orafit/blob/v1.0.3/LICENSE)
- [官方资料](https://github.com/ltslab/orafit/blob/v1.0.3/COMPATIBILITY.md)
- [官方资料](https://github.com/ltslab/orafit/releases/tag/v1.0.3)

orafit 是 Orafit JDBC 兼容层的服务端 SQL runtime。1.0.3 分发包中的原生控制文件和安装 SQL 仍为 1.0.0；本次修复位于 JDBC 驱动，不要求重新安装未变化的 runtime。PostgreSQL 17 是主要目标，16 受支持，18 仅进行兼容检查。

### 启用

```sql
CREATE EXTENSION orafit;
```

原生 runtime 将不可迁移的对象安装到 orafit 模式，无需预加载。托管服务应使用官方事务式纯 SQL 安装器。应用 classpath 需包含匹配的 Java 17+ Orafit 驱动、其声明的 JSqlParser 依赖，以及单独提供的 pgJDBC；只有 SQL 辅助函数不会翻译应用 SQL。

```properties
driver=io.github.orafit.jdbc.OrafitDriver
url=jdbc:orafit://localhost:5432/app
```

### 兼容边界

Oracle 18c XE 是兼容性参照。COMPATIBILITY.md 列出支持子集与明确拒绝项；解析通过不等于结果等价。驱动默认采用 autosave=always 和 cleanupSavepoints=true 以保留语句错误行为，连接属性可覆盖这些值。1.0.3 修复了不安全的运算符及外连接改写；应重新执行分析器，因为此前被错误翻译的语句现在可能被拒绝。

Ora2Pg 迁移后，仅当源列及源/目标编码确实要求该策略时，才用 orafit.apply_byte_length_semantics('public') 添加 BYTE 长度约束，应在加载或接受写入之前审核。Runtime 不负责结构或数据迁移，也不提供通用 Oracle 仿真。
