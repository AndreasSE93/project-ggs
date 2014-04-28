package packageManaging;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;


public class LobbyHandler extends Handler {

	public LobbyHandler() {
		
	}

	
	public String encode(LobbyClientMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
	    obj.put("actionType", mess.actionType);
	    obj.put("actioneId", mess.actionId);
	   
		String message = obj.toString();
		return message;
		
	}

	
	public LobbyServerMessage decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		
		
		LobbyServerMessage lsm = new LobbyServerMessage();
		
		JSONArray JArr = obj.getJSONArray("userList");
		for (int i=0; i<JArr.length(); i++) {
		    lsm.userList.add( JArr.getString(i));
		}
		JSONArray JArr2 = obj.getJSONArray("userList");
		for (int i=0; i<JArr2.length(); i++) {
		    lsm.gameHost.add(JArr2.getString(i));
		}

		
		return lsm;
	}
}
