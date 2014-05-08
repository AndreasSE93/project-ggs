package packageManaging;

import org.json.JSONException;
import org.json.JSONObject;

public class JoinMessageEncoder {

	public JoinMessageEncoder(){
		
	}
	/* NOT DONE YET */
	
public String encode(JoinMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
	    obj.put("joinName", mess.joinName);
	    
	   
		String message = obj.toString();
		return message;
		
	}

	
	public JoinMessage decode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		
		
		JoinMessage lsm = new JoinMessage();

		lsm.joinName = obj.getString("joinName");
		
		return lsm;
	}
}

	

