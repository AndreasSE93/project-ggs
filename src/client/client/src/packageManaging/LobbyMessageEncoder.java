package packageManaging;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;


public class LobbyMessageEncoder{

	public LobbyMessageEncoder() {
		
	}

	
	public String encode(LobbyClientMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
	    obj.put("actionType", mess.actionType);
	    obj.put("actionId", mess.actionId);
	    obj.put("Message", mess.mess);
	   
		String message = obj.toString();
		return message;
		
	}

	
	public LobbyServerMessage decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		
		
		LobbyServerMessage lsm = new LobbyServerMessage();
		
		JSONArray JArr = obj.getJSONArray("UserList");
		for (int i=0; i<JArr.length(); i++) {
		    lsm.UserList.add( JArr.getString(i));
		}
		JSONArray JArr2 = obj.getJSONArray("GameHost");
		for (int i=0; i<JArr2.length(); i++) {
		    lsm.GameHost.add(JArr2.getString(i));
		}
		
		lsm.s = obj.getString("Message");
		
		return lsm;
	}
}
