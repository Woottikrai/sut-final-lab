package entity

import (


"gorm.io/gorm"

)
type Employee struct {
	gorm.Model
	Name string 		`valid:"required~name cant ba blank"`// ต้องไม่เป็นค่าว่าง 
	Email string		`valid:"email~email incorrect"`
	EmployeeID string 	`valid:"matches(^[JMS]//d{8}฿)"`// รหัสพนักงานขึ7นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว
   }



