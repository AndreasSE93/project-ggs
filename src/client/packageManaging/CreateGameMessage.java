package packageManaging;


/**
 * Used to create a generic gameroom on the server with a fixed size.
 * 
 * 
 * 
 * @author Grupp 8
 *
 */

public class CreateGameMessage {
	final int PacketID = 101;
	public String gameName;
	public int maxSize;
	public String roomName;
	
	public CreateGameMessage(){
		
	}
	
	public CreateGameMessage(String m, int maxSize, String rm){
		
		gameName = m;
		this.maxSize = maxSize;
		this.roomName= rm;
		
	}
	
}
