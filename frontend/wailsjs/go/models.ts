export namespace station {
	
	export class Hill {
	    signal: string;
	    repeaters: string;
	
	    static createFrom(source: any = {}) {
	        return new Hill(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.signal = source["signal"];
	        this.repeaters = source["repeaters"];
	    }
	}
	export class Definition {
	    hills: Hill[];
	
	    static createFrom(source: any = {}) {
	        return new Definition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hills = this.convertValues(source["hills"], Hill);
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

