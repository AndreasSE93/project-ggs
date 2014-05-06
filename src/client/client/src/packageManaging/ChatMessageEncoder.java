package packageManaging;

import org.json.JSONObject;
import org.json.JSONException;


public class ChatMessageEncoder extends Handler {
    
	public String encode(Message mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
	    obj.put("message", mess.message);
	    obj.put("PacketID", mess.PacketID);
	    obj.put("user", mess.user);

		String message = obj.toString();
		return message;
		
	}

	public Message decode(String enc) throws JSONException {
		
		JSONObject obj = new JSONObject(enc);
		
		Message chatMessage = new Message();
		chatMessage.message = obj.getString("message");
		chatMessage.user = obj.getString("user");
		
		return chatMessage;
		
	}


}
