package db

import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  Username  string `json:"username" gorm:"size:30"`
  Password  string `json:"password" gorm:"size:40"`
  Usertype  string `json:"usertype" gorm:"size:20"`
  Firstname string `json:"firstname" gorm:"size:20"`
  Lastname  string `json:"lastname" gorm:"size:20"`
  Address   string `json:"address" gorm:"size:30"`
  Phone     string `json:"phone" gorm:"size:20"`
  Email     string `json:"email" gorm:"size:40"`
}

type Vehicle struct {
  gorm.Model
  Type  string `json:"type" gorm:"size:20"`
  Brand  string `json:"brand" gorm:"size:20"`
  Number  string `json:"number" gorm:"size:20"`
  OwnerEmail string `json:"owneremail" gorm:"size:30"`
  Owner  User  `json:"owner" gorm:"foreignKey:UserID;references:id"`
  Locations []Location `json:"locations"`
  UserID uint `json:"userid"`
}

type Location struct {
  gorm.Model
  Location  string `json:"location" gorm:"size:30"`
  Latitude  string `json:"latitude" gorm:"size:20"`
  Longitude  string `json:"longitude" gorm:"size:20"`
  VehicleID uint `json:"vehicleid"`
}

type Device struct {
  gorm.Model
  UniqueLabel string  `json:"uniquelabel"`
  Ownername string `json:"ownername"`
  Ownerphone string `json:"ownerphone"`
  Owneremail  string  `json:"owneremail"`
  Type string `json:"type" gorm:"default:car"`
  State string `json:"state" gorm:"default:stationary"`
  CurrentHolder string `json:"currentHolder" gorm:"default:owner"`
  Longitude float32 `json:"longitude"`
  Latitude float32 `json:"latitude"`
  ObsvrLongitude float32 `json:"obsvrlongitude"`
  ObsvrLatitude float32 `json:"obsvrlatitude"`
  Missing bool `json:"missing" gorm:"default:false"`
}

type DeviceHistory struct {
  gorm.Model
  UniqueLabel string  `json:"uniquelabel"`
  Type string `json:"type" gorm:"default:car"`
  State string `json:"state" gorm:"default:stationary"`
  CurrentHolder string `json:"currentHolder" gorm:"default:owner"`
  Longitude float32 `json:"longitude"`
  Latitude float32 `json:"latitude"`
  Missing bool `json:"missing" gorm:"default:false"`
  Notice string `json:"notice"`
}



