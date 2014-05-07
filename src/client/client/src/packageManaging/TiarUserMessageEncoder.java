package packageManaging;

import org.json.JSONException;
import org.json.JSONObject;

public class TiarUserMessageEncoder {

	public String encode(TiarUserMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
	    obj.put("isDraw", mess.isDraw);
	    obj.put("isVictorious", mess.isVictorious);
	    obj.put("PacketID", mess.PacketID);
	   
		String message = obj.toString();
		return message;
	
	}
	
	public TiarUserMessage decode(String enc) throws JSONException {
		
		JSONObject obj = new JSONObject(enc);
		
		TiarUserMessage TiarMessage = new TiarUserMessage(obj.getInt("isDraw"), obj.getInt("isVictorious"), obj.getString("Move"));
	
		return TiarMessage;
		
		
	}
}
