import { Alarm } from "../interfaces/alarms";

export class alarm implements Alarm{
    id:number;
    name:string;
    hour:number;
    location:string;
    ringtone:string;
    prepTime:number;
    active:boolean;

    constructor(alarmData:Alarm){
        this.id=alarmData.id;
        this.name=alarmData.name;
        this.hour=alarmData.hour;
        this.location=alarmData.location;
        this.ringtone=alarmData.ringtone;
        this.prepTime=alarmData.prepTime;
        this.active=alarmData.active;
    }
}