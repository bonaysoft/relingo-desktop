// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {relingo} from '../models';

export function GetUserInfo():Promise<relingo.RespUserInfo>;

export function GetVocabulary(arg1:string,arg2:string):Promise<Array<string>>;

export function GetVocabularyList():Promise<Array<relingo.VocabularyListItem>>;

export function MasteredWords(arg1:string):Promise<Array<string>>;

export function Ready():Promise<boolean>;

export function SetToken(arg1:string):Promise<void>;

export function SubmitVocabulary(arg1:Array<string>):Promise<void>;
