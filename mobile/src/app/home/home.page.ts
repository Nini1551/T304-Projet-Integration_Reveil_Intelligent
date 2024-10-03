import { Component } from '@angular/core';
import { alarm } from '../classes/alarms';
import { Alarm } from '../interfaces/alarms';

@Component({
  selector: 'app-home',
  templateUrl: 'home.page.html',
  styleUrls: ['home.page.scss'],
})
export class HomePage {

  alarmData:Alarm={
    id:0,
    name:"",
    hour:0,
    location:"",
    ringtone:"",
    prepTime:0,
    active:false
  }

  constructor() {}

}
