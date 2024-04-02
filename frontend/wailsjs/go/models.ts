export namespace models {
	
	export class Doc {
	    parentId: string;
	    docType: string;
	    title: string;
	    fileId: string;
	    children: Doc[];
	    createdBy: string;
	    updatedBy: string;
	    id: string;
	    createdAt: string;
	    updatedAt: string;
	
	    static createFrom(source: any = {}) {
	        return new Doc(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.parentId = source["parentId"];
	        this.docType = source["docType"];
	        this.title = source["title"];
	        this.fileId = source["fileId"];
	        this.children = this.convertValues(source["children"], Doc);
	        this.createdBy = source["createdBy"];
	        this.updatedBy = source["updatedBy"];
	        this.id = source["id"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
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

export namespace services {
	
	export class ServiceResponse {
	    code: number;
	    message: string;
	    data: any;
	    success: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ServiceResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = source["data"];
	        this.success = source["success"];
	    }
	}

}

