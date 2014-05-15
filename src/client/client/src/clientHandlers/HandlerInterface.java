package clientHandlers;

import java.io.IOException;

public interface HandlerInterface {

	public void sendMessage(String message);
	public String receiveMessage() throws IOException;
	public int retrieveId(String mess);
	public void init(); 
	
}
