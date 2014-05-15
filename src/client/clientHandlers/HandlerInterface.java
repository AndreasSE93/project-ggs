package clientHandlers;

import java.io.IOException;

import org.json.JSONException;

public interface HandlerInterface {

	public void sendMessage(String message);
	public String recieveMessage() throws IOException;
	public int retrieveId(String mess);
	public void decodeAndRender(int id, String message) throws JSONException;
	public int init(); 
	
}
