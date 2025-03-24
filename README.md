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
2. ~~⚠️(FeatureReported) Feature report: BeforeDelete Hooks with Struct, Slice, Map~~
3. ~~⚠️(FeatureReported) AfterFind Hooks with Find Map: ~~map[string]any~~, *map[string]any, *[]map[string]any~~
4. ✅(BugFix;2025.03.08) New Feature report: TenantID as a type do data isolation, can be overlapped: tenant_id, user_id, admin_id, operator_id, ...
5. ⚠️New Plugin: codegen for gin+gorm, fiber+gorm short for `CodeGenerator`
6. ✅New Plugin: optlock short for `OptimisticLock`
7. ✅Feature report: new type in tenants.ID, tenants.HideID

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
2. ✅(Tested)Feature report: uniques IN columns expression support omitempty value

> 2025.03.13 
1. ✅(2025.03.14)(FixedBug report): uniques IN columns list, skip all-null-expr groups
2. ✅(2025.03.14)(Feature report): create map not null Constraint, set a default or go-zero value

> 2025.03.14
1. ✅(2025.03.15)Feature report(Plugin Tests Failed): array or slice of struct delete by Pk skip check in clauses.
2. ✅Fix bug report: update or delete primaryKeyClause by model (struct, slice, array).

> 2025.03.15
1. ✅CRUD all basic plugin feature full test passed.
2. ⚠️TODO: plugin more complex case:
   1. uniques: different group combine; some group lack of all fields, some group lack of some fields
   2. ✅ tenants: type string, *string, *uint, tenants.ID, tenants.HideID
   3. dynamicsql: remove not known column but not to conflict with schemas basic.
   4. schemas: not to parse schemas when both model and dest is nil for .Raw(`table`).Scan(&dest)
   5. schemas: to test with schemas when model not conform with dest
   6. user manual
   7. export the callback and clauses function to outside

> 2025.03.17
1. ✅feature support tenants.ID, tenants.HideID, tenants.SID
2. ⚠️full test of uniques plugin with complex uniques

> 2025.03.19
1. ✅optimistic lock by updatedAt is ok.
2. ✅clause where infer by PrimaryKey much stronger.
3. ✅need to test versionTag field.

> 2025.03.20
1. ✅OptimisticLock full feature tested.
2. ✅clause where infer by PrimaryKey test passed.

> 2025.03.22
1. codegen
2. enum value mapping (codegen)
3. PrimaryKey generator when create : can be replaced by BeforeCreate
4. field encrypt when create and update
5. field sensitive
6. ✅TODO: test about types

> 2025.03.24
1. ConfigOption: Explain in query or Row
2. Fix Bug: callback.hasSchemaAndDestIsMap gorm only support three Map type, not derived type.
3. Fix Bug: schemas.RMBCent Type Scan From DB
Notice of best practice: 
- MarshalJSON: realizeOf `json.Marshaler`
  > Get Field Value From Memory, Write to ioBuffer
  - Method receiver should be ValueReceiver
  - json.Marshal must not be receiver itself, will be stackoverflow
    - use typeCast receiver or it's Fields
  - use EncodeByHand for Customization
    - literal value: Number,null,true,false, can encode itself
    - combine value: time or string or special, notice of escape: ""
- UnmarshalJSON: realizeOf `json.Unmarshaler`
  > Parse Field Value From []byte, Assign to Memory
  - Method receiver must be PointerReceiver
  - must not json.Unmarshal From receiver itself, will be stackoverflow
  - should use another value unmarshal from bytes, assign to receiver itself by typeCast or FieldsFill
  - should use receiver Fields to unmarshal from bytes
- Scan: realizeOf `sql.Scanner`
  > Read Field Value From Database, Assign to Memory
  - Method receiver must be PointerReceiver
  - typeAssert or typeCast isOk must return nil error
  - typeAssert or typeCast notOk should return readable error
  - typeAssert type should equal to ValueReturnType
  - must assign src to receiver itself by typeCast or fieldsAssign
  - if not change receiver anything, Scan Field From DB will nothing
- Value: realizeOf `driver.Valuer`
  > Get Field Value From Memory, Write to Database
  - Method receiver should be ValueReceiver
  - must not return receiver itself, will be stackoverflow
  - value must typeCast or typeConvert to driver.Value
  - return type should equal to ScanFromType