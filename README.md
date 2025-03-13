# gormx
Reinforced for gorm, aims to be best partner of gorm, rich plugins and toolkits for ORM.

## Migrate From kube
> 2025.03.05
1. ✅(2025.03.09) Do Full Feature Tests
2. ✅(2025.03.09) Add TenantID Default Clause
3. ✅(2025.03.09) Fix Time Type to receive: timestamp, time.Time
4. ✅(2025.03.09) Move update and delete clause check to clause_checker

> 2025.03.06
1. ✅(BugFixed)BUG report: uniques and tenants not consistency, because of tx session not clone current settings.
   1. ✅(BugFix;2025.03.08) tenants setting eg. tenant_id in callbacks
   2. ✅(BugFix;2025.03.08) hooks setting as well, eg. user_id in BeforeUpdate hooks
2. ⚠️(FeatureReported) Feature report: BeforeDelete Hooks with Struct, Slice, Map
3. ⚠️(FeatureReported) AfterFind Hooks with Find Map: map[string]any, *map[string]any, *[]map[string]any
4. ✅(BugFix;2025.03.08) New Feature report: TenantID as a type do data isolation, can be overlapped: tenant_id, user_id, admin_id, operator_id, ...
5. New Plugin: codegen for gin+gorm, fiber+gorm short for `CodeGenerator`
6. New Plugin: optlock short for `OptimisticLock`
7. Feature report: new type in tenants.ID, tenants.HideID

> 2025.03.08
1. ✅(2025.03.09)Feature report: complex dup check before count do self check for group mode, >=2 rows needed, one row do not need.
2. ✅(2025.03.09)BugFix: build primaryKey clause with multiple Pk clause.
3. ✅(2025.03.11)Feature report: MySQL optimized columns IN for uniques plugin(field dup check)

> 2025.03.09
1. ✅(2025.03.10)Feature report: need to merge clauses plugins to gormx scopes, as well dependency conflicting resolve.

> 2025.03.10
1. ✅(2025.03.11)Feature report: returning clause in delete and update

> 2025.03.11
1. ✅(2025.03.11)Feature report: MySQL8, SQLite3, Postgres17 all supports IN columns expression

> 2025.03.12
1. Full Feature Test in pkg: tests/tests_v2
2. ⚠️(NotTest)Feature report: uniques IN columns expression support omitempty value

> 2025.03.13 
1. ✅(2025.03.11)(FixedBug report): uniques IN columns list, skip all-null-expr groups
2. ✅(2025.03.11)(Feature report): create map not null Constraint, set a default or go-zero value