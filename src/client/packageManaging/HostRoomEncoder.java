package packageManaging;

import org.json.JSONException;
import org.json.JSONObject;

public class HostRoomEncoder {

	public HostRoom decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		
		
		HostRoom hr = new HostRoom();

		hr.Player = obj.getInt("Player");
		
		
		return hr;
	}

}
