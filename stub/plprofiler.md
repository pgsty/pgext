


## Usage

> [plprofiler: server-side support for profiling PL/pgSQL functions](https://github.com/bigsql/plprofiler)

`plprofiler` is a profiling tool for PL/pgSQL functions that identifies performance bottlenecks and generates interactive flame graph reports.

```sql
CREATE EXTENSION plprofiler;
```

### Configuration

Add to `postgresql.conf`:

```
shared_preload_libraries = 'plprofiler'
```

### Command-Line Tool

The `plprofiler` command-line utility controls profiling and generates reports:

```bash
# Profile a specific SQL command
plprofiler run -d mydb --command "SELECT my_function()" --output report.html

# Monitor profiling in real-time
plprofiler monitor -d mydb

# Save profiling data for later analysis
plprofiler save -d mydb --name "my_profile"

# Generate HTML report with flame graphs
plprofiler report -d mydb --from-data "my_profile" --output report.html

# List saved profiling datasets
plprofiler list -d mydb

# Reset profiling data
plprofiler reset-data -d mydb

# Export/import profiling data
plprofiler export -d mydb --from-data "my_profile" > profile.json
plprofiler import -d mydb --into-data "imported" < profile.json
```

### SQL Interface

```sql
-- Enable profiling for the current session
SELECT pl_profiler_set_enabled_local(true);

-- Execute functions to be profiled
SELECT my_function();

-- Collect profiling data into shared hash tables
SELECT pl_profiler_collect_data();

-- Disable profiling
SELECT pl_profiler_set_enabled_local(false);

-- Enable profiling globally (for all sessions)
SELECT pl_profiler_set_enabled_global(true);

-- Reset local/shared profiling data
SELECT pl_profiler_reset_local();
SELECT pl_profiler_reset_shared();
```

### Report Output

Generated HTML reports include:

- **Interactive flame graphs** showing wall-clock time spent in PL/pgSQL code
- **Per-function statistics** with self-time (total minus children)
- **Top functions** ranked by time consumption (default: top 10)
- Self-contained HTML requiring no external dependencies

### Profiling Methods

- **Direct profiling**: Run specific SQL while collecting data
- **Timed collection**: Interval-based statistics gathering
- **Per-user profiling**: Enable profiling for specific database users via `ALTER USER`
- **Production monitoring**: Low-overhead profiling on live systems
