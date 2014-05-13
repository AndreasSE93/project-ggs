package packageManaging;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

public class TiarUserMessageEncoder {

	public String encode(TiarUserMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		

	    obj.put("PacketID", mess.PacketID);
	    obj.put("Move", mess.Move);
	    obj.put("Player", mess.player);
	   
		String message = obj.toString();
		return message;
	
	}
	
	public TiarUserMessage decode(String enc) throws JSONException {
		
		JSONObject obj = new JSONObject(enc);
		
		TiarUserMessage TiarMessage = new TiarUserMessage();
		JSONArray JArray = obj.getJSONArray(("GameBoard"));
		int[] array = new int[9];
			for (int i = 0; i < array.length; i++){
			array[i] = JArray.getInt(i); 
			}
		TiarMessage.Gameboard = array;
		TiarMessage.HasWon = obj.getInt("HasWon");
		TiarMessage.IsDraw = obj.getInt("IsDraw");
		TiarMessage.isValid = obj.getInt("IsValid");
	
		return TiarMessage;
		
		
	}
}
