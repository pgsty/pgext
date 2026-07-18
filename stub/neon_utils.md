## Usage

Sources:

- [Official Neon neon_utils documentation](https://neon.com/docs/extensions/neon-utils)

`neon_utils` is a Neon-provided extension for observing Autoscaling. Its `num_cpus()` function reports the current whole number of CPU cores allocated to the connected Neon compute, making it useful while load-testing an autoscaling range.

### Read the Current Allocation

```sql
CREATE EXTENSION IF NOT EXISTS neon_utils;
SELECT num_cpus();
```

Run the query repeatedly from a Neon compute while applying load to observe allocation changes.

### Caveats

- `neon_utils` is provider-managed and documented for Neon; no portable source repository or self-hosted installation is published.
- `num_cpus()` is correct only for computes with Autoscaling enabled. The official documentation says it does not return a correct value for a fixed-size compute.
- Fractional Compute Unit allocations are rounded up to a whole number: allocations of 0.25 or 0.5 CU are reported as 1.
- No preload step is documented. Availability is controlled by the Neon service and may depend on the compute's PostgreSQL version and service configuration.
