// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {station} from '../models';

export function CheckForUpdate(arg1:string):Promise<boolean>;

export function Connect(arg1:string):Promise<void>;

export function Disconnect():Promise<void>;

export function GetStationHash():Promise<string>;

export function LoadStationFile():Promise<station.Definition>;

export function SetSignal(arg1:station.Hill,arg2:string):Promise<void>;
