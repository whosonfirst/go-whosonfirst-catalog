var mapzen = mapzen || {};
mapzen.whosonfirst = mapzen.whosonfirst || {};

mapzen.whosonfirst.inspector = (function(){

    var cache = {};

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
	    
	    var by_type = {};
	    var types = [];
	    var sources = [];
	    
	    for (var i=0; i < count; i++){
		
		var r = records[i];
		var t = r["type"];
		
		if (! by_type[t]){
		    
		    by_type[t] = [ r ];
		    types.push(t);
		}
		
		else {
		    by_type[t].push(r);
		}
	    }
	    
	    types.sort();
	    
	    var table = document.createElement("table");
	    
	    var source_header = document.createElement("th");
	    source_header.appendChild(document.createTextNode("source"));
	    
	    var type_header = document.createElement("th");
	    type_header.appendChild(document.createTextNode("type"));
	    
	    var hash_header = document.createElement("th");
	    hash_header.appendChild(document.createTextNode("hash"));
	    
	    var uri_header = document.createElement("th");
	    uri_header.appendChild(document.createTextNode("uri"));
	    
	    var timing_header = document.createElement("th");
	    timing_header.appendChild(document.createTextNode("timing"));
	    
	    var show_header = document.createElement("th");
	    show_header.appendChild(document.createTextNode("-"));
	    
	    var details_header = document.createElement("th");
	    details_header.appendChild(document.createTextNode("-"));
	    
	    var header_row = document.createElement("tr");
	    header_row.appendChild(source_header);
	    header_row.appendChild(type_header);
	    header_row.appendChild(hash_header);
	    header_row.appendChild(uri_header);
	    header_row.appendChild(timing_header);
	    header_row.appendChild(details_header);
	    
	    table.appendChild(header_row);
	    
	    for (var t in types){
		
		t = types[t];
		var type_records = by_type[t];
		
		var count_records = type_records.length;
		
		var by_source = {};
		var sources = [];
		
		for (var i=0; i < count_records; i++){
		    
		    var data = type_records[i];

		    var source = data["source"];				
		    var type = data["type"];
		    var uri = data["uri"];
		    var timing = data["timing"] / 1000000000;	// nanoseconds
		   
		    var hash = data["hash"];			
		   
		    console.log("DATA", uri, hash); 

		    var source_cell = document.createElement("th");
		    source_cell.appendChild(document.createTextNode(source));
		    
		    var type_cell = document.createElement("td");
		    type_cell.appendChild(document.createTextNode(type));
		    
		    var hash_cell = document.createElement("td");
		    hash_cell.appendChild(document.createTextNode(hash));

		    var uri_cell = document.createElement("td");
		    // uri_cell.setAttribute("data-remote-uri", uri);		  
		    uri_cell.appendChild(document.createTextNode(uri));
		    
		    var timing_cell = document.createElement("td");
		    timing_cell.appendChild(document.createTextNode(timing.toFixed(3) + "s"));

		    var details_cell = document.createElement("td");
		    details_cell.setAttribute("class", "click-me");
		    details_cell.setAttribute("data-remote-uri", uri);		  
		    details_cell.appendChild(document.createTextNode("details"));

		    var meta_row = document.createElement("tr");

		    if (type == "error"){
			meta_row.setAttribute("class", "error");
		    }

		    meta_row.appendChild(source_cell);
		    meta_row.appendChild(type_cell);
		    meta_row.appendChild(hash_cell);
		    meta_row.appendChild(uri_cell);
		    meta_row.appendChild(timing_cell);
		    meta_row.appendChild(details_cell);
		    
		    details_cell.onclick = function(e){
			
			var el = e.target;
			var uri = el.getAttribute("data-remote-uri");
		
			self.show_details(uri);
		    };

		    table.append(meta_row);

		    cache[uri] = data;   
		}
	    }
	    
	    return table;
	},

	'hide_details': function(){

	    var details = document.getElementById("details");
	    details.innerHTML = "";
	},
	
	'show_details': function(uri){

	    self.hide_details();

	    var data = cache[uri];
	    var body = data["body"];
	    var type = data["type"];

	    var wrapper = document.createElement("div");
	    wrapper.setAttribute("id", "source");

	    var feature = self.data_to_feature(data);
	    var props = feature["properties"];

	    props = mapzen.whosonfirst.render.render_data(props);
	    wrapper.appendChild(props);

	    var map = document.createElement("div");
	    map.setAttribute("id", "map");

	    var table = document.createElement("table");
	    var tr = document.createElement("tr");

	    var left_cell = document.createElement("td");
	    var right_cell = document.createElement("td");

	    left_cell.append(map);
	    right_cell.append(wrapper);

	    tr.appendChild(left_cell);
	    tr.appendChild(right_cell);

	    table.appendChild(tr);

	    var details = document.getElementById("details");
	    details.appendChild(table);

	    self.draw_map(feature);
	},

	'draw_map': function(feature){

	    var props = feature["properties"];

	    if (! props){
		return;
	    }

	    var bbox = mapzen.whosonfirst.geojson.derive_bbox(feature);

	    var sw = L.latLng(bbox[1], bbox[0]);
	    var ne = L.latLng(bbox[3], bbox[2]);

	    console.log("BOUNDS", sw, ne);
	    console.log("TEST", (sw == ne));

            L.Mapzen.apiKey = document.body.getAttribute("data-mapzen-api-key");
	    
            map = L.Mapzen.map('map');

	    if ((sw["lat"] == ne["lat"]) && (sw["lng"] == ne["lng"])){
		map.setView(sw, 12);
	    }
	    
	    else {
		map.fitBounds([sw, ne]);
	    }

	    var layer = L.geoJSON(feature);
	    layer.addTo(map);
	},

	'data_to_feature': function(data){
	    
	    var type = data["type"];

	    if (type == "geojson"){
		return data["body"];
	    }
	    
	    if (type == "postgis"){
		
		var body = data["body"];
		body["Meta"] = JSON.parse(body["Meta"]);
		
		var geom;
		
		if (body["Geom"]){
		    geom = JSON.parse(body["Geom"]);
		}
		
		else {
		    geom = JSON.parse(body["Centroid"]);
		}
		
		// delete(body["Geom"]);
		// delete(body["Centroid"]);
		
		var props = body;
		
		var feature = {
		    type: "Feature",
		    geometry: geom,
		    properties: props,
		}
		
		return feature;
	    }

	    if (type == "elasticsearch"){

		var props = data["body"];
		var bbox = props["geom:bbox"];
		bbox = bbox.split(",");

		var swlon = bbox[0];
		var swlat = bbox[1];
		var nelon = bbox[2];
		var nelat = bbox[3];

		var coords = [
		    [ swlat, swlon ],
		    [ nelat, swlon ],
		    [ nelat, nelon ],
		    [ swlat, nelon ],
		    [ swlat, swlon ],		    
		];

		var geom = {
		    "type": "Polygon",
		    "coordinates": [ coords ],
		}

		var feature = {
		    "type": "Feature",
		    "geometry": geom,
		    "properties": props,
		};

		return feature;
	    }

	    if (type == "tile38"){

		var props = {};
		var geom = {};

		var id = data["id"];
		var re = "^" + id + "#";

		var body = data["body"];

		for (var k in body){

		    if (! k.match(re)){
			continue;
		    }

		    if (k != id + "#meta"){
			
			geom = body[k]["object"];
			// delete(body[k]["object"]);
		    }

		    props[k] = body[k];
		}

		var feature = {
		    "type": "Feature",
		    "geometry": geom,
		    "properties": props,
		};

		return feature;
	    }

	    console.log("FIX ME", data);
	    return data["body"];
	}

    };
    
    return self;
    
})();	
