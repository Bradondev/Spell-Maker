export namespace main {
	
	export class SpellStats {
	    damage: number;
	    Debuff: string;
	    mana_cost: number;
	    casttime: number;
	    range: string;
	
	    static createFrom(source: any = {}) {
	        return new SpellStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.damage = source["damage"];
	        this.Debuff = source["Debuff"];
	        this.mana_cost = source["mana_cost"];
	        this.casttime = source["casttime"];
	        this.range = source["range"];
	    }
	}
	export class Spell {
	    name: string;
	    description: string;
	    nameOfType: string;
	    levelOfSpell: string;
	    stats: SpellStats;
	
	    static createFrom(source: any = {}) {
	        return new Spell(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.nameOfType = source["nameOfType"];
	        this.levelOfSpell = source["levelOfSpell"];
	        this.stats = this.convertValues(source["stats"], SpellStats);
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

