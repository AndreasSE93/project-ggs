package clientHandlers;

import java.io.IOException;

public interface HandlerInterface {

	public void sendMessage(String message);
	public String recieveMessage() throws IOException;
	public int retrieveId(String mess);
	public int init(); 
	
}
