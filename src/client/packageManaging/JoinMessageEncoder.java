package packageManaging;

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
		
		
		JoinMessage lsm = new JoinMessage();

		lsm.RoomID = obj.getInt("RoomID");
		
		return lsm;
	}
}

	

