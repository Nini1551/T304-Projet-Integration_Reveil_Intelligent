import { Component, OnInit } from '@angular/core';
import { Alarm } from '../../interfaces/alarms';
import { AlarmService } from 'src/app/services/alarm.service';

@Component({
  selector: 'app-alarms-list',
  templateUrl: './alarms-list.component.html',
  styleUrls: ['./alarms-list.component.scss'],
})
export class AlarmsListComponent implements OnInit {
  alarms: Alarm[] = [];

  constructor(private alarmService: AlarmService) {

  }
  
  setAlarms(){
    this.alarmService.getAlarms().subscribe((data: any) => {
      console.log(data);
      for (let alarmData of data) {
        let alarm: Alarm = {
          id: alarmData.ID,
          name: alarmData.Name,
          ringDate: alarmData.RingDate,
          createdAt: alarmData.CreatedAt,
          location: alarmData.Location,
          ringtone: alarmData.Ringtone,
          active: alarmData.IsActive
        }
        this.alarms.push(alarm);
      }
      console.log(this.alarms);
    });
  }

  ngOnInit() {
    this.setAlarms();
  }

}