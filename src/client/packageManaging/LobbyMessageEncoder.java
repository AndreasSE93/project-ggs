package packageManaging;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;


public class LobbyMessageEncoder{

	public LobbyMessageEncoder() {
		
	}



	
	public LobbyServerMessage decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		
		
		LobbyServerMessage lsm = new LobbyServerMessage();
		
		JSONArray JArr = obj.getJSONArray("UserList");
		for (int i=0; i<JArr.length(); i++) {
		   JSONObject l = JArr.getJSONObject(i);
		   HostRoom h = new HostRoom (l.getInt("RoomID"), l.getInt("MaxSize"), l.getInt("ClientCount"), l.getString("RoomName"), l.getString("GameName"));
		   lsm.UserList.add(h);
		}
		JSONArray JArr2 = obj.getJSONArray("GameHost");
		for (int i=0; i<JArr2.length(); i++) {
		    lsm.GameHost.add(JArr2.getString(i));
		}
		
		
		
		return lsm;
	}
}
