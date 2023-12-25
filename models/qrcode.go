package models

import "github.com/google/uuid"

type QrCode struct {
	QRCode string
	ID     uuid.UUID
}
