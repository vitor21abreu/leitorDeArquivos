package domain

import "time"

type Vehicle struct {
    ID        string
    Model     string
    Status    string
    SoldAt    time.Time
}