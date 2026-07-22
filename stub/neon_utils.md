## Usage

Sources:

- [Official Neon documentation](https://neon.com/docs/extensions/neon-utils)

`neon_utils` is a Neon-managed extension that reports the CPU capacity currently allocated to an autoscaling Neon compute. It is meaningful only inside Neon and is not a portable community package or a general operating-system monitoring API.

### Install and Read Capacity

Create the extension in a Neon database, then call its single documented function:

```sql
CREATE EXTENSION IF NOT EXISTS neon_utils;

SELECT num_cpus();
```

`num_cpus()` returns the current number of allocated CPU cores. Poll it alongside workload metrics to observe how a compute moves between its configured autoscaling minimum and maximum.

### Interpretation Limits

Neon measures compute size in Compute Units, which can be fractional. `num_cpus()` does not report fractional sizes: for example, Neon documents that a compute at 0.25 or 0.5 CU returns `1`. The value is therefore an operational signal, not an exact billing, memory, or entitlement measurement.

The function only works correctly on computes with Autoscaling enabled. On a fixed-size compute it does not return a correct value. Confirm the compute configuration in the Neon console before consuming the result.

### Monitoring Notes

Sample at a reasonable interval and record timestamps, workload latency, and the configured autoscaling bounds; a single value cannot explain why a scale event occurred. Avoid using the rounded result as a hard scheduler or admission-control input. Availability, versioning, privileges, and behavior follow the Neon service, so recheck the official page after compute or platform upgrades. The public official page documents the function and limitations but does not publish an extension version; verify the installed `extversion` on the target Neon database when version evidence is required.
