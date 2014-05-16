package snake;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;




public class SnakeMessageEncoder {

	
	
	public String encodeSnakeUserMessage(SnakeUserMessage mess) throws JSONException {
		
		JSONObject obj = new JSONObject();
		
	    obj.put("PacketID", mess.PacketID);
	    obj.put("PlayerID", mess.PlayerID);
	    obj.put("KeyEvent", mess.KeyEvent);
	   
		String message = obj.toString();
		return message;
		
		
		
	}
	
	public SnakeServerMessage decode(String mess) throws JSONException{
		JSONObject obj = new JSONObject(mess);
		
		
		SnakeServerMessage SSM = new SnakeServerMessage();
		double[]p;
		JSONArray JArr = obj.getJSONArray("PlayerArray");
		for (int i=0; i<JArr.length(); i++) {
			p = new double[10];
		   JSONObject l = JArr.getJSONObject(i);
		   JSONArray JArr2 = l.getJSONArray("Coordinates");
		   
		   for(int d =0; d<p.length;d++){
			   p[d]= JArr2.getDouble(d);
			  
			  
		   }
		   SnakePlayer sp = new SnakePlayer(l.getInt("PlayerID"), l.getDouble("PosX"), l.getDouble("PosY"), l.getBoolean("Alive"),l.getInt("Score"),p);
		   SSM.Players[i] = sp;
		}
		

		return SSM;
	}

}
