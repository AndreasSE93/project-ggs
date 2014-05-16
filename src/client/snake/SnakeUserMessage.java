package snake;

public class SnakeUserMessage {
	public final int PacketID = 301;
	public int PlayerID;
	public String KeyEvent;
	
	public SnakeUserMessage(int id, String event){
		PlayerID = id;
		KeyEvent = event;
	}
	

}
