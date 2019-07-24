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

		var spinner = document.getElementById("spinner");
		spinner.style.display = "inline";
		mapzen.whosonfirst.inspector.lookup(id, function(rsp){
		    
			spinner.style.display = "none";

		    var r = mapzen.whosonfirst.inspector.render(rsp);
		    
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
