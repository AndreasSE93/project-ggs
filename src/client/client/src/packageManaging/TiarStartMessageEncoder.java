package packageManaging;

import org.json.JSONException;
import org.json.JSONObject;

public class TiarStartMessageEncoder {

	
	public TiarStartMessage decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		
		
		TiarStartMessage tum = new TiarStartMessage();

		tum.opponent = obj.getString("Opponent");
		tum.turn  = obj.getInt("Turn");
		
		return tum;
	}
}


