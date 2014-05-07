package packageManaging;

public class TiarUserMessage {

	final int PacketID = 200;
	int isDraw;
	int isVictorious;
	public String Move;
	
	public TiarUserMessage(int d, int v, String m){
		isDraw = d;
		isVictorious = v;
		Move = m;
		
	}
	

}
