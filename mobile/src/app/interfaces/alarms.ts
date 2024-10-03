export interface Alarm{
    id: number;
    name: string;
    hour: number;
    location: string;
    ringtone: string;
    prepTime: number;
    active: boolean;
}