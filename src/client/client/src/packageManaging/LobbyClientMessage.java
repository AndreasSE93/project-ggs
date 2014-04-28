package packageManaging;

public class LobbyClientMessage {

	public static final int HOST = 0;
	public static final int JOIN = 1;
	
	public int actionType;
	public int actionId;
	
	public String mess;

	public LobbyClientMessage(int actionType, int id, String s) {
		this.actionType = actionType;
		this.actionId = id; 
		this.mess = s;
		
	}
	
	
}
