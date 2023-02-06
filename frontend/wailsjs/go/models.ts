export namespace model {
	
	export class Word {
	    id: number;
	    source: string;
	    exposures: number;
	    mastered: boolean;
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
	        this.source = source["source"];
	        this.exposures = source["exposures"];
	        this.mastered = source["mastered"];
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

}

export namespace relingo {
	
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

