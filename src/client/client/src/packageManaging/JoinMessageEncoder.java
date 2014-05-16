package packageManaging;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

public class JoinMessageEncoder {

	public JoinMessageEncoder(){
		
	}
	/* NOT DONE YET */
	
public String encode(JoinMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
		obj.put("PacketID", mess.PacketID);
	    obj.put("RoomID", mess.RoomID);
	    
	   
		String message = obj.toString();
		return message;
		
	}

	
	public JoinMessage decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		System.out.println(mess);
		
		JoinMessage lsm = new JoinMessage();
		JSONObject i = obj.getJSONObject("hostRoom");

		lsm.RoomID = i.getInt("RoomID");
		lsm.GameType = i.getString("GameType");
		lsm.GameName = i.getString("GameName");
		
		return lsm;
	}
}

	

