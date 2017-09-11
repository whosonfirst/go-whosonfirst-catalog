window.addEventListener("load", function load(event){
	
	var button = document.getElementById("lookup");
	
	button.onclick = function(){

		try {
    			var id = document.getElementById("wof_id");
			id = parseInt(id.value);

			mapzen.whosonfirst.catalog.lookup(id, function(rsp){

				var r = mapzen.whosonfirst.catalog.render(rsp);
				
				var results = document.getElementById("results");
				results.appendChild(r);
			});
		}

		catch (e){
			console.log("WHAT", e);
		}
		
    		return false;
	};
});
