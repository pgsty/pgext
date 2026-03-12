


## Usage

> [omni_vfs_types_v1: Virtual File System types (v1)](https://github.com/omnigres/omnigres)

The `omni_vfs_types_v1` extension defines the core types used by `omni_vfs`.

### Types

**`omni_vfs_types_v1.file`** -- Describes a file entry:
- `name` (text) -- File name
- `kind` (file_kind) -- Either `file` or `dir`

**`omni_vfs_types_v1.file_info`** -- Describes file metadata:
- `size` (bigint) -- File size
- `created_at` (timestamp) -- Creation time
- `accessed_at` (timestamp) -- Access time
- `modified_at` (timestamp) -- Modification time
- `kind` (file_kind) -- File kind

**`omni_vfs_types_v1.file_kind`** -- Enum with values `file` and `dir`.
