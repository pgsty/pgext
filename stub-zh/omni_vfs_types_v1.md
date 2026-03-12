


## 用法

> [omni_vfs_types_v1: 虚拟文件系统类型（v1）](https://github.com/omnigres/omnigres)

`omni_vfs_types_v1` 扩展定义了 `omni_vfs` 使用的核心类型。

### 类型

**`omni_vfs_types_v1.file`** -- 描述文件条目：
- `name` (text) -- 文件名
- `kind` (file_kind) -- `file` 或 `dir`

**`omni_vfs_types_v1.file_info`** -- 描述文件元数据：
- `size` (bigint) -- 文件大小
- `created_at` (timestamp) -- 创建时间
- `accessed_at` (timestamp) -- 访问时间
- `modified_at` (timestamp) -- 修改时间
- `kind` (file_kind) -- 文件类型

**`omni_vfs_types_v1.file_kind`** -- 枚举，值为 `file` 和 `dir`。
