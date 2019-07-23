var mapzen = mapzen || {};
mapzen.whosonfirst = mapzen.whosonfirst || {};

mapzen.whosonfirst.render = (function(){

	var self = {

		'render_data': function(d, ctx){

			var wrapper = document.createElement("div");
			wrapper.setAttribute("class", "dumper table-responsive");

			var render;
			
			if (Array.isArray(d)){
				render = self.render_list(d, ctx);
			}
	    
			else if (typeof(d) == "object"){
				render = self.render_dict(d, ctx);
			}

			else {
				var content = self.render_text(d, ctx);

				render = document.createElement("span");				
				render.appendChild(content);
			}

			wrapper.appendChild(render);
			return wrapper;			
		},

		'render_dict': function(d, ctx){

			var table = document.createElement("table");
			table.setAttribute("class", "table table-details");
	    
			for (k in d){
		
				var row = document.createElement("tr");
				var label_text = k;
				
				var _ctx = (ctx) ? ctx + "." + k : k;
				
				var header = document.createElement("th");
				var label = document.createTextNode(label_text);
				header.appendChild(label);
		
				var content = document.createElement("td");
				
				var body = self.render_data(d[k], _ctx);
				content.appendChild(body);
				
				row.appendChild(header);
				row.appendChild(content);
				
				table.appendChild(row);
			}
			
			return table;
		},
	
		'render_list': function(d, ctx){
	    
			var count = d.length;
			
			if (count == 0){
				return self.render_text("–", ctx);
			}
	    
			if (count <= 1){
				return self.render_data(d[0], ctx);
			}
	    
			var list = document.createElement("ul");
	    
			for (var i=0; i < count; i++){
				
				var item = document.createElement("li");
				var body = self.render_data(d[i], ctx + "#" + i);
				
				item.appendChild(body);
				list.appendChild(item);
			}
			
			return list;
		},
	
		'render_text': function(d, ctx){
			
			var text = d;	// self.htmlspecialchars(d);
			
			var span = document.createElement("span");
			span.setAttribute("id", ctx);
			span.setAttribute("title", ctx);
	    		
			var el = document.createTextNode(text);
			span.appendChild(el);
			
			return span;
		},
		
	};

	return self;

})();	
