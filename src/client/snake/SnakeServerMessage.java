package snake;

public class SnakeServerMessage {
	
	final int PacketID = 302;
	public SnakePlayer[] Players;
	public boolean clearBoard;

public SnakeServerMessage(int len){
	this.Players = new SnakePlayer[len];
}


}
