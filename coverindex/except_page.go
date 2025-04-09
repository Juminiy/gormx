package coverindex

// ONLY a NOTATION
// after create batch records (by Create(&StructList)), they have same `created_at` time
// when queryList Statement uses order by `created_at`, driver will return unOrdered list
// solution is to add order by <pk_cols> same order with `created_at`
