package packageManaging;


/* NOT DONE YET */

public class CreateGameMessage {
	final int PacketID = 101;
	public String gameName;
	public int maxSize;
	
	public CreateGameMessage(){
		
	}
	
	public CreateGameMessage(String m, int maxSize){
		
		gameName = m;
		this.maxSize = maxSize;
		
	}
	
}
