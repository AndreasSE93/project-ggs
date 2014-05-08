package packageManaging;


import org.json.JSONException;
import org.json.JSONObject;

public class CreateGameMessageEncoder {

	public CreateGameMessageEncoder(){
		
	}
	/* NOT DONE YET */
	
	
public String encode(CreateGameMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
	    obj.put("gameName", mess.gameName);
	    
	   
		String message = obj.toString();
		return message;
		
	}

	
	public CreateGameMessage decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		
		
		CreateGameMessage lsm = new CreateGameMessage();

		lsm.gameName = obj.getString("gameName");
		
		return lsm;
	}
}

	

