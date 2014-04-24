package packageManaging;

import org.json.JSONObject;
import org.json.JSONException;

import clientNetworking.Handler;

public class MessageHandler extends Handler {

	public String encode(Message mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
	    obj.put("message", mess.message);
	    obj.put("color", mess.color);
	    obj.put("username", mess.username);

		String message = obj.toString();
		return message;
		
	}

	public Message decode(String enc) throws JSONException {
		
		JSONObject obj = new JSONObject(enc);
		
		Message chatMessage = new Message();
		chatMessage.message = obj.getString("message");
		chatMessage.color =  obj.getString("color");
		chatMessage.username = obj.getString("username");
		
		return chatMessage;
		
	}	
}
