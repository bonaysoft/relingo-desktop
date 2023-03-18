export namespace main {
	
	export class Word {
	    id: number;
	    name: string;
	    exposures: number;
	    mastered: boolean;
	    raw_json: string;
	    raw_object: relingo.DictItem;
	    root: model.Result;
	    // Go type: time.Time
	    created_at: any;
	    // Go type: time.Time
	    updated_at: any;
	    // Go type: time.Time
	    deleted_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Word(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.exposures = source["exposures"];
	        this.mastered = source["mastered"];
	        this.raw_json = source["raw_json"];
	        this.raw_object = this.convertValues(source["raw_object"], relingo.DictItem);
	        this.root = this.convertValues(source["root"], model.Result);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.deleted_at = this.convertValues(source["deleted_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ListResult {
	    items: Word[];
	    total: number;
	
	    static createFrom(source: any = {}) {
	        return new ListResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], Word);
	        this.total = source["total"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace model {
	
	export class Vocabulary {
	    name: string;
	    phonetic: string;
	    mnemonic: string;
	    constitute: string[];
	    meaning: string;
	    tags: string[];
	    children: Vocabulary[];
	
	    static createFrom(source: any = {}) {
	        return new Vocabulary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.phonetic = source["phonetic"];
	        this.mnemonic = source["mnemonic"];
	        this.constitute = source["constitute"];
	        this.meaning = source["meaning"];
	        this.tags = source["tags"];
	        this.children = this.convertValues(source["children"], Vocabulary);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Result {
	    self?: Vocabulary;
	    root?: Vocabulary;
	
	    static createFrom(source: any = {}) {
	        return new Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.self = this.convertValues(source["self"], Vocabulary);
	        this.root = this.convertValues(source["root"], Vocabulary);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace relingo {
	
	export class Translation {
	    target: string;
	    pos: string;
	    score: number;
	
	    static createFrom(source: any = {}) {
	        return new Translation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.target = source["target"];
	        this.pos = source["pos"];
	        this.score = source["score"];
	    }
	}
	export class DictItem {
	    phonetic: string[];
	    variant: string[];
	    wordFrequency: number;
	    definition: string;
	    _id: string;
	    source: string;
	    lang: string;
	    translations: Translation[];
	    display: string;
	    mastered: boolean;
	    stared: boolean;
	    revision: boolean;
	    needRevise: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DictItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.phonetic = source["phonetic"];
	        this.variant = source["variant"];
	        this.wordFrequency = source["wordFrequency"];
	        this.definition = source["definition"];
	        this._id = source["_id"];
	        this.source = source["source"];
	        this.lang = source["lang"];
	        this.translations = this.convertValues(source["translations"], Translation);
	        this.display = source["display"];
	        this.mastered = source["mastered"];
	        this.stared = source["stared"];
	        this.revision = source["revision"];
	        this.needRevise = source["needRevise"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RespUserInfo {
	    name: string;
	    avatar: string;
	    email: string;
	    token: string;
	
	    static createFrom(source: any = {}) {
	        return new RespUserInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.avatar = source["avatar"];
	        this.email = source["email"];
	        this.token = source["token"];
	    }
	}
	
	export class VocabularyListItem {
	    count?: number;
	    mastered?: number;
	    _id?: string;
	    uid?: string;
	    name: string;
	    id: string;
	    type: string;
	    // Go type: time.Time
	    createdAt?: any;
	    updatedAt?: number;
	    privilege?: string;
	    scope?: string;
	    level?: string;
	    words?: string[];
	
	    static createFrom(source: any = {}) {
	        return new VocabularyListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.count = source["count"];
	        this.mastered = source["mastered"];
	        this._id = source["_id"];
	        this.uid = source["uid"];
	        this.name = source["name"];
	        this.id = source["id"];
	        this.type = source["type"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = source["updatedAt"];
	        this.privilege = source["privilege"];
	        this.scope = source["scope"];
	        this.level = source["level"];
	        this.words = source["words"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

