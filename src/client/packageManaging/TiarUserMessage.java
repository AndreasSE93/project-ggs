package packageManaging;



public class TiarUserMessage {

	final int PacketID = 201;
	public int Move;
	public int HasWon;
	public int IsDraw;
	public int[] Gameboard;
	public int isValid;
	public int player;
	
	public TiarUserMessage(int m , int player){
		
		Move = m;
		this.player = player;
	}
	
	public TiarUserMessage(){
		
	}

}
