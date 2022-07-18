package main 


// PaymentID is a function a user can serve
type PaymentID string

// NilPaymentID is an empty Payment
type NilPaymentID PaymentID


type PaymentsChannels string 
type PaymentsType string 
type TransactionMode string 
type TransactionStatus string 
type TransactionType string 
const (
Mpesapayments PaymentsChannels = "mpesapayments"
TKashpayments PaymentsChannels = "tKashpayments"
AirtelMoneypayments PaymentsChannels = "airtelMoneypayments"
Debitcreditprepaidpayments PaymentsChannels = "debitcreditprepaidpayments"
Eazzypaypayments PaymentsChannels = "eazzypaypayments"
Eagentpayments PaymentsChannels = "eagentpayments"
Kcbcashpayments PaymentsChannels = "kcbcashpayments"
Equitypayments PaymentsChannels = "equitypayments"
Pesalinkpayments PaymentsChannels = "pesalinkpayments"
Paypalpayments PaymentsChannels = "paypalpayments"
Cashpayments PaymentsChannels = "cashpayments"
)

const (
Rentalfee PaymentsType = "rentalfee"
Buyingfee PaymentsType = "buyingfee"
Leasingfee PaymentsType = "leasingfee"
Amenitiesfee PaymentsType = "amenitiesfee"
Powerconnfee PaymentsType = "powerconnfee"
Waterconnfee PaymentsType = "waterconnfee"
Powertokenfee PaymentsType = "powertokenfee"
Watertokenfee PaymentsType = "watertokenfee"
Viewingfee PaymentsType = "viewingfee"
)

const (
Offline TransactionMode = "offline"
Online TransactionMode = "online"
Wired TransactionMode = "wired"
Draft TransactionMode = "draft"
Cheque TransactionMode = "cheque"
Cashondelivery TransactionMode = "cashondelivery"
)

const (
New TransactionStatus = "new"
Cancelled TransactionStatus = "cancelled"
Failed TransactionStatus = "failed"
Pending TransactionStatus = "pending"
Declined TransactionStatus = "declined"
Rejected TransactionStatus = "rejected"
Success TransactionStatus = "success"
)

const (
Credit TransactionType = "credit"
Debit TransactionType = "debit"
)

    //payments
    type Payments struct { 
    PaymentId PaymentID `json:"PaymentId,omitempty" db:"p_payment_id"`
      ConnAcctId *string `json:"connacctid" db:"p_conn_acct_id"`
    TokenEntryId *string `json:"tokenentryid" db:"p_token_entry_id"`
    UserId *string `json:"userid" db:"p_user_id"`
    BusinessId *string `json:"businessid" db:"p_business_id"`
    ScheduleId *string `json:"scheduleid" db:"p_schedule_id"`
    Code *string `json:"code" db:"p_code"`
    Type TransactionType `json:"type" db:"p_type"`
    Amount *int64 `json:"amount" db:"p_amount"`
    Status TransactionStatus `json:"status" db:"p_status"`
    PaymentType PaymentsType `json:"paymenttype" db:"p_payment_type"`
    Mode TransactionMode `json:"mode" db:"p_mode"`
    PaymentChannel PaymentsChannels `json:"paymentchannel" db:"p_payment_channel"`
    CreatedAt *time.Time `json:"createdat" db:"p_created_at"`
    UpdatedAt *time.Time `json:"updatedat" db:"p_updated_at"`
    DeletedAt *time.Time `json:"deletedat" db:"p_deleted_at"`
    }
  

  var paymentsValidationRules = map[int]func(p Payments) bool{
      0: func(p Payments) bool {
      ///Here2
      return len(*p.ConnAcctId) != 0 && p.ConnAcctId != nil
      },
      1: func(p Payments) bool {
      ///Here2
      return len(*p.TokenEntryId) != 0 && p.TokenEntryId != nil
      },
      2: func(p Payments) bool {
      ///Here2
      return len(*p.UserId) != 0 && p.UserId != nil
      },
      3: func(p Payments) bool {
      ///Here2
      return len(*p.BusinessId) != 0 && p.BusinessId != nil
      },
      4: func(p Payments) bool {
      ///Here2
      return len(*p.ScheduleId) != 0 && p.ScheduleId != nil
      },
      5: func(p Payments) bool {
      ///Here2
      return len(*p.Code) != 0 && p.Code != nil
      },
      6: func(p Payments) bool {
      ///Here5
      return len(*p.Type) != 0 && p.Type != nil
      },
      7: func(p Payments) bool {
      ///Here4
      return p.Amount != nil
      },
      8: func(p Payments) bool {
      ///Here5
      return len(*p.Status) != 0 && p.Status != nil
      },
      9: func(p Payments) bool {
      ///Here5
      return len(*p.PaymentType) != 0 && p.PaymentType != nil
      },
      10: func(p Payments) bool {
      ///Here5
      return len(*p.Mode) != 0 && p.Mode != nil
      },
      11: func(p Payments) bool {
      ///Here5
      return len(*p.PaymentChannel) != 0 && p.PaymentChannel != nil
      }, 
  }

  func (p *Payments) IsMyPaymentsValid() bool {
    for _, rule := range paymentsValidationRules {
      if !rule(*p) {
        return false
      }
    }
    return true
  }

