package clientHandlers;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.IOException;

import org.json.JSONException;
import org.json.JSONObject;

import clientNetworking.NetManager;



/* Handler and initializer for a Tic Tac Toe game */

public class TiarHandler implements HandlerInterface, ActionListener{

	NetManager network;
	
	TiarHandler(NetManager net){
		this.network = net;
		
	}

	public void init(){
		
		/* Create new Tiar lobby, probably with a graphical three in a row and a chat. 
		 * Add actionLsiteners to resp.. */
		
		
		
		
		while (true) {
			try {
				String mess = recieveMessage();
			
				int id = retrieveId(mess);
				decodeAndRender(id, mess);
				
			} catch (IOException | JSONException e) {
				e.printStackTrace();
			}
		}
	}
	
	public void decodeAndRender(int id, String message) throws JSONException {
		switch(id){
		case 100: //Chat message
				break;
				
		case 200: //Move message
				break;
				
		default: //Should not come here
				break;
		
		
		}
		
		
		
	}
	
	public void sendMessage(String message) {
		network.send(message);
		
	}

	
	public String recieveMessage() throws IOException {
		return network.receiveMessage();
	}
	
	
	public int retrieveId(String mess){
		try {
			JSONObject obj = new JSONObject(mess);
			return obj.getInt("PacketID");
		} catch (JSONException e) {
			
		}
		return 0;
		
	}

	
	public void actionPerformed(ActionEvent arg0) {
		// TODO Auto-generated method stub
		
	}

}
