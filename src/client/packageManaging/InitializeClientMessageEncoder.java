package packageManaging;

import org.json.JSONException;
import org.json.JSONObject;

public class InitializeClientMessageEncoder {

	public String encode(InitializeClientMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
		obj.put("PacketID", mess.PacketID);
	    obj.put("UserName", mess.userName);
	    
	   
		String message = obj.toString();
		return message;
		
	}

}
