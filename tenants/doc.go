package tenants

/*type Product struct {
	gorm.Model
	Name      string `mt:"unique"`
	Code      uint   `mt:"unique"`
	Price     int64
	TenantID  uint `mt:"tenant"`
	UserID    uint `mt:"user"`
	ProjectID uint `mt:"project"`
}*/

// Or

/*type Product struct {
	gorm.Model
	Name      string `mt:"unique"`
	Code      uint   `mt:"unique"`
	Price     int64
	TenantID  tenants.ID `mt:"tenant"`
	UserID    tenants.ID `mt:"user"`
	ProjectID tenants.ID `mt:"project"`
}*/

// Or

/*type Product struct {
	gorm.Model
	Name      string `mt:"unique"`
	Code      uint   `mt:"unique"`
	Price     int64
	TenantID  tenants.HideID `mt:"tenant"`
	UserID    tenants.HideID `mt:"user"`
	ProjectID tenants.HideID `mt:"project"`
}
*/

// Or

/*type MyTenantIDType tenants.ID

func (t MyTenantIDType) MarshalJSON() ([]byte, error) {
	// do your marshal value to json logic
	// ...

	return ...,...
}

func (t MyTenantIDType) UnmarshalJSON(b []byte) error {
	// do your unmarshal value from bytes logic
	// ...

	return ...
}*/
