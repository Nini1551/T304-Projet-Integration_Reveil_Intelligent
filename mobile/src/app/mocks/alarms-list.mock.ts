// Valeurs moquées pour des alarmes à afficher au démarrage de l'application

import { Alarm } from "../interfaces/alarms";

export const ALARMS_LIST_MOCK: Alarm[] = [
  { id: 1, name: "Projet d'intégration", hour: 8, location: "EPHEC", ringtone: "Ringtone 1", prepTime: 10, active: true },
  { id: 2, name: "Sécurité des réseaux (Théorie)", hour: 12, location: "EPHEC", ringtone: "Ringtone 2", prepTime: 5, active: false },
  { id: 3, name: "Sécurité des réseaux (Pratique)", hour: 14, location: "EPHEC", ringtone: "Ringtone 3", prepTime: 15, active: true }
];