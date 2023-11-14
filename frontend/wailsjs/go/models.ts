export namespace main {
	
	export class Spell {
	    name: string;
	    des: string;
	
	    static createFrom(source: any = {}) {
	        return new Spell(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.des = source["des"];
	    }
	}

}

