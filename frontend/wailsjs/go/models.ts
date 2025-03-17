export namespace excel {
	
	export class PersonData {
	    name: string;
	    age: number;
	    height: number;
	    birthday: string;
	
	    static createFrom(source: any = {}) {
	        return new PersonData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.age = source["age"];
	        this.height = source["height"];
	        this.birthday = source["birthday"];
	    }
	}

}

