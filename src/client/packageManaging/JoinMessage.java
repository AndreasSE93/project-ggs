package packageManaging;

public class JoinMessage {
	final int PacketID = 102;
	public int RoomID;
	public String GameType, GameName;
	
	

	
	public JoinMessage() {
		
	}
	
	public JoinMessage(String m){
		
		this.RoomID = Integer.parseInt(m);
		
	}
	
}
