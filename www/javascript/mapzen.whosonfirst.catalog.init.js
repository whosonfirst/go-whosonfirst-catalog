window.addEventListener("load", function load(event){
	
	var button = document.getElementById("lookup");
	
	button.onclick = function(){

	    try {
    		var id = document.getElementById("wof_id");

		id = id.value;
		id = id.trim();

		if (id == ""){
		    return false;
		}

		id = parseInt(id);

		if (id < 0){
		    return false;
		}

		mapzen.whosonfirst.catalog.lookup(id, function(rsp){
		    
		    var r = mapzen.whosonfirst.catalog.render(rsp);
		    
		    var results = document.getElementById("results");
		    results.innerHTML = "";

		    results.appendChild(r);
		});
	    }
	    
	    catch (e){
		console.log("WHAT", e);
	    }
	    
    	    return false;
	};
});
