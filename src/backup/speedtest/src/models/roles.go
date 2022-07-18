package main 

// Role is a function a user can serve
type Role string
// NilRoleID is an empty Role
type NilRoleID Role



type UserRole string 
const (
Admin UserRole = "admin"
)

    //roles
    type UserRoles struct { 
    Role Role `json:"role" db:"r_role"`
    }
  

  var rolesValidationRules = map[int]func(r UserRoles) bool{
      0: func(r UserRoles) bool {
      return len(r.Role) != 0
      }, 
  }

  func (r *UserRoles) IsMyRolesValid() bool {
    for _, rule := range rolesValidationRules {
      if !rule(*r) {
        return false
      }
    }
    return true
  }

