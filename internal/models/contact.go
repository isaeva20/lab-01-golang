package models

import (
    "time"
)

type Contact struct {
    ID          int                    `json:"id"`
    Username    string                 `json:"username"`
    GivenName   string                 `json:"given_name"`
    FamilyName  string                 `json:"family_name"`
    Phones      []Phones               `json:"phone"`
    Email       []string               `json:"email"`
    Birthdate   time.Time             `json:"birthdate"`
}

type Phones struct {
    TypeID      int    `json:"type_id"`
    CountryCode int    `json:"country_code"`
    Operator    int    `json:"operator"`
    Number      int    `json:"number"`
}
