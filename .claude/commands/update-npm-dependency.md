---
name: update-npm-dependency
description: Workflow command scaffold for update-npm-dependency in grafana.
allowed_tools: ["Bash", "Read", "Write", "Grep", "Glob"]
---

# /update-npm-dependency

Use this workflow when working on **update-npm-dependency** in `grafana`.

## Goal

Automated or manual update of a JavaScript/TypeScript dependency across multiple package.json files and yarn.lock.

## Common Files

- `package.json`
- `packages/*/package.json`
- `public/app/plugins/datasource/*/package.json`
- `yarn.lock`

## Suggested Sequence

1. Understand the current state and failure mode before editing.
2. Make the smallest coherent change that satisfies the workflow goal.
3. Run the most relevant verification for touched files.
4. Summarize what changed and what still needs review.

## Typical Commit Signals

- Update version in root package.json
- Update version in relevant packages/*/package.json and/or public/app/plugins/datasource/*/package.json
- Regenerate yarn.lock
- Commit all changes

## Notes

- Treat this as a scaffold, not a hard-coded script.
- Update the command if the workflow evolves materially.