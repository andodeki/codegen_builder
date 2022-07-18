package main 


// UserID is a function a user can serve
type UserID string

// NilUserID is an empty User
type NilUserID UserID


    //users
    type Users struct { 
    UserId UserID `json:"UserId,omitempty" db:"u_user_id"`
      Email *string `json:"email" db:"u_email"`
    Phone *string `json:"phone" db:"u_phone"`
    PasswordHash *[]byte  `json:"passwordhash" db:"u_password_hash"`
    CreatedAt *time.Time `json:"createdat" db:"u_created_at"`
    UpdatedAt *time.Time `json:"updatedat" db:"u_updated_at"`
    DeletedAt *time.Time `json:"deletedat" db:"u_deleted_at"`
    }
  

  var usersValidationRules = map[int]func(u Users) bool{
      0: func(u Users) bool {
      ///Here1
          return strings.Contains(*u.Email, "@") && len(*u.Email) > 5
      },
      1: func(u Users) bool {
      ///Here2
      return len(*u.Phone) != 0 && u.Phone != nil
      },
      2: func(u Users) bool {
      ///Here5
      return len(*u.PasswordHash) != 0 && u.PasswordHash != nil
      }, 
  }

  func (u *Users) IsMyUsersValid() bool {
    for _, rule := range usersValidationRules {
      if !rule(*u) {
        return false
      }
    }
    return true
  }

