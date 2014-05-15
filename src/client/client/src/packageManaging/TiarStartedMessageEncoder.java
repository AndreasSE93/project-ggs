package packageManaging;

import org.json.JSONException;
import org.json.JSONObject;

public class TiarStartedMessageEncoder {
	
	public TiarStartedMessage decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		
		TiarStartedMessage tum = new TiarStartedMessage();

		tum.started = obj.getBoolean("Started");
		tum.playerID = obj.getInt("Player");
		
		return tum;
	}
}
