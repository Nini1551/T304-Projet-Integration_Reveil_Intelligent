import { Component, OnInit } from '@angular/core';
import { alarm } from '../../classes/alarms';
import { Alarm } from '../../interfaces/alarms';

@Component({
  selector: 'app-alarms-list',
  templateUrl: './alarms-list.component.html',
  styleUrls: ['./alarms-list.component.scss'],
})
export class AlarmsListComponent  implements OnInit {

  alarmData: Alarm = {
    id: 0,
    name: "",
    hour: 0,
    location: "",
    ringtone: "",
    prepTime: 0,
    active: false
  };

  mockAlarms: Alarm[] = [
    { id: 1, name: "Projet d'intégration", hour: 8, location: "EPHEC", ringtone: "Ringtone 1", prepTime: 10, active: true },
    { id: 2, name: "Sécurité des réseaux (Théorie)", hour: 12, location: "EPHEC", ringtone: "Ringtone 2", prepTime: 5, active: false },
    { id: 3, name: "Sécurité des réseaux (Pratique)", hour: 14, location: "EPHEC", ringtone: "Ringtone 3", prepTime: 15, active: true }
  ];

  constructor() {}

  ngOnInit() {}

}