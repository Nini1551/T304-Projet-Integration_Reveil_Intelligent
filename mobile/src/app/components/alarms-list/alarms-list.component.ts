import { Component, OnInit } from '@angular/core';
import { alarm } from '../../classes/alarms';
import { Alarm } from '../../interfaces/alarms';
import { ALARMS_LIST_MOCK } from 'src/app/mocks/alarms-list.mock';

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

  alarms: Alarm[] = ALARMS_LIST_MOCK;

  constructor() {}

  ngOnInit() {}

}