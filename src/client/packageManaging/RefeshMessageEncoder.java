package packageManaging;

import org.json.JSONException;
import org.json.JSONObject;

public class RefeshMessageEncoder {

	public String encode (RefreshMessage mess) throws JSONException {
	JSONObject obj = new JSONObject();
	
	obj.put("PacketID", mess.PacketID);
    
    
   
	String message = obj.toString();
	return message;

	}
/*
	public  RefreshMessage decode (String mess) throws JSONException {
	JSONObject obj = new JSONObject(mess);
		
		
		RefreshMessage lsm = new RefreshMessage();

	
		
		return lsm;
		
	}
*/
}
