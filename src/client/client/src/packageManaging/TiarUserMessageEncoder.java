package packageManaging;

import org.json.JSONException;
import org.json.JSONObject;

public class TiarUserMessageEncoder {

	public String encode(TiarUserMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		

	    obj.put("PacketID", mess.PacketID);
	    obj.put("Move", mess.Move);
	   
		String message = obj.toString();
		return message;
	
	}
	
	public TiarUserMessage decode(String enc) throws JSONException {
		
		JSONObject obj = new JSONObject(enc);
		
		TiarUserMessage TiarMessage = new TiarUserMessage(obj.getString("Move"));
	
		return TiarMessage;
		
		
	}
}
