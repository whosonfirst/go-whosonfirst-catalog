var mapzen = mapzen || {};
mapzen.whosonfirst = mapzen.whosonfirst || {};

mapzen.whosonfirst.catalog = (function(){

	var self = {

		'lookup': function(id, cb){

			var url = 'http://' + location.host + '/id/' + id;

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

			for (var i=0; i < count; i++){

				var data = records[i];

				var source = data["source"];				
				var uri = data["uri"];
				var hash = data["hash"];			

				var source_cell = document.createElement("td");
				source_cell.appendChild(document.createTextNode(source));

				var hash_cell = document.createElement("td");
				hash_cell.appendChild(document.createTextNode(hash));

				var uri_cell = document.createElement("td");
				uri_cell.appendChild(document.createTextNode(uri));

				var show_cell = document.createElement("td");
				show_cell.setAttribute("data-uri", uri);				
				show_cell.appendChild(document.createTextNode("show"));
				
				var meta_row = document.createElement("tr");
				meta_row.appendChild(source_cell);
				meta_row.appendChild(hash_cell);
				meta_row.appendChild(uri_cell);
				meta_row.appendChild(show_cell);								

				show_cell.onclick = function(e){

					var el = e.target;
					var uri = el.getAttribute("data-uri");

					var details = document.getElementById("details-" + uri);

					if (details.style.display == "block"){
						details.style.display = "none";
						el.innerText = "show";						
					}

					else {
						details.style.display = "block";
						el.innerText = "hide";
					}
				};
				
				var details = mapzen.whosonfirst.render.render_data(data);
				
				var details_cell = document.createElement("td");
				details_cell.setAttribute("colspan", "4");				
				details_cell.appendChild(details);

				var details_row = document.createElement("tr");	
				details_row.setAttribute("id", "details-" + uri);
				details_row.setAttribute("class", "details");				
				details_row.appendChild(details_cell);
				
				table.append(meta_row);
				table.append(details_row);				
			}

			return table;
		}
	};

	return self;

})();	
