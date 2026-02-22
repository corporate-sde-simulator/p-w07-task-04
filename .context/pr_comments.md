# PR Review - Blue-green deployment switcher (by Suresh)

## Reviewer: Nisha Gupta
---

**Overall:** Good foundation but critical bugs need fixing before merge.

### `deploymentSwitcher.go`

> **Bug #1:** Switch action swaps traffic before health check passes sending traffic to unhealthy green
> This is the higher priority fix. Check the logic carefully and compare against the design doc.

### `healthGate.go`

> **Bug #2:** Rollback does not restore the previous blue version and just stops green without switching back
> This is more subtle but will cause issues in production. Make sure to add a test case for this.

---

**Suresh**
> Acknowledged. I have documented the issues for whoever picks this up.
