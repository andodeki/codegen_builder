package main 


// UnitID is a function a user can serve
type UnitID string

// NilUnitID is an empty Unit
type NilUnitID UnitID


type OwnerType string 
const (
Rental OwnerType = "rental"
Buying OwnerType = "buying"
Leasing OwnerType = "leasing"
)

    //units
    type Units struct { 
    UnitId UnitID `json:"UnitId,omitempty" db:"u_unit_id"`
      PropertyId *string `json:"propertyid" db:"u_property_id"`
    TenantId *string `json:"tenantid" db:"u_tenant_id"`
    SpaceUnitsUnit *string `json:"spaceunitsunit" db:"u_space_units_unit"`
    SpaceUnitsType *string `json:"spaceunitstype" db:"u_space_units_type"`
    SpaceUnitsCapacity *string `json:"spaceunitscapacity" db:"u_space_units_capacity"`
    SpaceUnitsNo *string `json:"spaceunitsno" db:"u_space_units_no"`
    SpaceUnitsFlrLevel *string `json:"spaceunitsflrlevel" db:"u_space_units_flr_level"`
    SpaceUnitsSqArea *string `json:"spaceunitssqarea" db:"u_space_units_sq_area"`
    SpaceUnitsFlrPlans *string `json:"spaceunitsflrplans" db:"u_space_units_flr_plans"`
    Furnished *bool `json:"furnished" db:"u_furnished"`
    Refurbishing *bool `json:"refurbishing" db:"u_refurbishing"`
    BronchureUploads *string `json:"bronchureuploads" db:"u_bronchure_uploads"`
    OwnershipType OwnerType `json:"ownershiptype" db:"u_ownership_type"`
    OwnershipDocs *string `json:"ownershipdocs" db:"u_ownership_docs"`
    OccupancyDocs *string `json:"occupancydocs" db:"u_occupancy_docs"`
    Currency *string `json:"currency" db:"u_currency"`
    CreatedAt *time.Time `json:"createdat" db:"u_created_at"`
    UpdatedAt *time.Time `json:"updatedat" db:"u_updated_at"`
    DeletedAt *time.Time `json:"deletedat" db:"u_deleted_at"`
    }
  

  var unitsValidationRules = map[int]func(u Units) bool{
      0: func(u Units) bool {
      ///Here2
      return len(*u.PropertyId) != 0 && u.PropertyId != nil
      },
      1: func(u Units) bool {
      ///Here2
      return len(*u.TenantId) != 0 && u.TenantId != nil
      },
      2: func(u Units) bool {
      ///Here2
      return len(*u.SpaceUnitsUnit) != 0 && u.SpaceUnitsUnit != nil
      },
      3: func(u Units) bool {
      ///Here2
      return len(*u.SpaceUnitsType) != 0 && u.SpaceUnitsType != nil
      },
      4: func(u Units) bool {
      ///Here2
      return len(*u.SpaceUnitsCapacity) != 0 && u.SpaceUnitsCapacity != nil
      },
      5: func(u Units) bool {
      ///Here2
      return len(*u.SpaceUnitsNo) != 0 && u.SpaceUnitsNo != nil
      },
      6: func(u Units) bool {
      ///Here2
      return len(*u.SpaceUnitsFlrLevel) != 0 && u.SpaceUnitsFlrLevel != nil
      },
      7: func(u Units) bool {
      ///Here2
      return len(*u.SpaceUnitsSqArea) != 0 && u.SpaceUnitsSqArea != nil
      },
      8: func(u Units) bool {
      ///Here2
      return len(*u.SpaceUnitsFlrPlans) != 0 && u.SpaceUnitsFlrPlans != nil
      },
      9: func(u Units) bool {
      ///Here3
      return *u.Furnished && u.Furnished != nil
      },
      10: func(u Units) bool {
      ///Here3
      return *u.Refurbishing && u.Refurbishing != nil
      },
      11: func(u Units) bool {
      ///Here2
      return len(*u.BronchureUploads) != 0 && u.BronchureUploads != nil
      },
      12: func(u Units) bool {
      ///Here5
      return len(*u.OwnershipType) != 0 && u.OwnershipType != nil
      },
      13: func(u Units) bool {
      ///Here2
      return len(*u.OwnershipDocs) != 0 && u.OwnershipDocs != nil
      },
      14: func(u Units) bool {
      ///Here2
      return len(*u.OccupancyDocs) != 0 && u.OccupancyDocs != nil
      },
      15: func(u Units) bool {
      ///Here2
      return len(*u.Currency) != 0 && u.Currency != nil
      }, 
  }

  func (u *Units) IsMyUnitsValid() bool {
    for _, rule := range unitsValidationRules {
      if !rule(*u) {
        return false
      }
    }
    return true
  }

