package clientHandlers;

import java.io.IOException;

import packageManaging.StageFlipper;

public interface HandlerInterface {

	public void sendMessage(String message);
	public String receiveMessage() throws IOException;
	public int retrieveId(String mess);
	public StageFlipper init(StageFlipper sf); 
	
}
