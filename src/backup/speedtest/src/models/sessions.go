package main 


// SessionID is a function a user can serve
type SessionID string

// Sessions2ID is a function a user can serve
type Sessions2ID string

// NilSessionID is an empty Session
type NilSessionID SessionID
// NilSessions2ID is an empty Sessions2
type NilSessions2ID Sessions2ID


type BrowserName string 
type BrowserVersion string 
type DeviceId string 
type IpType string 
type OsName string 
type OsVersion string 
const (
Browsername BrowserName = "browsername"
)

const (
Browserversion BrowserVersion = "browserversion"
)

const (
Device DeviceId = "device"
)

const (
Ipv4 IpType = "ipv4"
Ipv6 IpType = "ipv6"
)

const (
Osname OsName = "osname"
)

const (
Osversion OsVersion = "osversion"
)

    //sessions
    type Sessions struct { 
    UserId SessionID `json:"UserId,omitempty" db:"s_user_id"`
      DeviceId DeviceId `json:"deviceid" db:"s_device_id"`
    OsName OsName `json:"osname" db:"s_os_name"`
    OsVersion OsVersion `json:"osversion" db:"s_os_version"`
    BrowserName BrowserName `json:"browsername" db:"s_browser_name"`
    BrowserVersion BrowserVersion `json:"browserversion" db:"s_browser_version"`
    Ip IpType `json:"ip" db:"s_ip"`
    RefreshToken *string `json:"refreshtoken" db:"s_refresh_token"`
    ExpiresAt *time.Time `json:"expiresat" db:"s_expires_at"`
    }
  

    //sessions2
    type Sessions2 struct { 
    UserId Sessions2ID `json:"UserId,omitempty" db:"s_user_id"`
      DeviceId DeviceId `json:"deviceid" db:"s_device_id"`
    OsName OsName `json:"osname" db:"s_os_name"`
    OsVersion OsVersion `json:"osversion" db:"s_os_version"`
    BrowserName BrowserName `json:"browsername" db:"s_browser_name"`
    BrowserVersion BrowserVersion `json:"browserversion" db:"s_browser_version"`
    Ip IpType `json:"ip" db:"s_ip"`
    RefreshToken *string `json:"refreshtoken" db:"s_refresh_token"`
    ExpiresAt *time.Time `json:"expiresat" db:"s_expires_at"`
    }
  

  var sessionsValidationRules = map[int]func(s Sessions) bool{
      0: func(s Sessions) bool {
      ///Here5
      return len(*s.DeviceId) != 0 && s.DeviceId != nil
      },
      1: func(s Sessions) bool {
      ///Here5
      return len(*s.OsName) != 0 && s.OsName != nil
      },
      2: func(s Sessions) bool {
      ///Here5
      return len(*s.OsVersion) != 0 && s.OsVersion != nil
      },
      3: func(s Sessions) bool {
      ///Here5
      return len(*s.BrowserName) != 0 && s.BrowserName != nil
      },
      4: func(s Sessions) bool {
      ///Here5
      return len(*s.BrowserVersion) != 0 && s.BrowserVersion != nil
      },
      5: func(s Sessions) bool {
      ///Here5
      return len(*s.Ip) != 0 && s.Ip != nil
      },
      6: func(s Sessions) bool {
      ///Here2
      return len(*s.RefreshToken) != 0 && s.RefreshToken != nil
      },
      7: func(s Sessions) bool {
      ///Here5
      return len(*s.ExpiresAt) != 0 && s.ExpiresAt != nil
      }, 
  }

  func (s *Sessions) IsMySessionsValid() bool {
    for _, rule := range sessionsValidationRules {
      if !rule(*s) {
        return false
      }
    }
    return true
  }


  var sessions2ValidationRules = map[int]func(s Sessions2) bool{
      0: func(s Sessions2) bool {
      ///Here5
      return len(*s.DeviceId) != 0 && s.DeviceId != nil
      },
      1: func(s Sessions2) bool {
      ///Here5
      return len(*s.OsName) != 0 && s.OsName != nil
      },
      2: func(s Sessions2) bool {
      ///Here5
      return len(*s.OsVersion) != 0 && s.OsVersion != nil
      },
      3: func(s Sessions2) bool {
      ///Here5
      return len(*s.BrowserName) != 0 && s.BrowserName != nil
      },
      4: func(s Sessions2) bool {
      ///Here5
      return len(*s.BrowserVersion) != 0 && s.BrowserVersion != nil
      },
      5: func(s Sessions2) bool {
      ///Here5
      return len(*s.Ip) != 0 && s.Ip != nil
      },
      6: func(s Sessions2) bool {
      ///Here2
      return len(*s.RefreshToken) != 0 && s.RefreshToken != nil
      },
      7: func(s Sessions2) bool {
      ///Here5
      return len(*s.ExpiresAt) != 0 && s.ExpiresAt != nil
      }, 
  }

  func (s *Sessions2) IsMySessions2Valid() bool {
    for _, rule := range sessions2ValidationRules {
      if !rule(*s) {
        return false
      }
    }
    return true
  }

