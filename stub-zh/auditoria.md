## 用法

来源：

- [官方扩展控制文件（auditoria.control）](https://github.com/zsnails/postgres-audit-log/blob/44368b5696450699618b11c26bfcc2ff40af9bda/auditoria.control)
- [官方扩展 SQL 脚本（auditoria--0.0.1.sql）](https://github.com/zsnails/postgres-audit-log/blob/44368b5696450699618b11c26bfcc2ff40af9bda/auditoria--0.0.1.sql)

`auditoria` — PostgreSQL 审计扩展。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION auditoria;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小 SQL 脚本，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `auditoria.registrar_actividad(tabla_afectada TEXT, tipo_operacion TEXT, datos_anteriores JSONB, datos_nuevos JSONB)` 是一个扩展函数，返回 `VOID`。
- `auditoria.registrar_en_tablas_nuevas()` 是一个扩展函数，返回 `event_trigger`。
- `auditoria.trig_registrar_actividad()` 是一个扩展函数，返回 `TRIGGER`。
- `auditoria.registrar_auditores` 是一个扩展过程。
- `auditoria.registrar_en_todo_lado` 是一个扩展过程。
- `auditoria.registro_auditoria` 是一个由扩展安装或管理的表。
- `auditoria` 是一个由扩展创建的模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
