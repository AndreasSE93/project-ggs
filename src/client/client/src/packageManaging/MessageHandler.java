package packageManaging;

import org.json.JSONObject;
import org.json.JSONException;


public class MessageHandler extends Handler {
    
	public String encode(Message mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
	    obj.put("message", mess.message);
	    obj.put("id", mess.id);
	    obj.put("user", mess.user);

		String message = obj.toString();
		return message;
		
	}

	public Message decode(String enc) throws JSONException {
		
		JSONObject obj = new JSONObject(enc);
		
		Message chatMessage = new Message();
		chatMessage.message = obj.getString("message");
		chatMessage.id =  obj.getInt("id");
		chatMessage.user = obj.getString("user");
		
		return chatMessage;
		
	}


}
