---
name: add-or-update-feature-toggle
description: Workflow command scaffold for add-or-update-feature-toggle in grafana.
allowed_tools: ["Bash", "Read", "Write", "Grep", "Glob"]
---

# /add-or-update-feature-toggle

Use this workflow when working on **add-or-update-feature-toggle** in `grafana`.

## Goal

Introduce or update a feature toggle, updating registry, generated types, and documentation.

## Common Files

- `pkg/services/featuremgmt/toggles_gen.csv`
- `pkg/services/featuremgmt/registry.go`
- `pkg/services/featuremgmt/toggles_gen.go`
- `pkg/services/featuremgmt/toggles_gen.json`
- `packages/grafana-data/src/types/featureToggles.gen.ts`
- `docs/sources/setup-grafana/configure-grafana/feature-toggles/index.md`

## Suggested Sequence

1. Understand the current state and failure mode before editing.
2. Make the smallest coherent change that satisfies the workflow goal.
3. Run the most relevant verification for touched files.
4. Summarize what changed and what still needs review.

## Typical Commit Signals

- Edit or add entry in pkg/services/featuremgmt/toggles_gen.csv
- Update pkg/services/featuremgmt/registry.go
- Regenerate pkg/services/featuremgmt/toggles_gen.go and toggles_gen.json
- Update packages/grafana-data/src/types/featureToggles.gen.ts
- Update docs/sources/setup-grafana/configure-grafana/feature-toggles/index.md if needed

## Notes

- Treat this as a scaffold, not a hard-coded script.
- Update the command if the workflow evolves materially.