package packageManaging;

import org.json.JSONException;
//import org.json.JSONObject;
import org.json.JSONObject;

public class KickMessageEncoder {

	public KickMessage decode(String mess) throws JSONException {
		//JSONObject obj = new JSONObject(mess);
		
		KickMessage km = new KickMessage();
		
		return km;
		
	}
	
	public String encode(KickMessage mess) throws JSONException{
		
		JSONObject obj = new JSONObject();
		
	    obj.put("PacketID", mess.PacketID);
	   
		String message = obj.toString();
		return message;
		
		
	}
}
