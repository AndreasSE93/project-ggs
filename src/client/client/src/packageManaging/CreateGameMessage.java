package packageManaging;


/* NOT DONE YET */

public class CreateGameMessage {
	final int PacketID = 101;
	public String gameName;
	
	public CreateGameMessage(){
		
	}
	
	public CreateGameMessage(String m){
		
		gameName = m;
		
	}
	
}
