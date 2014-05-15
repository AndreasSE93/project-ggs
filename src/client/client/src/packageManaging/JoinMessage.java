package packageManaging;

public class JoinMessage {
	final int PacketID = 102;
	public int RoomID;
	public String GameType, GameName;
	
	
	/* NOT DONE YET */
	
	public JoinMessage() {
		
	}
	
	public JoinMessage(String m){
		
		this.RoomID = Integer.parseInt(m);
		
	}
	
}
