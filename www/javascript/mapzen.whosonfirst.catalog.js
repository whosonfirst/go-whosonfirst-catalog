var mapzen = mapzen || {};
mapzen.whosonfirst = mapzen.whosonfirst || {};

mapzen.whosonfirst.catalog = (function(){

	var self = {

		'lookup': function(id, cb){

		    	var proto = document.location.protocol;
			var url = proto + '//' + location.host + location.pathname + 'id/' + id;

			var onload = function(rsp){
			
				var target = rsp.target;
				
				if (target.readyState != 4){
					return;
				}
				
				var status_code = target['status'];
				var status_text = target['statusText'];
				
				var raw = target['responseText'];
				var data = null;
				
				try {
					data = JSON.parse(raw);
				}
				
				catch (e){
					
					console.log("ERR", url, e)
					return false;
				}

				cb(data);
			};

			var req = new XMLHttpRequest();
			req.addEventListener("load", onload);
		    
			req.open("GET", url, true);
			req.send();
		},

		'render': function(rsp){
		    
		    var records = rsp["recordset"]["records"];
		    var count = records.length;
		    
		    var table = document.createElement("table");

		    var source_header = document.createElement("th");
		    source_header.appendChild(document.createTextNode("source"));
		    
		    var type_header = document.createElement("th");
		    type_header.appendChild(document.createTextNode("type"));
		    
		    var hash_header = document.createElement("th");
		    hash_header.appendChild(document.createTextNode("hash"));
		    
		    var uri_header = document.createElement("th");
		    uri_header.appendChild(document.createTextNode("uri"));
		    
		    var show_header = document.createElement("th");
		    show_header.appendChild(document.createTextNode("-"));
		    
		    var header_row = document.createElement("tr");
		    header_row.appendChild(source_header);
		    header_row.appendChild(type_header);
		    header_row.appendChild(hash_header);
		    header_row.appendChild(uri_header);
		    header_row.appendChild(show_header);								
		    
		    table.appendChild(header_row);

			for (var i=0; i < count; i++){

				var data = records[i];

				var source = data["source"];				
				var type = data["type"];				
				var uri = data["uri"];
				var hash = data["hash"];			

				var source_cell = document.createElement("td");
				source_cell.appendChild(document.createTextNode(source));

				var type_cell = document.createElement("td");
				type_cell.appendChild(document.createTextNode(type));

				var hash_cell = document.createElement("td");
				hash_cell.appendChild(document.createTextNode(hash));

				var uri_cell = document.createElement("td");
				uri_cell.appendChild(document.createTextNode(uri));

				var show_cell = document.createElement("td");

				var show_button = document.createElement("button");
				show_button.setAttribute("data-uri", uri);				
				show_button.appendChild(document.createTextNode("show"));

				show_cell.appendChild(show_button);
				
				var meta_row = document.createElement("tr");
				meta_row.appendChild(source_cell);
				meta_row.appendChild(type_cell);
				meta_row.appendChild(hash_cell);
				meta_row.appendChild(uri_cell);
				meta_row.appendChild(show_cell);								

				show_button.onclick = function(e){

					var el = e.target;
					var uri = el.getAttribute("data-uri");

					var details = document.getElementById("details-" + uri);

					if (details.style.display == "none"){
						details.style.display = "block";
						el.innerText = "hide";						
					}

					else {
						details.style.display = "none";
						el.innerText = "show";						
					}
				};
				
				var details = mapzen.whosonfirst.render.render_data(data);
				
				var details_cell = document.createElement("td");
				details_cell.setAttribute("colspan", "5");
				details_cell.setAttribute("id", "details-" + uri);
				details_cell.setAttribute("class", "details");
				details_cell.appendChild(details);

				var details_row = document.createElement("tr");
				details_row.appendChild(details_cell);
				
				table.append(meta_row);
				table.append(details_row);				
			}

			return table;
		}
	};

	return self;

})();	
