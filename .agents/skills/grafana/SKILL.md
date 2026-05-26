```markdown
# grafana Development Patterns

> Auto-generated skill from repository analysis

## Overview

This skill teaches you how to contribute to the [grafana](https://github.com/grafana/grafana) codebase, a large-scale observability platform written primarily in Go and React. You'll learn the repository's coding conventions, commit patterns, and common development workflows, including dependency updates, feature toggling, CI script changes, database migrations, and API spec updates. This guide also covers testing practices and provides ready-to-use commands for frequent tasks.

## Coding Conventions

### File Naming

- **CamelCase** is used for file names.
  - Example: `featureToggleRegistry.go`, `alertingRules.tsx`

### Import Style

- **Relative imports** are preferred.
  - Example (Go):
    ```go
    import "../models"
    ```
  - Example (TypeScript/React):
    ```tsx
    import { PanelOptions } from '../types'
    ```

### Export Style

- **Named exports** are used.
  - Example (TypeScript):
    ```ts
    export const PanelOptions = { ... };
    export function getDashboard() { ... }
    ```

### Commit Patterns

- Commit messages often use prefixes (e.g., `release`, `ci`, `i18n`, `chore`, `alerting`, `docs`, `dashboards`, `fix`).
- Messages are concise (average ~57 characters).
  - Example: `fix: resolve dashboard loading error`

## Workflows

### Update NPM Dependency

**Trigger:** When a dependency needs to be updated to a new version  
**Command:** `/update-dependency`

1. Update the version in the root `package.json`.
2. Update the version in relevant `packages/*/package.json` and/or `public/app/plugins/datasource/*/package.json`.
3. Regenerate `yarn.lock` (e.g., run `yarn install`).
4. Commit all changes.

**Files Involved:**
- `package.json`
- `packages/*/package.json`
- `public/app/plugins/datasource/*/package.json`
- `yarn.lock`

---

### Add or Update Feature Toggle

**Trigger:** When a new feature flag is added or an existing one is updated  
**Command:** `/add-feature-toggle`

1. Edit or add an entry in `pkg/services/featuremgmt/toggles_gen.csv`.
2. Update `pkg/services/featuremgmt/registry.go`.
3. Regenerate `pkg/services/featuremgmt/toggles_gen.go` and `toggles_gen.json`.
4. Update `packages/grafana-data/src/types/featureToggles.gen.ts`.
5. Update documentation at `docs/sources/setup-grafana/configure-grafana/feature-toggles/index.md` if needed.

**Files Involved:**
- `pkg/services/featuremgmt/toggles_gen.csv`
- `pkg/services/featuremgmt/registry.go`
- `pkg/services/featuremgmt/toggles_gen.go`
- `pkg/services/featuremgmt/toggles_gen.json`
- `packages/grafana-data/src/types/featureToggles.gen.ts`
- `docs/sources/setup-grafana/configure-grafana/feature-toggles/index.md`

---

### Update Changelog

**Trigger:** When a new release is prepared or significant changes are made  
**Command:** `/update-changelog`

1. Edit `CHANGELOG.md` with new entries.
2. Commit the updated file.

**Files Involved:**
- `CHANGELOG.md`

---

### CI Build Script Update

**Trigger:** When the build process needs adjustment (e.g., new flags, missing steps, new environments)  
**Command:** `/update-ci-script`

1. Edit one or more scripts under `pkg/build/daggerbuild/scripts/`.
2. Optionally update `.drone.yml` or `scripts/drone/*.star` if CI config is affected.
3. Commit all changes.

**Files Involved:**
- `pkg/build/daggerbuild/scripts/*.sh`
- `.drone.yml`
- `scripts/drone/*.star`

---

### Add or Update Database Table or Index

**Trigger:** When a new table/index is needed or a schema change is required  
**Command:** `/new-table`

1. Add or update a migration file (e.g., `*_mig.go` or `*.sql`).
2. Update or add related model/store code (`*.go`).
3. Add or update tests if needed.
4. Commit all changes.

**Files Involved:**
- `pkg/services/sqlstore/migrations/*.go`
- `pkg/storage/secret/metadata/data/*.sql`
- `pkg/storage/secret/metadata/keeper_model.go`
- `pkg/storage/secret/metadata/keeper_store.go`
- `pkg/storage/secret/metadata/keeper_store_test.go`

---

### Update API or OpenAPI Spec

**Trigger:** When API endpoints are added/changed or documentation needs to be updated  
**Command:** `/update-api-docs`

1. Edit docs in `docs/sources/developers/http_api/*.md`.
2. Update `public/api-*.json` or `public/openapi*.json`.
3. Commit all changes.

**Files Involved:**
- `docs/sources/developers/http_api/*.md`
- `public/api-enterprise-spec.json`
- `public/api-merged.json`
- `public/openapi3.json`

---

## Testing Patterns

- **Framework:** [Jest](https://jestjs.io/)
- **Test File Pattern:** `*.test.tsx`
- **Location:** Alongside source files or in dedicated test directories.

**Example:**
```tsx
// panelOptions.test.tsx
import { render } from '@testing-library/react';
import { PanelOptions } from './PanelOptions';

test('renders panel options', () => {
  const { getByText } = render(<PanelOptions />);
  expect(getByText('Panel Options')).toBeInTheDocument();
});
```

## Commands

| Command              | Purpose                                                      |
|----------------------|--------------------------------------------------------------|
| /update-dependency   | Update a JS/TS dependency across package.json and yarn.lock   |
| /add-feature-toggle  | Add or update a feature toggle and regenerate related files   |
| /update-changelog    | Add entries to the changelog for a new release or changes    |
| /update-ci-script    | Update CI build or publish scripts                           |
| /new-table           | Add or modify a database table or index with migrations      |
| /update-api-docs     | Update API documentation and OpenAPI spec files              |
```
