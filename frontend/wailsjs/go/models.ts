export namespace commandUtils {
	
	export class Response {
	    status: number;
	    message: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}

}

export namespace domain {
	
	export class CreateCompanyRequest {
	    name?: string;
	    area?: string;
	    tannssid?: string;
	    is_active?: boolean;
	    comments?: string[];
	
	    static createFrom(source: any = {}) {
	        return new CreateCompanyRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.area = source["area"];
	        this.tannssid = source["tannssid"];
	        this.is_active = source["is_active"];
	        this.comments = source["comments"];
	    }
	}
	export class CreateSubscriptionRequest {
	    company_id: number[];
	    name: string;
	    license: string;
	    tier?: string;
	    features?: string[];
	    status?: string;
	    // Go type: time
	    start_date?: any;
	    // Go type: time
	    end_date?: any;
	    auto_renew?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new CreateSubscriptionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.company_id = source["company_id"];
	        this.name = source["name"];
	        this.license = source["license"];
	        this.tier = source["tier"];
	        this.features = source["features"];
	        this.status = source["status"];
	        this.start_date = this.convertValues(source["start_date"], null);
	        this.end_date = this.convertValues(source["end_date"], null);
	        this.auto_renew = source["auto_renew"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class DeleteCompanyRequest {
	    id: number[];
	
	    static createFrom(source: any = {}) {
	        return new DeleteCompanyRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}
	export class DeleteSubscriptionRequest {
	    id: number[];
	
	    static createFrom(source: any = {}) {
	        return new DeleteSubscriptionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}
	export class GetCompanyRequest {
	    id: number[];
	
	    static createFrom(source: any = {}) {
	        return new GetCompanyRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}
	export class GetSubscriptionRequest {
	    id: number[];
	
	    static createFrom(source: any = {}) {
	        return new GetSubscriptionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}
	export class ListCompanySubscriptionsRequest {
	    company_id: number[];
	
	    static createFrom(source: any = {}) {
	        return new ListCompanySubscriptionsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.company_id = source["company_id"];
	    }
	}
	export class UpdateCompanyRequest {
	    id: number[];
	    name?: string;
	    tannssid?: string;
	    area?: string;
	    is_active?: boolean;
	    comments?: string[];
	
	    static createFrom(source: any = {}) {
	        return new UpdateCompanyRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.tannssid = source["tannssid"];
	        this.area = source["area"];
	        this.is_active = source["is_active"];
	        this.comments = source["comments"];
	    }
	}
	export class UpdateSubscriptionRequest {
	    id: number[];
	    name?: string;
	    license?: string;
	    tier?: string;
	    features?: string[];
	    status?: string;
	    // Go type: time
	    start?: any;
	    // Go type: time
	    end_date?: any;
	    auto_renew?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new UpdateSubscriptionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.license = source["license"];
	        this.tier = source["tier"];
	        this.features = source["features"];
	        this.status = source["status"];
	        this.start = this.convertValues(source["start"], null);
	        this.end_date = this.convertValues(source["end_date"], null);
	        this.auto_renew = source["auto_renew"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

