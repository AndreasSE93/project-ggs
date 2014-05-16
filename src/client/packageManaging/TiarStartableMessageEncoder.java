package packageManaging;

import org.json.JSONException;
import org.json.JSONObject;

public class TiarStartableMessageEncoder {

	
	public TiarStartableMessage decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		
		TiarStartableMessage tum = new TiarStartableMessage();

		tum.isStartable = obj.getBoolean("IsStartable");
		
		return tum;
	}
}


