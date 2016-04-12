$( document ).ready(function() {
	
	var conexion =new WebSocket('ws://localhost:8000/ws');
	conexion.onopen = function(){
	conexion.onmessage = function(response){
			val = $("#chat_area").val();
    		$("#chat_area").val(val + "\n" + response.data); 
		}
	}

    $('#form_message').on('submit', function(e) {
    	e.preventDefault();
    	conexion.send($('#msg').val());
    	$('#msg').val("")

    });
});