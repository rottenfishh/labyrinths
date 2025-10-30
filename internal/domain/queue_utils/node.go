package queue_utils

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type Node interface {
	Distance() int
	Prev() Node
	Coords() domain.Coords
}
