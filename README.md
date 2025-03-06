# gormx
Reinforced for gorm, aims to be best partner of gorm, rich plugins and toolkits for ORM.

## Migrate From kube
> 2025.03.05
1. Do Full Feature Tests
2. Add TenantID Default Clause
3. Fix Time Type to receive: timestamp, time.Time, sql.NullTime
4. Move update and delete clause check to clause_checker

> 2025.03.06
1. BUG report: uniques and tenants not consistency, because of tx session not clone current settings.
   1. tenants setting eg. tenant_id in callbacks
   2. hooks setting as well, eg. user_id in BeforeUpdate hooks
2. Feature report: BeforeDelete Hooks with Array, Slice, Map
3. Feature report: AfterFind Hooks with Find Map: map[string]any, *map[string]any, *[]map[string]any
4. New Feature report: TenantID as a type do data isolation, can be overlapped: tenant_id, user_id, admin_id, operator_id, ...
5. New Plugin: codegen for gin+gorm, fiber+gorm short for `CodeGenerator`
6. New Plugin: optlock short for `OptimisticLock`