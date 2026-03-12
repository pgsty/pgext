


## Usage

> [basic_archive: an example of an archive module](https://www.postgresql.org/docs/current/basic-archive.html)

The `basic_archive` module is a WAL archive module that copies completed WAL segment files to a specified directory. It serves as a reference implementation for custom archive modules.

### Configuration

Add to `postgresql.conf`:

```ini
archive_mode = 'on'
archive_library = 'basic_archive'
basic_archive.archive_directory = '/path/to/archive/directory'
```

### Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `basic_archive.archive_directory` | string | Directory to copy WAL files to (must already exist) |

If `archive_mode` is enabled but `basic_archive.archive_directory` is empty (default), the server will accumulate WAL files until a directory path is configured.

### Notes

- The target directory must be created before use; the module will not create it
- After a server crash, temporary files with the `archtemp` prefix may be left in the archive directory and should be deleted before restarting
- These temporary files can also be safely removed while the server is running, provided they are not related to an ongoing archive operation
- This module is primarily intended as a simple example and starting point for developing custom archive modules
