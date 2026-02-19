export namespace commandUtils {
	
	export class Response__Groundwork_backend_internal_domain_CompanyResponse_ {
	    status: number;
	    message: string;
	    data?: domain.CompanyResponse;
	
	    static createFrom(source: any = {}) {
	        return new Response__Groundwork_backend_internal_domain_CompanyResponse_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], domain.CompanyResponse);
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
	export class Response__Groundwork_backend_internal_domain_SubscriptionResponse_ {
	    status: number;
	    message: string;
	    data?: domain.SubscriptionResponse;
	
	    static createFrom(source: any = {}) {
	        return new Response__Groundwork_backend_internal_domain_SubscriptionResponse_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], domain.SubscriptionResponse);
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
	export class Response____Groundwork_backend_internal_domain_CompanyResponse_ {
	    status: number;
	    message: string;
	    data: domain.CompanyResponse[];
	
	    static createFrom(source: any = {}) {
	        return new Response____Groundwork_backend_internal_domain_CompanyResponse_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], domain.CompanyResponse);
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
	export class Response____Groundwork_backend_internal_domain_SubscriptionResponse_ {
	    status: number;
	    message: string;
	    data: domain.SubscriptionResponse[];
	
	    static createFrom(source: any = {}) {
	        return new Response____Groundwork_backend_internal_domain_SubscriptionResponse_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], domain.SubscriptionResponse);
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
	export class Response_interface____ {
	    status: number;
	    message: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Response_interface____(source);
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
	
	export class CompanyBase {
	    name: string;
	    area: string;
	    tannssid: string;
	    comments: string[];
	    is_active: boolean;
	
	    static createFrom(source: any = {}) {
	        return new CompanyBase(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.area = source["area"];
	        this.tannssid = source["tannssid"];
	        this.comments = source["comments"];
	        this.is_active = source["is_active"];
	    }
	}
	export class CompanyResponse {
	    id: number[];
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new CompanyResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = source["created_at"];
	    }
	}
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
	    start_date?: string;
	    end_date?: string;
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
	        this.start_date = source["start_date"];
	        this.end_date = source["end_date"];
	        this.auto_renew = source["auto_renew"];
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
	export class SubscriptionBase {
	    name: string;
	    license: string;
	    tier: string;
	    features: string[];
	    status: string;
	    start_date: string;
	    end_date?: string;
	    auto_renew: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SubscriptionBase(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.license = source["license"];
	        this.tier = source["tier"];
	        this.features = source["features"];
	        this.status = source["status"];
	        this.start_date = source["start_date"];
	        this.end_date = source["end_date"];
	        this.auto_renew = source["auto_renew"];
	    }
	}
	export class SubscriptionResponse {
	    id: number[];
	    company_id: number[];
	    SubscriptionBase: SubscriptionBase;
	
	    static createFrom(source: any = {}) {
	        return new SubscriptionResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.company_id = source["company_id"];
	        this.SubscriptionBase = this.convertValues(source["SubscriptionBase"], SubscriptionBase);
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
	    start?: string;
	    end_date?: string;
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
	        this.start = source["start"];
	        this.end_date = source["end_date"];
	        this.auto_renew = source["auto_renew"];
	    }
	}

}

