package coverindex

// when Query List ON gorm.CreateInBatches
// Statement uses order by created_at desc
// as well add order by pk desc, for same create_time wrong order case
