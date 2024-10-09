package mocks

import (
	"server/db/models"
	"time"
)

// ALARM_LIST contient les valeurs mockées pour les alarmes
var ALARM_LIST = []models.Alarm{
	{Name: "Projet d'intégration", RingDate: time.Now().Add(24 * time.Hour), Location: "EPHEC", Ringtone: "Ringtone 1", IsActive: true},
	{Name: "Sécurité des réseaux (Théorie)", RingDate: time.Now().Add(48 * time.Hour), Location: "EPHEC", Ringtone: "Ringtone 2", IsActive: false},
	{Name: "Sécurité des réseaux (Pratique)", RingDate: time.Now().Add(72 * time.Hour), Location: "EPHEC", Ringtone: "Ringtone 3", IsActive: true},
}
