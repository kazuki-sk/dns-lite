package main

type Zone struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:255;uniqueIndex;not null"`
	DefaultTTL int32  `gorm:"not null"`
}

type Record struct {
	ID     uint   `gorm:"primaryKey"`
	ZoneID uint   `gorm:"not null;index"`
	Zone   Zone   `gorm:"constraint:OnDelete:CASCADE"`
	Name   string `gorm:"size:255;not null"`
	Class  uint16 `gorm:"not null"`
	Type   uint16 `gorm:"not null"`
	Value  string `gorm:"not null"`
	TTL    int32  `gorm:"not null"`
}

type SOA struct {
	ID      uint   `gorm:"primaryKey"`
	ZoneID  uint   `gorm:"uniqueIndex;not null"`
	Zone    Zone   `gorm:"constraint:OnDelete:CASCADE"`
	MName   string `gorm:"size:255;not null"`
	RName   string `gorm:"size:255;not null"`
	Serial  int64  `gorm:"not null"`
	Refresh int64  `gorm:"not null"`
	Retry   int64  `gorm:"not null"`
	Expire  int64  `gorm:"not null"`
	Minimum int64  `gorm:"not null"`
	TTL     int32  `gorm:"not null"`
}
