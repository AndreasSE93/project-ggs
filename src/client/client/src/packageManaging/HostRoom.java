package packageManaging;

public class HostRoom {
	public int RoomID;
	public int MaxSize;
	public int ClientCount;
	public String RoomName;
	
	public HostRoom (int RID, int MaxSize, int ClientCount, String RoomName) {
		this.RoomID = RID;
		this.MaxSize = MaxSize;
		this.ClientCount = ClientCount;
		this.RoomName = RoomName;
		
	}
	
	@Override
	public String toString(){
		return this.RoomName;
	}
}
