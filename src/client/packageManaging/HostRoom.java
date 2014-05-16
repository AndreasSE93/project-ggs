package packageManaging;

public class HostRoom {
	public int RoomID;
	public int MaxSize;
	public int ClientCount;
	public String RoomName;
	public String GameName;
	public int Player;
	
	public HostRoom(){
		
	}
	public HostRoom (int RID, int MaxSize, int ClientCount, String RoomName, String GameName) {
		this.RoomID = RID;
		this.MaxSize = MaxSize;
		this.ClientCount = ClientCount;
		this.RoomName = RoomName;
		this.GameName = GameName;
		
	}
	
	@Override
	public String toString(){
		return this.RoomName;
	}
}
