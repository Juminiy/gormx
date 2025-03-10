# gormx
Reinforced for gorm, aims to be best partner of gorm, rich plugins and toolkits for ORM.

## Migrate From kube
> 2025.03.05
1. ✅(2025.03.09) Do Full Feature Tests
2. ✅(2025.03.09) Add TenantID Default Clause
3. ✅(2025.03.09) Fix Time Type to receive: timestamp, time.Time, sql.NullTime
4. ✅(2025.03.09) Move update and delete clause check to clause_checker

> 2025.03.06
1. BUG report: uniques and tenants not consistency, because of tx session not clone current settings.
   1. ✅(BugFix;2025.03.08) tenants setting eg. tenant_id in callbacks
   2. ✅(BugFix;2025.03.08) hooks setting as well, eg. user_id in BeforeUpdate hooks
2. ⚠️(FeatureMerged) Feature report: BeforeDelete Hooks with Struct, Slice, Map
3. Feature report: AfterFind Hooks with Find Map: map[string]any, *map[string]any, *[]map[string]any
4. ✅(BugFix;2025.03.08) New Feature report: TenantID as a type do data isolation, can be overlapped: tenant_id, user_id, admin_id, operator_id, ...
5. New Plugin: codegen for gin+gorm, fiber+gorm short for `CodeGenerator`
6. New Plugin: optlock short for `OptimisticLock`
7. Feature report: new type in tenants.ID, tenants.HideID

> 2025.03.08
1. ✅(2025.03.09)Feature report: complex dup check before count do self check for group mode, >=2 rows needed, one row do not need.
2. ✅(2025.03.09)BugFix: build primaryKey clause with multiple Pk clause.
3. Feature report: MySQL optimized columns IN for uniques plugin(field dup check)

> 2025.03.09
1. ✅(2025.03.10)Feature report: need to merge clauses plugins to gormx scopes, as well dependency conflicting resolve.

> 2025.03.10
1. ⚠️(FeatureMerged) Feature report: returning clause in delete and update